package stackandqueue

type Queue struct {
	Front int
	Rear  int
	Data  []int
}

func NewQueue(data []int) Queue {
	return Queue{
		Front: len(data) - 1,
		Rear:  0,
		Data:  data,
	}
}

func (q *Queue) Enqueue(data int) {
	q.Data = append([]int{data}, q.Data...)
	q.Front += 1
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	data := q.Data[q.Front]
	q.Data = q.Data[:q.Front]
	q.Front -= 1
	return data
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.Data[q.Front]
}

func (q *Queue) IsEmpty() bool {
	return q.Front == -1
}

type StaticQueue struct {
	Front   int
	Rear    int
	Data    []int
	MaxSize int
}

func NewStaticQueue(maxSize int, data []int) StaticQueue {
	if len(data) > maxSize {
		panic("data length is greater than the queue max size value")
	}
	return StaticQueue{
		Front:   0,
		Rear:    len(data) - 1,
		Data:    data,
		MaxSize: maxSize,
	}
}

func (q *StaticQueue) Enqueue(data int) {
	if q.IsFull() {
		panic("queue is full")
	}
	q.Data = append(q.Data, data)
	q.Rear += 1
}

func (q *StaticQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	data := q.Data[0]
	q.Data = q.Data[1:]
	q.Rear -= 1
	return data
}

func (q *StaticQueue) Peek() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.Data[0]
}

func (q *StaticQueue) IsFull() bool {
	return q.Rear == q.MaxSize-1
}

func (q *StaticQueue) IsEmpty() bool {
	return q.Rear == -1
}
