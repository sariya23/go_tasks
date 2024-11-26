# Превратить синхронный код в потоковую обработку

В файле `./src.go` лежит исходный код. 

## Что делает код

Есть слайс из `50_00_000` элементов типа `Job`. Есть три этапа обработки этих данных. Визуально можно представить так:
Поля структуры перечислены в порядке их объявления.
```
                                  first stage                                 second stage
[Job{0, 0}, Job{1, 0}, Job{2, 0}] ->  [Job{0, 1}, Job{3.14, 1}, Job{6.28, 1}] -> [Job{0, 2}, Job{3.14 * math.E, 2}, Job{6.28 * math.E, 2}] ->
last stage
-> [Job{0, 3}, Job{int64(float64(3.14 * math.E) / float64(rand.Intn(10))), 3}, Job{(float646.28 * math.E) / float64(rand.Intn(10)), 3}]
```

## Что нужно сделать

Нам не нужны все `50_000_000` значений сразу, но текущий код пока не посчитает все, не успокоится. Нужно как-то сделать так, чтобы можно было получать значения последовательно, одно за другим.