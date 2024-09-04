## Views

"Here lies all of my *views*... I'm not sure if the pun was worth it..."

I'm using [templ](https://github.com/a-h/templ) for content managment. Templ is
a templating system that uses Go's own template engine to make content
generation a little smoother.

### Content Generation?

Yes. Just like in React, we can pass properties to the components and read from
them in the different views to generate our webpages.

```go
package components

templ ComponentName( prop propType, ... ){
    // all component code
}
```

To have the components generate callable `Go` code, we have to run one
more command `templ generate`.

I have had some problems running `templ [cmd]` in the past and found that even
having `GOBIN`, `GOPATH`, and `GOROOT` set, you may get `templ: command not found`. At this point, I
type it in manually. I haven't been able to find where it's breaking at to tell
if it's and env thing or the face that I'm using a package manager for Golang.

### Attempt at a Dynamic Site

Now, the thing we have to keep in mind, is tracking. It's terrible I know, but
that's not the end goal for me. I simply want to load content from my own server at
input. The page `Render`s the data to the view without reloading. Who knew a
dynamic webpage could be made without spamming ads?
