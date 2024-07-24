package main

import (
	"log"
	"sync"
	"time"
)

func worker(id int) {
	log.Printf("Worker %d iniciando\n", id)
	time.Sleep(time.Second)
	log.Printf("Worker %d terminou\n", id)
}

func teste() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}
	wg.Wait()
}
