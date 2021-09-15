//github: https://github.com/Special23/queue.git

package queue

import (
	"time"
)

const (
	DATA_NOT_EXISTS = 0 //无数据
	DATA_EXISTS     = 1 //存在数据
)

const (
	QUEUE_DEFAULT_SIZE = 1024 //队列默认容量
)

type T interface{}

type Queue struct {
	size  uint32  //容量
	queue []T     //队列数据
	sign  []uint8 //队列标记
	ri    uint32  //写入索引
	wi    uint32  //读出索引
	rc    uint32  //写入计数
	wc    uint32  //读出计数
}

func NewQueue(s uint32) *Queue {
	if s == 0 {
		s = QUEUE_DEFAULT_SIZE
	}
	return &Queue{
		size:  s,
		queue: make([]T, s),
		sign:  make([]uint8, s),
		ri:    0,
		wi:    0,
		rc:    0,
		wc:    0,
	}
}

//入队
//注意：一直阻塞到入成功为止
func (m *Queue) Push(src T) {
	sleepTime := time.Duration(1) * time.Second
	for m.sign[m.wi] != DATA_NOT_EXISTS {
		time.Sleep(sleepTime)
	}

	m.queue[m.wi] = src
	m.sign[m.wi] = DATA_EXISTS
	m.wi++
	m.wc++
	if m.wi >= m.size {
		m.wi = 0
	}
}

//尝试入队
//注意：成功返回true，失败为false
func (m *Queue) TryPush(src T) bool {
	if m.sign[m.wi] != DATA_NOT_EXISTS {
		return false
	}

	m.queue[m.wi] = src
	m.sign[m.wi] = DATA_EXISTS
	m.wi++
	m.wc++
	if m.wi >= m.size {
		m.wi = 0
	}
	return true
}

//出队
//注意：会一直到成功出队为止
func (m *Queue) Pull() T {
	sleepTime := time.Duration(1) * time.Second
	for m.sign[m.ri] != DATA_EXISTS {
		time.Sleep(sleepTime)
	}

	dst := m.queue[m.ri]
	m.sign[m.ri] = DATA_NOT_EXISTS
	m.rc++
	m.ri++
	if m.ri >= m.size {
		m.ri = 0
	}
	return dst
}

//尝试出队
//注意：成功返回对应元素，失败返回nil
func (m *Queue) TryPull() T {
	if m.sign[m.ri] != DATA_EXISTS {
		return nil
	}

	dst := m.queue[m.ri]
	m.sign[m.ri] = DATA_NOT_EXISTS
	m.rc++
	m.ri++
	if m.ri >= m.size {
		m.ri = 0
	}
	return dst
}

//队列长度
func (m *Queue) Len() int {
	return int(m.wc - m.rc)
}

//是否为空
func (m *Queue) IsEmpty() bool {
	return m.rc == m.wc
}
