<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Anonymous Functions in Go</title>
<style>
body{
  font-family: Arial, sans-serif;
  line-height:1.6;
  margin:40px;
  background:#f7f7f7;
}
h1,h2{
  color:#333;
}
pre{
  background:#1e1e1e;
  color:#eaeaea;
  padding:15px;
  border-radius:6px;
  overflow-x:auto;
}
code{
  color:#9cdcfe;
}
.section{
  background:white;
  padding:20px;
  border-radius:8px;
  margin-bottom:20px;
  box-shadow:0 2px 6px rgba(0,0,0,0.08);
}
</style>
</head>

<body>

<h1>Anonymous Functions in Go</h1>

<div class="section">
<h2>What is an Anonymous Function?</h2>
<p>
An <strong>anonymous function</strong> is a function that does not have a name.
It is usually used when the function is needed only once or when we want to
create a quick <strong>closure</strong>.
</p>

<p>
In <strong>Go</strong>, functions can be passed as parameters to other functions.
</p>
</div>

<div class="section">
<h2>Example Function That Accepts Another Function</h2>

<p>This function takes another function (<code>converter</code>) and applies it to three numbers.</p>

<pre><code>func conversions(converter func(int) int, x, y, z int) (int, int, int) {
    convertedX := converter(x)
    convertedY := converter(y)
    convertedZ := converter(z)

    return convertedX, convertedY, convertedZ
}</code></pre>

</div>

<div class="section">
<h2>Using a Named Function</h2>

<p>First we create a normal function and pass it to <code>conversions</code>.</p>

<pre><code>func double(a int) int {
    return a + a
}

func main() {
    newX, newY, newZ := conversions(double, 1, 2, 3)
}
</code></pre>

<p>Results:</p>

<ul>
<li>newX = 2</li>
<li>newY = 4</li>
<li>newZ = 6</li>
</ul>

</div>

<div class="section">
<h2>Using an Anonymous Function</h2>

<p>
Instead of creating a separate function, we can define the function directly
inside the call.
</p>

<pre><code>func main() {

    newX, newY, newZ := conversions(
        func(a int) int {
            return a + a
        },
        1, 2, 3,
    )

}
</code></pre>

<p>
This function has <strong>no name</strong>, so it is called an anonymous function.
</p>

</div>

<div class="section">
<h2>Why Use Anonymous Functions?</h2>

<ul>
<li>Useful for small one-time functions</li>
<li>Keeps code shorter</li>
<li>Good for functional patterns</li>
<li>Can capture variables from the surrounding scope (closures)</li>
</ul>

</div>

</body>
</html>