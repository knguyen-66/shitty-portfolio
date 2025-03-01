---
Title: Building a (kinda) high performance chat application with Python & Golang
Summary: Take a look on my process of designing, building & testing a performance-prioritized application for a school project.
Tags:
    - python
    - golang
    - chat-app
    - performance
    - k3s
Date: 30/09/2024
---

## 1. Context

On my junior year, I need to work on a project as a preparation course before working on the graduation thesis. At that time I had done a couple school projects / big assignments, got a bit of experience on how to build a system, working with API and UI interfaces, but just for general, "educational" use cases at most. There were some fun & exciting projects too, but not too much.

So for that project, I feel like I want to tackle something more on the technical side, mainly trying to work on something that will:

- Process & provide interaction in real-time
- Focus on the performance side (basically want to make it run fast)
- Give me an excuse to use new tech stack & toolkit

![Tired Mr.Incredible](https://i.imgflip.com/92ecuk.jpg)
Phone shop inventory, book shop inventory, flower shop inventory, ... what else?

I was really inspired from [this blog from Discord](https://discord.com/blog/how-discord-stores-billions-of-messages), and I was able to borrow some spare PC hardware from my classmate (I'm a self-hoster, cloud is scary, and I just don't wanna be racked up with a bill at the end of the day). So I decided to build & deploy a chat app. **Basically a Discord clone**. Kinda reasonable, yeah?

## 2. Preparation

### 2.1. Research phase

It sounded really exciting, but I lacked a lot of information on how to build it the right way. I didn't even know how a chat app works back then :P

So like other devs, I searched for articles and watched YouTube vids repeatedly. Then I found this [WhatsApp system design overview](https://www.codekarle.com/system-design/Whatsapp-system-design.html) (ofc it's from an Indian chad).

![Chat application system design overview](/static/imgs/blog/01-whatsapp-codekarle.png)
System Design Overview of a chat application

Basically, there are a couple components that I need to take notice on:

- **The connection layer** between the server and all the active users in order to:
  - Keep the connections open for each users & remember what user is currently connected to it.
  - Handle the message flow for both directions.
  - Expose APIs for other types of operations (user search, get group images, ...)

- **The storage structure** to handle the application data. It needs to have good *read & write performance*, and provides *the ability to scale* (unless we are working with an end-to-end app where data is stored on user's machine, chat messages gonna be all over the places).

- **The main app structure**: of course we need it. But there are a couple things to consider: I also want it to be scalable, because of that we also need a good communication interface between services that we are expecting to build & scale up.

- **The hardware infrastructure** to host the application. It will also host all the additional services for the application to use. We won't talk about cloud, as mentioned before and also for educational purposes on how to design a DevOps flow effectively.

> But is reinstalling the OS 15 times by yourself educational? I hope so -_-!

By the way, I also recommend this [Uber/Lyft system design overview](https://www.youtube.com/watch?v=R_agd5qZ26Y) video, it helped me a lot on understanding how to design an application that focused on scalability aspects.

### 2.2. Design phase

#### Application features

As I was focusing a lot on the technical aspect, I've decided to keep the feature list short. For this app, I wanted it to be a group-based chat app, with options to:

- Create & login to personal account
- Search for users / groups
- Create group, add people to group, edit group information (there would be user roles on each group)
- Send text / file messages
- Receive notifications on joined group's activities

![Main app workflow](/static/imgs/blog/01-app-workflow.png)
The expected application workflow with available features, with the blue features as the main route

#### Nitpicking on the tech stack

For the **databases**, as for other big chat applications, they usually use a cluster of database nodes other than a single instance, especially using a "wide-column store, distributed" database like **Cassandra**, or **ScyllaDB** for Discord. There are a couple things to know for this type of database:

- **Wide-column**: It's a type of data structure (column-oriented to be specific) that these databases used. Basically, a wide-column database stores data in tables, where the attribute names and formats are structured in columns.

```text
# row-oriented, as like most traditional SQL databases:

| id1 | fullName1 | age1 | email1 |
| id2 | fullName2 | age2 | email2 |
| id3 | fullName3 | age3 | email3 |
| ... | ...       | ...  | ...    |

# column-oriented:

| id1       | id2       | id3       | id4       | ... |
| fullName1 | fullName2 | fullName3 | fullName4 | ... |
| age1      | age2      | age3      | age4      | ... |
| email1    | email2    | email3    | email4    | ... |
```

- **Distributed**: These databases are designed specifically so that the data can be splited and organized onto multiple database instances (nodes) located on different sites, therefore decrease the latency between calls to the database node and allow it to be scaled as needed by adding or removing nodes from the system.

I chose to use ScyllaDB for this project, as it's said to be a more performant version of Cassandra. I also wanted to try out a new database type, so why not?

```yml
# docker-compose-staging.yml
# A Scylla DB cluster with 3 nodes, db-scylla-1 is the main node for the app to connect to

db-scylla-1:
  image: scylladb/scylla
  restart: unless-stopped
  ports:
    - '127.0.0.1:9042:9042'
  command: --smp 2 --memory 2500M --reserve-memory 4G
  networks:
    - network-1

db-scylla-2:
  image: scylladb/scylla
  restart: unless-stopped
  depends_on:
    - db-scylla-1
  ports:
    - '127.0.0.1:9043:9042'
  command: --smp 2 --memory 2500M --reserve-memory 4G --seeds=db-scylla-1
  networks:
    - network-1

db-scylla-3:
  image: scylladb/scylla
  restart: unless-stopped
  depends_on:
    - db-scylla-1
  ports:
    - '127.0.0.1:9044:9042'
  command: --smp 2 --memory 2500M --reserve-memory 4G --seeds=db-scylla-1
  networks:
    - network-1
```

I also add a Postgres instance for general user information and app configuration. It's a good choice for a relational database, I already have some experience with it, and it's safer to keep ACID properties for the user data.

For the **backend service**, I initially wanted to use Python for both HTTP and Websocket services, but as the project going on, I realized it's not really performant for the Websocket part. So I switched to Golang for the Websocket service, and keep Python for the HTTP service.

#### Data schema

