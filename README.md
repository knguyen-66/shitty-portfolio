# shitty-portfolio

As the name suggest.

A place to:

- Show people that I have a portfolio
- Document & link my projects (also want to improve my grammar on writing long English articles)
- Other types of showcase

Stack: Golang, HTMX, Templ, TailwindCSS

!["Approved by HTMX"](/internal/static/imgs/htmx-banner.png)

## Prerequisites

- Golang v1.24+: <https://go.dev/doc/install>
- SQLite: `apt install sqlite3`

## Build

- Install [Go](https://go.dev/doc/install) and clone the repository
- Download & install required dependencies

    ```bash
    make prepare
    ```

- For development, run the following command

    ```bash
    make dev
    ```

- For production, run the following command

    ```bash
    make build-prod
    make prod
    ```
