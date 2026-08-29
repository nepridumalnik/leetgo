package workers

import (
	"errors"
	"sync"
)

type Task func()

type WorkerPool struct {
	workers   int8
	tasksC    chan Task
	isClosed  bool
	taskGuard sync.RWMutex
}

func New(workers int8) (*WorkerPool, error) {
	if workers == 0 {
		return nil, errors.New("workers number can not be zero")
	}

	return &WorkerPool{
		workers:   workers,
		tasksC:    make(chan Task, workers*5),
		isClosed:  false,
		taskGuard: sync.RWMutex{},
	}, nil
}

func (w *WorkerPool) Run() {
	wg := sync.WaitGroup{}
	for range w.workers {
		wg.Go(w.loop)
	}

	wg.Wait()
}

func (w *WorkerPool) Close() {
	w.taskGuard.Lock()
	defer w.taskGuard.Unlock()

	w.isClosed = true
	close(w.tasksC)
}

func (w *WorkerPool) Submit(task Task) error {
	w.taskGuard.RLock()
	defer w.taskGuard.RUnlock()

	if w.isClosed {
		return errors.New("pool closed")
	}

	w.tasksC <- task
	return nil
}

func (w *WorkerPool) loop() {
	for task := range w.tasksC {
		task()
	}
}
