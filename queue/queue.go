package main

import "fmt"

type Queue struct {
	data []int
}

var index = 0

func (q *Queue) enque(n int) {
	q.data[index] = n
	index++
}

func (q *Queue) front() {
	fmt.Println(q.data[0])
}

func (q *Queue) deque() {
	q.data = q.data[1:]
}

func (q *Queue) rear() {
	fmt.Println(q.data[len(q.data)-1])
}

func main() {
	queue := Queue{}
	queue.enque(1)
	queue.enque(2)
	queue.enque(3)
	queue.enque(4)
	queue.enque(5)
	queue.deque()
	fmt.Println(queue)
	queue.front()
	queue.rear()
}
