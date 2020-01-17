/*
Describe a scenario by code:
1. main thread read data from DB (fake DBdata is enough) and put the data in cache, refesh the data every 3s
2. other child threads read or write the cache
3. make sure that child threads doesn't mess the data during the write/read
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type fakeData struct {
	Name string
	Age  uint8
}

var fakeD fakeData = fakeData{Name: "Mike", Age: 9}

func reader(ctx context.Context, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			func(lock *sync.RWMutex) {
				lock.RLock()
				defer lock.RUnlock()
				fmt.Printf("Reader: %+v\n", fakeD)
			}(lock)
		case <-ctx.Done():
			return
		}
	}

}

func writer(ctx context.Context, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			func(lock *sync.RWMutex) {
				lock.Lock()
				defer lock.Unlock()
				fakeD.Age++
				fmt.Printf("Writer: %+v\n", fakeD)
			}(lock)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	lock := &sync.RWMutex{}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, lock *sync.RWMutex) {
		ticker := time.NewTicker(3 * time.Second)
		defer wg.Done()
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				func(lock *sync.RWMutex) {
					lock.Lock()
					defer lock.Unlock()
					fakeD = fakeData{Name: "Bob", Age: 20}
					fmt.Println("Fetch data from fake DB...")
				}(lock)
			case <-ctx.Done():
				return
			}
		}
	}(ctx, wg, lock)

	for index := 0; index < 10; index++ {
		wg.Add(1)
		go reader(ctx, wg, lock)
	}

	for index := 0; index < 10; index++ {
		wg.Add(1)
		go writer(ctx, wg, lock)
	}

	<-quit
	cancelFunc()
	wg.Wait()

}
