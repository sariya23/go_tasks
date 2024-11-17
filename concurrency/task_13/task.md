===========================================================
Задача 13
1. Запросить параллельно данные из источников. Если все где-то произошла ошибка, то вернуть ошибку, иначе вернуть nil.
2. Представим, что теперь функция должна возвращать результат int. Есть функция resp.Size(), для каждого url надо проссумировать и вернуть, если ошибок не было. Просто описать подход к решению
3. Что делать, если урлов у нас миллионы?

===========================================================

package main

func main() {
    _, err := download([]string{
        "https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
        "https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
        "https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
        "https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
        "https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
    })

    if err != nil {
        panic(err)
    }
}

func download(urls []string) (error) {
    return nil
}
