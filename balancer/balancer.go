package balancer

import (
	"container/heap"
)

type Pool []*Worker

func (p Pool) Len() int { return len(p) }
func (p Pool) Less(i, j int) bool {
	return p[i].Pending < p[j].Pending
}
func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].Index = i
	p[j].Index = j
}
func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*Worker)
	item.Index = n
	*p = append(*p, item)
}
func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*p = old[0 : n-1]
	return item
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func NewBalancer(workers []*Worker, done chan *Worker) *Balancer {
	b := &Balancer{
		pool: make(Pool, len(workers)),
		done: done,
	}
	for i, w := range workers {
		b.pool[i] = w
	}
	heap.Init(&b.pool)
	return b
}

func (b *Balancer) Balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.Requests <- req
	// One more in its work queue.
	w.Pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.Pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.Index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}
