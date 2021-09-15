/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lgraph

import (
	"fmt"

	"./lqueue"
	"./lstack"
)

// Edge to object.
type edgeToObject struct {
	from interface{}
	to   interface{}
}

// Distance to object.
type distanceToObject struct {
	distance float64
	to       interface{}
}

// This function checks if visited.
func isVisited(visited []interface{}, value interface{}) bool {
	for _, v := range visited {
		if v == value {
			return true
		}
	}
	return false
}

// This function add to visited list.
func addToVisited(visited *[]interface{}, value interface{}) {
	for _, v := range *visited {
		if v == value {
			return
		}
	}
	*visited = append(*visited, value)
}

// This function add to edge to list.
func addToEdgeTo(edgeTo *[]edgeToObject, from interface{}, to interface{}) {
	var et = *edgeTo
	for i, v := range et {
		if v.to == to {
			et[i].from = from
			return
		}
	}
	*edgeTo = append(*edgeTo, edgeToObject{from: from, to: to})
}

// This function find egde to given vertex.
func edgeToVertex(edgeTo []edgeToObject, to interface{}) interface{} {
	for _, v := range edgeTo {
		if v.to == to {
			return v.from
		}
	}
	return nil
}

// This function add to distance to list.
func addToDistanceTo(distanceTo *[]distanceToObject, to interface{}, distance float64) {
	var dt = *distanceTo
	for i, d := range dt {
		if d.to == to {
			dt[i].distance = distance
			return
		}
	}
	*distanceTo = append(*distanceTo, distanceToObject{distance: distance, to: to})
}

// This function return distance to given vertex.
func distanceToVertex(distanceTo []distanceToObject, to interface{}) float64 {
	for _, d := range distanceTo {
		if d.to == to {
			return d.distance
		}
	}
	return 0
}

// This function perform breadth first search from given source and return visited vertices and path.
func (graph GraphObject) breadthFirstSearch(source interface{}) ([]interface{}, []edgeToObject) {
	var visited []interface{}
	var edgeTo []edgeToObject
	queue := lqueue.Create(0)
	queue.Enqueue(source)
	addToVisited(&visited, source)
	for {
		if queue.IsEmpty() {
			// Close loop if queue is empty.
			return visited, edgeTo
		}
		// Get vertex value from queue.
		var vertexValue = queue.Dequeue()
		var vertex = graph.findVertex(vertexValue)
		// For all adjacent vertices.
		var currentNode = vertex.next
		for {
			if currentNode == nil {
				// No adjacent node.
				break
			}
			if !isVisited(visited, currentNode.other(vertexValue)) {
				// Save edge info and add node value in queue.
				addToEdgeTo(&edgeTo, vertexValue, currentNode.other(vertexValue))
				queue.Enqueue(currentNode.other(vertexValue))
				addToVisited(&visited, currentNode.other(vertexValue))
			}
			currentNode = currentNode.next
		}
	}
}

// This function perform depth first search from given source and return visited vertices and path.
func (graph GraphObject) depthFirstSearch(source interface{}) ([]interface{}, []edgeToObject) {
	var visited []interface{}
	var edgeTo []edgeToObject
	stack := lstack.Create()
	stack.Push(source)
	for {
		if stack.IsEmpty() {
			// Stop loop if stack is empty.
			return visited, edgeTo
		}
		// Get vertex from stack mark it visited.
		var vertexValue = stack.Pop()
		addToVisited(&visited, vertexValue)
		var vertex = graph.findVertex(vertexValue)
		// For all adjacent vertices.
		var currentNode = vertex.next
		for {
			if currentNode == nil {
				// No adjacent node.
				break
			}
			if !isVisited(visited, currentNode.other(vertexValue)) {
				// Save edge info and add node value in stack.
				addToEdgeTo(&edgeTo, vertexValue, currentNode.other(vertexValue))
				stack.Push(currentNode.other(vertexValue))
			}
			currentNode = currentNode.next
		}
	}
}

