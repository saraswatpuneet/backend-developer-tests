# Concurrency with worker pools

## SimplePool example
1. Implemented a worker pool with a buffered channel queue of specified size and buffer takes in function reference.
2. As tasks are submit so queue fills up and wait for worker to be available before accepting new task.
3. To test run:
    go run simple_main/simple_main.go: will start a simple worker pool with a task to send echo functions to execute by the pool workers. Check files for more details

## AdvancedPool example
1. Implemented a worker pool with a buffered channel queue of for total tasks running or waiting at a given point with only 2 maxConcurrent workers being initialized, maxConcurrent queues in waitgroup as many wait groups are created. 
2. As tasks are submit so queue fills up and wait for worker to be available before accepting new task.
3. To test run:
    go run advanced_main/advance_main.go: will start a simple worker pool with a task to send echo functions to execute by the pool workers. Check files for more details