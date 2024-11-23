===========================================================

# Задача 5

Ответить на вопросы:

1. Как будет работать код?
2. Как сделать так, чтобы выводился только первый ch?

===========================================================
```go
func main() {
        ch := make(chan bool)
        ch2 := make(chan bool)
        ch3 := make(chan bool)
        go func() {
                ch <- true
        }()
        go func() {
                ch2 <- true
        }()
        go func() {
                ch3 <- true
        }()

        select {
        case <-ch:
                fmt.Printf("val from ch")
        case <-ch2:
                fmt.Printf("val from ch2")
        case <-ch3:
                fmt.Printf("val from ch3")
        }
}
```