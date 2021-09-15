/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lgraph

import "fmt"

// Graph object.
type GraphObject struct {
	vertices []vertexObject
	directed bool
	weighted bool
}

// Vertex object.
type vertexObject struct {
	value interface{}
	next  *edgeObject
}

// Edge object.
type edgeObject struct {
	vertexA interface{}
	vertexB interface{}
	weight  float64
	flow    float64
	next    *edgeObject
}

// This function add node to given vertex.
func (vertex *vertexObject) addNode(valueA interface{}, valueB interface{}, weight float64) {
	var newNode = edgeObject{vertexA: valueA, vertexB: valueB, weight: weight}
	if vertex.next == nil {
		// No node is added.
		vertex.next = &newNode
		return
	}
	newNode.next = vertex.next
	vertex.next = &newNode
}

// This function find given vertex.
func (graph GraphObject) findVertex(value interface{}) vertexObject {
	for _, vertex := range graph.vertices {
		if vertex.value == value {
			return vertex
		}
	}
	return vertexObject{}
}

// This function save edge in graph.
func (graph *GraphObject) saveEdge(valueA interface{}, valueB interface{}, weight float64) {
	var vertexAUpdated bool
	var vertexBUpdated bool
	for i, vertex := range graph.vertices {
		if vertex.value == valueA {
			graph.vertices[i].addNode(valueA, valueB, weight)
			vertexAUpdated = true
		}
		if vertex.value == valueB && !graph.directed {
			graph.vertices[i].addNode(valueA, valueB, weight)
			vertexBUpdated = true
		}
	}
	if !vertexAUpdated {
		var vertexA = vertexObject{value: valueA}
		vertexA.addNode(valueA, valueB, weight)
		graph.vertices = append(graph.vertices, vertexA)
	}
	if !vertexBUpdated && !graph.directed {
		var vertexB = vertexObject{value: valueB}
		vertexB.addNode(valueA, valueB, weight)
		graph.vertices = append(graph.vertices, vertexB)
	}
}

// This function add edge in graph between given vertex.
func (graph *GraphObject) AddEdge(valueA interface{}, valueB interface{}) {
	graph.saveEdge(valueA, valueB, 0)
}

// This function add weighted edge in graph between given vertex.
func (graph *GraphObject) AddWeightedEdge(valueA interface{}, valueB interface{}, weight float64) {
	graph.weighted = true
	graph.saveEdge(valueA, valueB, weight)
}

// This function show other vertex of given edge.
func (edge edgeObject) other(value interface{}) interface{} {
	if edge.vertexA == value {
		return edge.vertexB
	}
	return edge.vertexA
}

// This function return adjacent vertices.
func (graph GraphObject) PrintAdjacent(value interface{}) {
	fmt.Println("Vertices adjacent to:", value)
	var vertex = graph.findVertex(value)
	var currentNode = vertex.next
	for {
		if currentNode == nil {
			// No node found.
			return
		}
		fmt.Println(currentNode.other(value))
		currentNode = currentNode.next
	}
}

// This function find total connected components in graph.
func (graph GraphObject) TotalConnectedComponents() int {
	var totalComponents int
	var visited []interface{}
	var edgeTo []edgeToObject
	for _, vertex := range graph.vertices {
		if !isVisited(visited, vertex.value) {
			// Depth first serach.
			graph.depthFirstSearchRecursive(vertex.value, &visited, &edgeTo)
			totalComponents++
		}
	}
	return totalComponents
}

// This function create new graph.
func Create(directed bool) GraphObject {
	return GraphObject{directed: directed}
}
