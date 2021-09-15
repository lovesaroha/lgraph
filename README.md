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

### Create Graph (Undirected)

![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/1.png)

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
  graph.AddEdge(2, 4)

  // Show path between 0 and 4.
  graph.PrintShortestPathBetween(0, 4)


```
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/2.png)

### Create Weighted Graph (Directed)

![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/3.png)

```Golang
  // Create a graph (false for undirected , true for directed).
  graph := lgraph.Create(true)

  // Add edges in graph.
	graph.AddWeightedEdge(0, 1, 5.0)
	graph.AddWeightedEdge(0, 7, 8.0)
	graph.AddWeightedEdge(0, 4, 9.0)
	graph.AddWeightedEdge(1, 7, 4.0)
	graph.AddWeightedEdge(1, 2, 12.0)
	graph.AddWeightedEdge(1, 3, 15.0)
	graph.AddWeightedEdge(7, 2, 7.0)
	graph.AddWeightedEdge(7, 5, 6.0)
	graph.AddWeightedEdge(5, 2, 1.0)
	graph.AddWeightedEdge(5, 6, 13.0)
	graph.AddWeightedEdge(4, 7, 5.0)
	graph.AddWeightedEdge(4, 5, 4.0)
	graph.AddWeightedEdge(4, 6, 20.0)
	graph.AddWeightedEdge(2, 3, 3.0)
	graph.AddWeightedEdge(2, 6, 11.0)
	graph.AddWeightedEdge(3, 6, 9.0)

  // Show path between 0 and 6.
  graph.PrintShortestPathBetween(0, 6)
```
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/4.png)

## Examples

### Six Degrees of Kevin Bacon

```js
  // Movies data sample.
  { "apollo13" : ["Kevin Bacon" , "Tom Hanks" , "Ed Harris"] ,
    "The Truman Show" : ["Ed Harris" , "Jim Carrey" , "Laura Linney"] ,
    "Joker" : ["Joaquin Phoenix" , "Robert De Niro" , "Zazie Beetz"] ,
    "Bruce Almighty" : ["Jim Carrey" , "Steve Carell" , "Morgan Freeman"] ,
    "Lucy" ["Morgan Freeman" , "Scarlett Johansson"] ,
    "Her" ["Joaquin Phoenix" , "Scarlett Johansson" , "Rooney Mara"]
  }
```
```Golang
  // Create a graph (false for undirected , true for directed).
  graph := lgraph.Create(false)

  // Add edges in graph.
  graph.AddEdge("apollo13", "Kevin Bacon")
  graph.AddEdge("apollo13", "Tom Hanks")
  graph.AddEdge("apollo13", "Ed Harris")
  graph.AddEdge("The Truman Show" , "Ed Harris")
  graph.AddEdge("The Truman Show" , "Jim Carrey")
  graph.AddEdge("The Truman Show" , "Laura Linney")
  graph.AddEdge("Joker" , "Joaquin Phoenix")
  graph.AddEdge("Joker" , "Robert De Niro")
  graph.AddEdge("Joker" , "Zazie Beetz")
  graph.AddEdge("Bruce Almighty" , "Jim Carrey")
  graph.AddEdge("Bruce Almighty" , "Steve Carell")
  graph.AddEdge("Bruce Almighty" , "Morgan Freeman")
  graph.AddEdge("Lucy" , "Morgan Freeman")
  graph.AddEdge("Lucy" , "Scarlett Johansson")
  graph.AddEdge("Her" , "Joaquin Phoenix")
  graph.AddEdge("Her" , "Scarlett Johansson")
  graph.AddEdge("Her" , "Rooney Mara")

  // Show path between Kevin Bacon and Joaquin Phoenix.
  graph.PrintShortestPathBetween("Kevin Bacon", "Joaquin Phoenix")
```

![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/5.png)
