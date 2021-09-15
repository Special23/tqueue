package queue

import (
	"fmt"
	"sync"
	"testing"
)

var testQueue *Queue

func TestQueue(t *testing.T) {
	//test normal
	testQueue = NewQueue(3)
	testQueue.Push(1)
	testQueue.Push(2)
	fmt.Printf("testQueue len:%d \n", testQueue.Len())

	ret3 := testQueue.TryPush(3)
	fmt.Printf("testQueue len:%d TryPush ret:%t \n", testQueue.Len(), ret3)
	ret4 := testQueue.TryPush(4)
	fmt.Printf("testQueue len:%d TryPush ret:%t \n", testQueue.Len(), ret4)

	v1 := testQueue.Pull()
	fmt.Printf("testQueue len:%d v1:%v \n", testQueue.Len(), v1)
	v2 := testQueue.Pull()
	fmt.Printf("testQueue len:%d v2:%v \n", testQueue.Len(), v2)

	v3 := testQueue.TryPull()
	fmt.Printf("testQueue len:%d v3:%v \n", testQueue.Len(), v3)
	v4 := testQueue.TryPull()
	fmt.Printf("testQueue len:%d v4:%v \n", testQueue.Len(), v4)

	//test muti-thread
	testQueue = NewQueue(100)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 200; i++ {
			testQueue.Push(1)
			fmt.Printf("testQueue TryPush len:%d \n", testQueue.Len())
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 200; i++ {
			ret := testQueue.Pull()
			fmt.Printf("testQueue TryPull len:%d ret:%v \n", testQueue.Len(), ret)
		}
		wg.Done()
	}()
	wg.Wait()
}
