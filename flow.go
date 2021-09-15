/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lgraph

import (
	"./lqueue"
)

// This function return residual capacity from given vertex to other.
func (edge edgeObject) residualCapacityTo(value interface{}) float64 {
	if value == edge.vertexB {
		// Forward from a to b.
		return edge.weight - edge.flow
	}
	// Backward edge from b to a.
	return edge.flow
}

// This function add residual flow to from given vertex.
func (edge *edgeObject) addResidualFlowTo(value interface{}, flow float64) {
	if value == edge.vertexB {
		// Forward from a to b.
		edge.flow += flow
		return
	}
	// Backward edge from b to a.
	edge.flow -= flow
}

// This function checks augmented path.
func (graph GraphObject) hasAugmentedPath(source interface{}, target interface{}) (bool, []edgeToObject, float64) {
	var visited []interface{}
	queue := lqueue.Create(0)
	var edgeTo []edgeToObject
	var minimum float64
	queue.Enqueue(source)
	addToVisited(&visited, source)
	// Perform breadth first search.
	for {
		if queue.IsEmpty() {
			// True if there is a path from source to target with capacity.
			return isVisited(visited, target), edgeTo, minimum
		}
		// Get vertex from queue.
		var vertexValue = queue.Dequeue()
		var vertex = graph.findVertex(vertexValue)
		// For all adjacent vertices.
		var currentNode = vertex.next
		for {
			if currentNode == nil {
				// No more adjacent nodes.
				break
			}
			var c = currentNode.residualCapacityTo(currentNode.other(vertexValue))
			if !isVisited(visited, currentNode.other(vertexValue)) && c > 0 {
				// Add to queue if not visited and has residual capacity.
				queue.Enqueue(currentNode.other(vertexValue))
				addToVisited(&visited, currentNode.other(vertexValue))
				addToEdgeTo(&edgeTo, vertexValue, currentNode.other(vertexValue))
				// Update minimum.
				if minimum == 0 || minimum > c {
					minimum = c
				}
			}
			currentNode = currentNode.next
		}
	}
}

// This function perform ford fulkerson to find max flow.
func (graph GraphObject) Maxflow(source interface{}, target interface{}) float64 {
	var maxFlow float64
	for {
		path, edgeTo, minimumCapacity := graph.hasAugmentedPath(source, target)
		if !path {
			// No more path from source to target.
			return maxFlow
		}
		var currentVertex = target
		for {
			var from = edgeToVertex(edgeTo, currentVertex)
			var vertex = graph.findVertex(from)
			var currentNode = vertex.next
			for {
				if currentNode == nil {
					break
				}
				if currentNode.other(from) == currentVertex {
					currentNode.addResidualFlowTo(currentVertex, minimumCapacity)
					break
				}
				currentNode = currentNode.next
			}
			if from == source || from == nil {
				// End of path.
				break
			}
			currentVertex = from
		}
		maxFlow += minimumCapacity
	}
}
