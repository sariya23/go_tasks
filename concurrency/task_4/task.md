===========================================================
Задача 4
1. Как это работает, что не так, что поправить?
===========================================================

func main() {
  ch := make(chan bool)
  ch <- true
  go func() {
    <-ch
  }()
  ch <-true
}