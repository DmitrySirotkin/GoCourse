package main

import "fmt"

var queue []string
var queueCount int = -1

func main() {
	queuePush("First")
	queuePush("Second")
	queuePush("Thrid")

	fmt.Println(queuePop())
	fmt.Println(queuePop())
	fmt.Println(queuePop())
}

func queuePush(str string) {
	queue = append(queue, str)
}

func queuePop() string {
	if queueCount > len(queue) {
		return ""
	}
	queueCount = queueCount + 1
	return queue[queueCount]
}
