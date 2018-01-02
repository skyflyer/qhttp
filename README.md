# qhttp

A very simple http server. Just run `qhttp` to serve files in the current directory and subdirectories on port 8000. Specify a port parameter (`qhttp -port 3000`) if you want to use a different port.

# Why

Just because I felt like it and coding is fun. I've been using `python -m SimpleHTTPServer` for years and it still works (it serves single client at a time). There are other similar utlities like node's [http-server](https://www.npmjs.com/package/http-server).

# Installation

* Install using Homebrew: `brew install qhttp`
* Download [binary](https://github.com/skyflyer/qhttp/releases/download/v0.1.0/qhttp.v0.1.0.zip) and extract. Copy to `/usr/local/bin`.

## From source

1. `go get github.com/skyflyer/qhttp`
1. Execute `qhttp` from a directory you wish to serve.