# Создание очереди

Нужно реализовать тип `Queue` - очередь FIFO фиксированного размера на `n` элементов. Поддерживает 2 операции:

```
Get(block bool) (int, error)
Put(val int, block bool) error
```

`Put`:
- Если в очереди есть место — помещает `val` в очередь и возвращает `nil`;
- Если очередь заполнена **и** `block = true` — блокируется, пока не освободится место. Затем помещает `val` в очередь и возвращает `nil`;
- Если очередь заполнена **и** `block = false` — возвращает ошибку `ErrFull`.

`Get`:
- Если в очереди есть значения — выбирает очередное значение, возвращает его и `nil`;
- Если очередь пуста **и** `block = true` — блокируется, пока не появится значение. Затем выбирает и возвращает его и `nil`;
- Если очередь пуста и `block = false` — возвращает нулевое значение и ошибку `ErrEmpty`.

Решение не должно использовать `len`, `cap`, счетчики длины очереди, цикли для подсчета длины очереди. Также не нужно менять уже написанный код - если метод принимает структуру по значению, то не нужно менять тип ресивера на указательный. 

## Код для проверки

Нужно реализовать только ту часть, **что между "Начало решения" и "Конец решения"**.

```go
package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	// ...
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	// ...
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	// ...
}

// конец решения

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
```