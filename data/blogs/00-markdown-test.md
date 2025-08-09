---
Title: Markdown Test View
Date: 01/01/1970
---

# Markdown test

## Heading 2

A paragraph: Dolor esse amet aute nulla. Veniam labore ut adipisicing nostrud. In consectetur dolore aliquip culpa in. Elit sint dolore aliquip laborum sint non non ullamco. Exercitation sunt incididunt laborum quis aliquip do sit laboris ipsum elit magna reprehenderit exercitation.

Another paragraph: Eiusmod duis veniam amet consectetur do fugiat esse non.

A small ` sudo apt install docker ` on some text to see if the code block margin is too large or not. Quis nisi pariatur velit minim veniam aliqua tempor quis nisi eiusmod cupidatat adipisicing elit mollit.

### Hello in heading 3

A line break

---

> Blockquote for Table 1: Table with header

| Syntax aligned left | Description aligned right | Price aligned center |
| :------------------ | ------------------------: | :------------------: |
| Header              |                     Title |         1000         |
| Paragraph           |                      Text |         2000         |

~~The world is flat.~~

That is so funny! :joy: 

I need to highlight these ==very important words==

H~2~O and X^2^

1. List item 1
   1. List item 1.1
   2. List item 1.2
   3. List item 1.3 
2. List item 2
   1. List item 2.1
      1. List item 2.1.1
      2. List item 2.1.2
3. List item 3

- List item 1
  - List item 1.1
  - List item 1.2
  - List item 1.3
- List item 2
  - List item 2.1
    - List item 2.1.1
    - List item 2.1.2
- List item 3

- [x] Write the press release
- [ ] Update the website
- [ ] Contact the media

```go
// a simple code block

package handlers

import (
	"net/http"
	"shitty-portfolio/internal/views"
)
// a really long command to test scrolling: Voluptate occaecat ullamco labore adipisicing commodo laboris. Et esse mollit proident quis eiusmod id voluptate officia.
func HandleHomePage(w http.ResponseWriter, r *http.Request) error {
	return RenderWithDefaultLayout(w, r, views.Home())
}

func HandleHomeTemplate(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, views.Home())
}

```

> Another blockquote
> > Blockquote inside a blockquote
> - Item 1
> - Item 2
