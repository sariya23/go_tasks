===========================================================
# Задача 14

1. Что выведет на экран и сколько времени будет работать?
2. Нужно ускорить, чтобы работало быстрее. Сколько будет работать теперь?
3. Если бы в networkRequest выполнялся реальный сетевой вызов, то какие с какими проблемами мы могли бы столкнуться в данном коде?
4. Если url немного, а запросов к ним много, то как можно оптимизировать?

===========================================================

```go
package main

import (
    "fmt"
    "time"
)

const numRequests = 10000

var count int

var m sync.Mutex

func networkRequest() {
    time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
    m.Lock()
    count++
    m.Unlock()
}

func main() {
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
```