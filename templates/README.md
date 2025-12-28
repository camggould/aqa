## Developer Guide

templates use [Templ](https://templ.guide) to generate static HTML from `.templ` files. There are extensions in VSCode to help with syntax highlighting.

Install the CLI tool using `go get -tool github.com/a-h/templ/cmd/templ@latest`. This requires go 1.24+ as of the time of writing this documentation, but latest information can be found [here](https://templ.guide/quick-start/installation).

In order to generate the `.go` files from `.templ`, you need to run `go tool templ generate`. Once the go code is generated, you can run the code. Any time a template is updated, this needs to be rerun. Make sure to commit the generated code.