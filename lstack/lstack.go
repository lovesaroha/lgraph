/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lstack

import (
	"fmt"
	"strconv"
)

// Stack object defined.
type StackObject struct {
	first *nodeObject
}

// Node object structure.
type nodeObject struct {
	value interface{}
	next  *nodeObject
}

// This function check if stack is empty.
func (stack StackObject) IsEmpty() bool {
	return stack.first == nil
}

// Push function to add value in stack.
func (stack *StackObject) Push(value interface{}) {
	var newNode = nodeObject{next: stack.first, value: value}
	stack.first = &newNode
}

// Push values function add multiple values in stack.
func (stack *StackObject) PushValues(values []interface{}) {
	for _, value := range values {
		stack.Push(value)
	}
}

// Pop function remove value from stack.
func (stack *StackObject) Pop() interface{} {
	if stack.first == nil {
		// Empty stack.
		return nil
	}
	value := stack.first.value
	stack.first = stack.first.next
	return value
}

// Pop int function remove int value from stack.
func (stack *StackObject) PopInt() int {
	value := stack.Pop()
	switch valueType := value.(type) {
	case int:
		return valueType
	case float64:
		return int(valueType)
	}
	return 0
}

// Pop float64 function remove float64 value from stack.
func (stack *StackObject) PopFloat64() float64 {
	value := stack.Pop()
	switch valueType := value.(type) {
	case int:
		return float64(valueType)
	case float64:
		return valueType
	}
	return 0.0
}

// Pop string function remove string value from stack.
func (stack *StackObject) PopString() string {
	value := stack.Pop()
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

// This function shows values of stack.
func (stack StackObject) Print() {
	var node = stack.first
	if node == nil {
		// Empty stack.
		return
	}
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

// This function create new stack.
func Create() StackObject {
	return StackObject{}
}
