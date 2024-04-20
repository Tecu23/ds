package main

import (
	"fmt"

	"github.com/Tecu23/ds/queue"
)

func main() {

	q := queue.New()

	q.Enqueue("a")
	q.Enqueue(2)

	for q.Len() > 0 {
		v := q.Dequeue()
		fmt.Println(v)
	}

}
