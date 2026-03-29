<!DOCTYPE html>
<html lang="uk">
<head>
<meta charset="UTF-8">
<title>Goroutine в Go</title>
<style>
    body {
        font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
        background-color: #f5f5f5;
        line-height: 1.6;
        margin: 20px;
        color: #333;
    }
    h1, h2, h3 {
        color: #222;
    }
    pre {
        background: #272822;
        color: #f8f8f2;
        padding: 15px;
        border-radius: 5px;
        overflow-x: auto;
    }
    code {
        font-family: "Courier New", Courier, monospace;
    }
    .note {
        background: #ffeeba;
        padding: 10px;
        border-left: 5px solid #ffc107;
        margin: 10px 0;
    }
    ul {
        margin: 5px 0 15px 20px;
    }
</style>
</head>
<body>

<h1>Goroutine в Go</h1>

<p>Goroutine — це легка нитка виконання в Go, яка керується рантаймом Go і є дешевшою за OS thread.</p>

<h2>Основи</h2>
<ul>
    <li>Запускається за допомогою ключового слова <code>go</code>.</li>
    <li>Може працювати паралельно з іншими goroutine.</li>
    <li>Необхідно синхронізувати доступ до спільних даних (mutex, channel).</li>
</ul>

<h2>Простий приклад</h2>
<pre><code>package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello() // запуск goroutine
    time.Sleep(1 * time.Second) // чекаємо, щоб goroutine встиг виконатись
    fmt.Println("Main finished")
}
</code></pre>

<div class="note">
<strong>Пам’ятай:</strong> goroutine може не встигнути виконатись, якщо main() завершиться раніше.
</div>

<h2>Анонімні goroutine</h2>
<pre><code>go func(msg string) {
    fmt.Println(msg)
}("Привіт, світ!")</code></pre>

<h2>Передача параметрів</h2>
<ul>
    <li>Аргументи передаються копією, щоб уникнути race condition.</li>
    <li>Якщо передавати змінні з зовнішнього scope — треба копіювати їх у локальні змінні.</li>
</ul>

<h2>Синхронізація</h2>

<h3>1. Channel</h3>
<pre><code>ch := make(chan int)

go func() {
    ch &lt;- 42 // відправка в канал
}()

val := &lt;-ch // отримання з каналу
fmt.Println(val)</code></pre>

<h3>2. WaitGroup</h3>
<pre><code>import "sync"

var wg sync.WaitGroup
wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("First goroutine")
}()

go func() {
    defer wg.Done()
    fmt.Println("Second goroutine")
}()

wg.Wait() // чекаємо всі goroutine
</code></pre>

<h2>Timeout та контекст</h2>
<pre><code>import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

go func(ctx context.Context) {
    select {
    case &lt;-time.After(3 * time.Second):
        fmt.Println("Finished work")
    case &lt;-ctx.Done():
        fmt.Println("Timeout:", ctx.Err())
    }
}(ctx)
</code></pre>

<h2>Важливі нюанси</h2>
<ul>
    <li>Goroutine дешевші за потоки ОС, можна запускати тисячі.</li>
    <li>Використовувати <code>defer wg.Done()</code> у goroutine для WaitGroup.</li>
    <li>Ніколи не залишати канал відкритим без закриття або споживача — може бути deadlock.</li>
    <li>При довгих операціях завжди думати про cancel через <code>context</code>.</li>
    <li>Race condition: перевіряємо <code>go run -race</code>.</li>
</ul>

<h2>Коротко</h2>
<ul>
    <li><code>go func()</code> — запускає легку паралельну нитку.</li>
    <li>Channel — безпечний спосіб передавати дані між goroutine.</li>
    <li>WaitGroup — чекаємо завершення групи goroutine.</li>
    <li>Context — керує життєвим циклом та таймаутами.</li>
</ul>

</body>
</html>