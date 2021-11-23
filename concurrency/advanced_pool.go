package concurrency

import (
	"context"
	"errors"
	"log"
	"sync"
)

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.Context) error
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	currentWg := sync.WaitGroup{}
	currentWg.Add(maxConcurrent)
	currentPooler := PoolCollection{
		Tasks:         make(chan func(context.Context), maxSlots),
		Wg:            &currentWg,
		maxConcurrent: maxConcurrent,
		maxCount:      0,
		isClosed:      false,
	}
	return &currentPooler, nil
}

func (p *PoolCollection) Close(ctx context.Context) error {
	// check if context was cancelled
	select {
	case <-ctx.Done():
		return ErrPoolClosed
	default:
		p.Wg.Wait()
		p.isClosed = true
		close(p.Tasks)
	}
	return nil
}

func (p *PoolCollection) Submit(ctx context.Context, task func(context.Context)) error {

	// if pool is closed return error
	if p.isClosed {
		return ErrPoolClosed
	}
	// initial workers with maxconcurrent task if not already initialized
	if p.maxCount == 0 {
		for i := 0; i < p.maxConcurrent; i++ {
			go p.run(ctx)
		}
		p.maxCount = p.maxConcurrent
	}
	// Submit task, if full wait, if context is done before execute return context error.
	select {
	case <-ctx.Done():
		log.Println("context cancelled or done")
		return ctx.Err()
	case p.Tasks <- task:
		log.Println("slot became available; task submitted")
	default:
		log.Println("no slot available; task queued")
		<-p.Tasks
		p.Tasks <- task
		log.Println("slot became available; task submitted")
	}

	return nil
}

func (p *PoolCollection) run(ctx context.Context) {
	defer p.Wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-p.Tasks:
			task(ctx)
		}
	}
}
