/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lqueue

import (
	"fmt"
	"strconv"
	"strings"
)

// Queue object.
type QueueObject struct {
	first    *nodeObject
	last     *nodeObject
	priority int
}

// Node object structure.
type nodeObject struct {
	value interface{}
	next  *nodeObject
}

// This function checks if queue is empty.
func (queue QueueObject) IsEmpty() bool {
	return queue.first == nil && queue.last == nil
}

// Enqueue values function add multiple values in queue.
func (queue *QueueObject) EnqueueValues(values []interface{}) {
	for _, value := range values {
		queue.Enqueue(value)
	}
}

// Enqueue function add new value in queue.
func (queue *QueueObject) Enqueue(value interface{}) {
	var newNode = nodeObject{value: value}
	if queue.first == nil {
		// First and last points to new node.
		queue.first = &newNode
		queue.last = &newNode
		return
	}
	if queue.priority > 0 {
		// Ascending order.
		queue.enqueueMin(&newNode)
	} else if queue.priority < 0 {
		// Descending order.
		queue.enqueueMax(&newNode)
	} else {
		// No priority.
		if queue.last != nil {
			// Last node next pointer now points to new node.
			queue.last.next = &newNode
		}
		// Last pointer points to new node.
		queue.last = &newNode
	}
}

// This function add new value in queue in min order.
func (queue *QueueObject) enqueueMin(newNode *nodeObject) {
	if isLessThan(newNode.value, queue.first.value) {
		// New node value is smaller than first.
		newNode.next = queue.first
		queue.first = newNode
		return
	}
	var currentNode = queue.first
	for {
		if currentNode.next == nil {
			// Add new node here.
			currentNode.next = newNode
			return
		}
		if isLessThan(newNode.value, currentNode.next.value) {
			// Add new node in between.
			newNode.next = currentNode.next
			currentNode.next = newNode
			return
		}
		// Move to next node.
		currentNode = currentNode.next
	}
}

// This function add new value in queue in max order.
func (queue *QueueObject) enqueueMax(newNode *nodeObject) {
	if !isLessThan(newNode.value, queue.first.value) {
		// New node value is more than first.
		newNode.next = queue.first
		queue.first = newNode
		return
	}
	var currentNode = queue.first
	for {
		if currentNode.next == nil {
			// Add new node here.
			currentNode.next = newNode
			return
		}
		if !isLessThan(newNode.value, currentNode.next.value) {
			// Add new node in between.
			newNode.next = currentNode.next
			currentNode.next = newNode
			return
		}
		// Move to next node.
		currentNode = currentNode.next
	}
}

// This function add new value in queue in given order.
func (queue *QueueObject) EnqueueWith(value interface{}, compare func(valueA interface{}, valueB interface{}) bool) {
	var newNode = nodeObject{value: value}
	if queue.first == nil {
		// First and last points to new node.
		queue.first = &newNode
		queue.last = &newNode
		return
	}
	if compare(newNode.value, queue.first.value) {
		// Save value based on compare.
		newNode.next = queue.first
		queue.first = &newNode
		return
	}
	var currentNode = queue.first
	for {
		if currentNode.next == nil {
			// Add new node here.
			currentNode.next = &newNode
			return
		}
		if compare(newNode.value, currentNode.next.value) {
			// Add new node in between.
			newNode.next = currentNode.next
			currentNode.next = &newNode
			return
		}
		// Move to next node.
		currentNode = currentNode.next
	}
}

// This function return if less than.
func isLessThan(valueA interface{}, valueB interface{}) bool {
	switch valueAType := valueA.(type) {
	case int:
		// For b.
		switch valueBType := valueB.(type) {
		case int:
			if valueAType > valueBType {
				return false
			}
		case float64:
			if valueAType > int(valueBType) {
				return false
			}
		}
	case float64:
		// For b.
		switch valueBType := valueB.(type) {
		case int:
			if valueAType > float64(valueBType) {
				return false
			}
		case float64:
			if valueAType > valueBType {
				return false
			}
		}
	case string:
		// For b.
		switch valueBType := valueB.(type) {
		case string:
			if strings.Compare(valueAType, valueBType) > 0 {
				return false
			}
		}
	}
	return true
}

// This function remove the first node from queue.
func (queue *QueueObject) Dequeue() interface{} {
	if queue.first == queue.last {
		// Queue have single value.
		queue.last = nil
	}
	value := queue.first.value
	// First points to second node.
	queue.first = queue.first.next
	return value
}

// Dequeue int function remove int value from queue.
func (queue *QueueObject) DequeueInt() int {
	value := queue.Dequeue()
	switch valueType := value.(type) {
	case int:
		return valueType
	case float64:
		return int(valueType)
	}
	return 0
}

// Dequeue float64 function remove float64 value from queue.
func (queue *QueueObject) DequeueFloat64() float64 {
	value := queue.Dequeue()
	switch valueType := value.(type) {
	case int:
		return float64(valueType)
	case float64:
		return valueType
	}
	return 0.0
}

// Dequeue string function remove string value from queue.
func (queue *QueueObject) DequeueString() string {
	value := queue.Dequeue()
	switch valueType := value.(type) {
	case int:
		return strconv.Itoa(valueType)
	case float64:
		s := fmt.Sprintf("%f", valueType)
		return s
	case string:
		return valueType
	}
	return ""
}

// This function shows values of queue.
func (queue QueueObject) Print() {
	if queue.first == nil {
		// Empty queue.
		return
	}
	var node = queue.first
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

// This function create new queue.
func Create(priority int) QueueObject {
	return QueueObject{priority: priority}
}
