# NOTE

I'm currently doing a complete re-write of this. Instead of using Hugo and Blogdown, I'm switching everything over to
Go/HTMX/Templ.

## Projects

Since I stopped working on this website, I have branched out further into embedded and software domains.<br>
Additional projects will be added.

# Website Re-Write

I'm re-writing my old personal website. I originally used R, Hugo, and blogdown.
For the re-write, I am using GO, HTMX, Templ, and possibly a css library. The
more I keep thinking about what I'm doing, I'm not sure I need HTMX for this.

If you'd like to follow along with this journey, I am (blogging)[https://dev.to/caffeineordeath/from-r-to-go-4pk8] it, and writing articles on fixing some the issues that arise along the way. You know, like `env` and `$PATH` problems.

## Packages Used (so far...)

(GO)[https://go.dev/]
There are a few needed packages, regardless of editor.
Gopls is the language server for Golang.
`go install golang.org/x/tools/gopls@latest`

(Templ)[https://github.com/a-h/templ]
`go get github.com/a-h/templ/cmd/templ@latest`

(HTMX)[https://htmx.org/]
`<script src="https://unpkg.com/htmx.org@2.0.2"></script>`
