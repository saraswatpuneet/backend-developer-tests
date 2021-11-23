package concurrency

import (
	"context"
	"sync"
)

// PoolCollection is a collection of worker pools that manages slots of workers running x concurrent tasks.
type PoolCollection struct {
	Tasks         chan func(context.Context)
	Wg            *sync.WaitGroup
	maxConcurrent int
	jobCount      int
}
