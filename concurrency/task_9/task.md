===========================================================
# Задача 9

Ответить на вопросы:

1. Что выведется и как исправить?
2. Что поправить, чтобы сохранить порядок?

===========================================================

```go
func main() {
  m := make(char string, 3)
  cnt := 5
  for i := 0; i < cnt; i++ {
    go func() {
      m <- fmt.Sprintf("Goroutine %d", i)
    }()
  }
  for i := 0; i < cnt; i++ {
    go ReceiveFromCh(m)
  }
}
func ReceiveFromCh(ch chan string) {
  fmt.Println(<-ch)
}
```