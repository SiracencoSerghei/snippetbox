<!DOCTYPE html>
<html lang="uk">
<head>
<meta charset="UTF-8">
<title>Go Context Example</title>
<style>
    body {
        font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
        line-height: 1.6;
        margin: 20px;
        background-color: #f5f5f5;
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
    h1, h2 {
        color: #333;
    }
    .note {
        background: #ffeeba;
        padding: 10px;
        border-left: 5px solid #ffc107;
        margin: 10px 0;
    }
</style>
</head>
<body>

<h1>Go <code>context</code> Пояснення</h1>

<p><strong>Context</strong> — це спосіб передавати сигнал відміни, таймаут або дедлайн між різними частинами програми, наприклад, між HTTP handler, базою даних і зовнішніми API.</p>

<h2>Основні види контексту:</h2>
<ul>
    <li><code>context.Background()</code> — базовий, порожній контекст, який нічого не відміняє.</li>
    <li><code>context.TODO()</code> — для місць, де ще не визначено який контекст використовувати.</li>
    <li><code>context.WithCancel(parent)</code> — створює контекст з можливістю ручного скасування.</li>
    <li><code>context.WithTimeout(parent, duration)</code> — створює контекст, який автоматично скасовується через <em>duration</em>.</li>
    <li><code>context.WithDeadline(parent, time)</code> — як <code>WithTimeout</code>, але з конкретним дедлайном.</li>
</ul>

<div class="note">
Важливо: <strong>передавати контекст як перший аргумент</strong> у функції, які можуть блокуватися (DB, HTTP, API calls), щоб уникнути зависання goroutine.
</div>

<h2>Простий приклад з <code>context.WithTimeout</code>:</h2>

<pre><code>package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Створюємо контекст з таймаутом 2 секунди
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // завжди викликаємо cancel, щоб уникнути утечок

    // Імітуємо довгу роботу
    done := make(chan string)
    go func() {
        time.Sleep(3 * time.Second) // довга операція
        done &lt;- "готово"
    }()

    // Очікуємо результат або таймаут
    select {
    case res := &lt;-done:
        fmt.Println("Результат:", res)
    case &lt;-ctx.Done():
        fmt.Println("Скасовано або таймаут:", ctx.Err())
    }
}
</code></pre>

<p>У цьому прикладі:</p>
<ul>
    <li>Горутина спить 3 секунди.</li>
    <li>Контекст таймаутить через 2 секунди.</li>
    <li>В результаті отримаємо: <code>Скасовано або таймаут: context deadline exceeded</code>.</li>
</ul>

<h2>Використання у HTTP handler:</h2>

<pre><code>func handler(w http.ResponseWriter, r *http.Request) {
    // Беремо контекст з HTTP request
    ctx := r.Context()

    result, err := slowQuery(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusRequestTimeout)
        return
    }

    fmt.Fprintf(w, "Результат: %v", result)
}
</code></pre>

<p>Якщо клієнт закриває з'єднання, <code>ctx.Done()</code> сигналізує про це, і запит можна зупинити.</p>

<h2>Коротко</h2>
<ul>
    <li>Контекст потрібен для керування життєвим циклом операцій.</li>
    <li>Завжди передавати його у функції, що блокуються.</li>
    <li>Не можна зберігати контекст у структурі на довго.</li>
    <li>Always call <code>cancel()</code> для контекстів з таймаутом або скасуванням.</li>
</ul>

</body>
</html>