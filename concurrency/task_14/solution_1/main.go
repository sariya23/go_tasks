package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const numRequests = 10000

var count int

var m sync.Mutex

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	defer timer("main")()
	RealRequestWebGetCloseBodyWithoutRead()
}

func InitFunc() {
	var wg sync.WaitGroup

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)

}

func RealRequestMyPostNewClient() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			client := http.Client{
				Timeout: time.Second * 20,
				Transport: &http.Transport{
					ResponseHeaderTimeout: time.Second * 10,
				},
			}
			defer wg.Done()
			m.Lock()
			req, err := http.NewRequest(
				http.MethodPost,
				"http://127.0.0.1:8082/register",
				bytes.NewBuffer([]byte(fmt.Sprintf(`{"login": "test%v@gmail.com","password": "qwerty"}`, i))),
			)
			m.Unlock()
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestMyPost() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	client := http.Client{
		Timeout: time.Second * 20,
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 10,
		},
	}

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			req, err := http.NewRequest(
				http.MethodPost,
				"http://127.0.0.1:8082/register",
				bytes.NewBuffer([]byte(fmt.Sprintf(`{"login": "test%v@gmail.com","password": "qwerty"}`, i))),
			)
			m.Unlock()
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestMyPostBigTimeout() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	client := http.Client{
		Timeout: time.Second * 1000,
	}

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			req, err := http.NewRequest(
				http.MethodPost,
				"http://127.0.0.1:8082/register",
				bytes.NewBuffer([]byte(fmt.Sprintf(`{"login": "test%v@gmail.com","password": "qwerty"}`, i))),
			)
			m.Unlock()
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestMyGet() {
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
				http.MethodGet,
				"http://127.0.0.1:8082/login",
				bytes.NewBuffer([]byte(`{"login": "test148@gmail.com","password": "qwerty"}`)),
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestMyGetWithoutReadBody() {
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
				http.MethodGet,
				"http://127.0.0.1:8082/login",
				bytes.NewBuffer([]byte(`{"login": "test148@gmail.com","password": "qwerty"}`)),
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestWebGetLongTimeout() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	client := http.Client{
		Timeout: time.Second * 100,
	}

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(
				http.MethodGet,
				"https://catfact.ninja/fact",
				nil,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestWebGet() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	client := http.Client{
		Timeout: time.Second * 20,
	}

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(
				http.MethodGet,
				"https://catfact.ninja/fact",
				nil,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
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
	fmt.Println(count)
}

func RealRequestWebGetCloseBodyWithoutRead() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var count int64
	const numRequests = 10000

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			client := http.Client{
				Timeout: time.Second * 5,
			}
			defer wg.Done()
			req, err := http.NewRequest(
				http.MethodGet,
				"https://catfact.ninja/fact",
				nil,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
			}()
			m.Lock()
			count++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
