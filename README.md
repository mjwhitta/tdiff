# TDiff

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/tdiff)](https://goreportcard.com/report/github.com/mjwhitta/tdiff)

## What is this?

Dumb. No, honestly, this is a very dumb example of how to play with
Time objects in Go. Do NOT use this for your own projects, it is
merely sample code.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/tdiff
$ go install --ldflags "-s -w" --trimpath \
    github.com/mjwhitta/tdiff/cmd/tdiff@latest
```

Or install from source:

```
$ git clone https://github.com/mjwhitta/tdiff.git
$ cd tdiff
$ git submodule update --init
$ make install
```

**Note:** `make install` will install to `$HOME/.local/bin`.

## Usage

```
$ tdiff [-y] 1988-05-22 now
```

## Links

- [Source](https://github.com/mjwhitta/tdiff)
