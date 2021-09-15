# lgraph
This is a generalized graph package with clean and transparent API for the Go language.

## Features
- Lightweight and Fast.
- Native Go implementation.
- Support all data types.

## Requirements
- Go 1.9 or higher. We aim to support the 3 latest versions of Go.

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get -u github.com/lovesaroha/lgraph
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage

### Create Graph

![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/graph.png)

```Golang
  // Create a graph (false for undirected , true for directed).
  graph := lgraph.Create(false)

  // Add edges in graph.
  graph.AddEdge(0, 1)
  graph.AddEdge(0, 2)
  graph.AddEdge(0, 5)
  graph.AddEdge(1, 2)
  graph.AddEdge(5, 3)
  graph.AddEdge(3, 2)
  graph.AddEdge(3, 4)

  // Show path between 0 and 4.
  graph.PrintShortestPathBetween(0, 4)


```
