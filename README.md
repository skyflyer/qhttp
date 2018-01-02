# qhttp

A very simple http server. Just run `qhttp` to serve files in the current directory and subdirectories on port 8000. Specify a port parameter (`qhttp -port 3000`) if you want to use a different port.

# Why

Just because I felt like it and coding is fun. I've been using `python -m SimpleHTTPServer` for years and it still works (it serves single client at a time). There are other similar utlities like node's [http-server](https://www.npmjs.com/package/http-server).

# Installation

* `brew install qhttp`

## From source

1. Clone this repo and
1. `go install`
1. Execute `qhttp` from a directory you wish to serve.