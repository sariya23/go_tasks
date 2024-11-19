===========================================================
# Задача 6

Ответить на вопросы:

1. Что выведет код и как исправить?

===========================================================
```go
var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
var a = 0
 
func main() {
    wg := sync.WaitGroup{}
    wg.Add(3)
    go func() {
        wg.Done()
        a=10
        globalMap["test"] = append(globalMap["test"], a)
         
    }()
    go func() {
        wg.Done()
        a=11
        globalMap["test2"] = append(globalMap["test2"], a)
    }()
    go func() {
        wg.Done()
        a=12
        globalMap["test3"] = append(globalMap["test3"], a)
    }()
    wg.Wait()
    fmt.Printf("%v", globalMap)
    fmt.Printf("%d", a)
}
```