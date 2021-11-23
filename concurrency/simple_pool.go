// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import "log"

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
	Execute()
}

type SimpleWorkerPool struct {
	maxConcurrent int
	queue         chan func()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	return &SimpleWorkerPool{
		maxConcurrent: maxConcurrent,
		queue:         make(chan func(), maxConcurrent),
	}
}

// Submit queues up the given task to be executed asynchronously via bufffered channel.
func (p *SimpleWorkerPool) Submit(task func()) {
	p.queue <- task
}

// Execute executes all queued tasks.
func (p *SimpleWorkerPool) Execute() {
	for i := 0; i < p.maxConcurrent; i++ {
		go func(i int) {
			for task := range p.queue {
				log.Printf("Started processing task on worker %v", i)
				task()
				log.Printf("Finished processing task on worker %v", i)
			}
		}(i)
	}
}
