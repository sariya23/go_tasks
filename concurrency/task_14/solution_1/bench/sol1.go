package solution

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func InitFunc() int64 {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			m.Lock()
			count++
			m.Unlock()
		}()
	}

	wg.Wait()
	return count
}

func AtomicFunc() int64 {
	var wg sync.WaitGroup
	var count int64
	const numRequests = 10000

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	return count
}

func SemaFunc() int64 {
	var wg sync.WaitGroup
	var count int64
	const numRequests = 10000
	sema := make(chan struct{}, 1)

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			sema <- struct{}{}
			count++
			<-sema
		}()
	}

	wg.Wait()
	return count
}

func CriticalSectionFunc() int64 {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			m.Lock()
			defer wg.Done()
			time.Sleep(time.Millisecond)
			count++
			m.Unlock()
		}()
	}

	wg.Wait()
	return count
}

func RealRequest() int64 {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	client := http.Client{
		Timeout: time.Second * 5,
	}

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(
				http.MethodPost,
				"http://127.0.0.1:8082/register",
				bytes.NewBuffer([]byte(fmt.Sprintf(`{"login": "test%v@gmail.com","password": "qwerty"}`, i))),
			)
			if err != nil {
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					return
				}
			}()
			_, _ = io.Copy(io.Discard, resp.Body)
			m.Lock()
			count++
			m.Unlock()
		}()
	}

	wg.Wait()
	return count
}