// This function perform depth first search usign recursion from given source.
func (graph GraphObject) depthFirstSearchRecursive(value interface{}, visited *[]interface{}, edgeTo *[]edgeToObject) {
	addToVisited(visited, value)
	var vertex = graph.findVertex(value)
	var currentNode = vertex.next
	// All adjacent vertices.
	for {
		if currentNode == nil {
			// No adjacent node.
			return
		}
		if !isVisited(*visited, currentNode.other(value)) {
			// Save path info with vertex value.
			addToEdgeTo(edgeTo, value, currentNode.other(value))
			// Search if node is not visited.
			graph.depthFirstSearchRecursive(currentNode.other(value), visited, edgeTo)
		}
		currentNode = currentNode.next
	}
}

// This function perform dijkstra algorithm from given source to all vertices.
func (graph GraphObject) shortestPathFrom(source interface{}) ([]distanceToObject, []edgeToObject, []interface{}) {
	var distanceTo []distanceToObject
	var edgeTo []edgeToObject
	var visited []interface{}
	var vertexValue = source
	for {
		if isVisited(visited, vertexValue) {
			// No more vertex.
			return distanceTo, edgeTo, visited
		}
		addToVisited(&visited, vertexValue)
		// Find all adjacent vertices.
		var vertex = graph.findVertex(vertexValue)
		var currentNode = vertex.next
		for {
			if currentNode == nil {
				// No more adjacent node.
				break
			}
			var dToB = distanceToVertex(distanceTo, currentNode.other(vertexValue))
			var ndToB = distanceToVertex(distanceTo, vertexValue) + currentNode.weight
			if dToB == 0 || dToB > ndToB {
				// Update distance.
				addToDistanceTo(&distanceTo, currentNode.other(vertexValue), ndToB)
				addToEdgeTo(&edgeTo, vertexValue, currentNode.other(vertexValue))
			}
			currentNode = currentNode.next
		}
		vertexValue = findClosestUnvisited(distanceTo, visited).to
	}
}

// This function return next minimum distance vertex.
func findClosestUnvisited(distanceTo []distanceToObject, visited []interface{}) distanceToObject {
	var closest distanceToObject
	var min float64
	for _, d := range distanceTo {
		if isVisited(visited, d.to) {
			continue
		}
		if (min == 0 || (d.distance < min)) && d.distance != 0 {
			min = d.distance
			closest = d
		}
	}
	return closest
}

// This function print path.
func printPath(edgeTo []edgeToObject, valueA interface{}, valueB interface{}) {
	var currentVertex = valueB
	stack := lstack.Create()
	for {
		var from = edgeToVertex(edgeTo, currentVertex)
		stack.Push(currentVertex)
		if from == valueA || from == nil {
			// End of path.
			stack.Push(from)
			break
		}
		currentVertex = from
	}
	for {
		if stack.IsEmpty() {
			return
		}
		v := stack.Pop()
		if stack.IsEmpty() {
			fmt.Printf("%v", v)
		} else {
			fmt.Printf("%v -> ", v)
		}
	}
}

// This function show shortest path between given vertices.
func (graph GraphObject) PrintShortestPathBetween(valueA interface{}, valueB interface{}) {
	if graph.weighted {
		// Perform dijkstra.
		distanceTo, edgeTo, visited := graph.shortestPathFrom(valueA)
		if !isVisited(visited, valueB) {
			// Vertex b is not connected to vertex a so no path.
			fmt.Println("No path between ", valueA, valueB)
			return
		}
		fmt.Println("Shortest path between", valueA, "and", valueB, "is", distanceToVertex(distanceTo, valueB))
		printPath(edgeTo, valueA, valueB)
		return
	}
	// Perform breadth first search from vertex A.
	visited, edgeTo := graph.breadthFirstSearch(valueA)
	if !isVisited(visited, valueB) {
		// Vertex b is not connected to vertex a so no path.
		fmt.Println("No path between ", valueA, valueB)
		return
	}
	fmt.Println("Shortest path between", valueA, "and", valueB)
	printPath(edgeTo, valueA, valueB)
}
