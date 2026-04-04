<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go fmt Package Cheatsheet</title>
<style>
    body { font-family: Arial, sans-serif; line-height: 1.5; margin: 20px; background: #f9f9f9; }
    h1, h2 { color: #2c3e50; }
    code { background: #ecf0f1; padding: 2px 4px; border-radius: 3px; }
    pre { background: #ecf0f1; padding: 10px; border-radius: 5px; }
    table { border-collapse: collapse; width: 100%; margin-top: 10px; }
    th, td { border: 1px solid #bdc3c7; padding: 8px; text-align: left; }
    th { background: #34495e; color: white; }
</style>
</head>
<body>

<h1>Go <code>fmt</code> Package Cheatsheet</h1>

<h2>Print Functions</h2>
<table>
<tr><th>Function</th><th>Usage</th><th>Example</th></tr>
<tr>
<td><code>fmt.Print()</code></td>
<td>Prints values to stdout without newline</td>
<td><pre>fmt.Print("Hello ", "world") // Hello world</pre></td>
</tr>
<tr>
<td><code>fmt.Println()</code></td>
<td>Prints values to stdout with space between args and newline</td>
<td><pre>fmt.Println("Hello", "world") // Hello world\n</pre></td>
</tr>
<tr>
<td><code>fmt.Printf()</code></td>
<td>Prints formatted string using verbs</td>
<td><pre>name := "Alice"; age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Name: Alice, Age: 30</pre></td>
</tr>
<tr>
<td><code>fmt.Sprintf()</code></td>
<td>Formats string and returns it (does not print)</td>
<td><pre>msg := fmt.Sprintf("Hello %s!", "Bob")
fmt.Println(msg) // Hello Bob!</pre></td>
</tr>
</table>

<h2>Scan Functions</h2>
<table>
<tr><th>Function</th><th>Usage</th><th>Example</th></tr>
<tr>
<td><code>fmt.Scan()</code></td>
<td>Reads space-separated values from stdin</td>
<td><pre>var name string; var age int
fmt.Print("Enter name and age: ")
fmt.Scan(&name, &age)
fmt.Println(name, age)</pre></td>
</tr>
<tr>
<td><code>fmt.Scanln()</code></td>
<td>Reads values from stdin until newline</td>
<td><pre>var name string; var age int
fmt.Print("Enter name and age: ")
fmt.Scanln(&name, &age)
fmt.Println(name, age)</pre></td>
</tr>
<tr>
<td><code>fmt.Sscanf()</code></td>
<td>Reads formatted values from string</td>
<td><pre>input := "Alice 30"
var name string; var age int
fmt.Sscanf(input, "%s %d", &name, &age)
fmt.Println(name, age) // Alice 30</pre></td>
</tr>
<tr>
<td><code>fmt.Fprint / Fprintln / Fprintf</code></td>
<td>Write formatted output to any <code>io.Writer</code></td>
<td><pre>fmt.Fprintf(os.Stdout, "Hello %s!\n", "world")</pre></td>
</tr>
</table>

<h2>Common Format Verbs</h2>
<table>
<tr><th>Verb</th><th>Description</th><th>Example</th></tr>
<tr><td>%v</td><td>Default format</td><td>fmt.Printf("%v", 42)</td></tr>
<tr><td>%+v</td><td>Struct with field names</td><td>fmt.Printf("%+v", User{"Alice", 30})</td></tr>
<tr><td>%#v</td><td>Go-syntax representation</td><td>fmt.Printf("%#v", 42)</td></tr>
<tr><td>%T</td><td>Type of value</td><td>fmt.Printf("%T", 42)</td></tr>
<tr><td>%d</td><td>Integer</td><td>fmt.Printf("%d", 42)</td></tr>
<tr><td>%f</td><td>Float</td><td>fmt.Printf("%f", 3.14)</td></tr>
<tr><td>%s</td><td>String</td><td>fmt.Printf("%s", "hello")</td></tr>
<tr><td>%q</td><td>Quoted string</td><td>fmt.Printf("%q", "hello")</td></tr>
<tr><td>%x / %X</td><td>Hexadecimal</td><td>fmt.Printf("%x", 255) // ff</td></tr>
<tr><td>%p</td><td>Pointer</td><td>fmt.Printf("%p", &x)</td></tr>
</table>

</body>
</html>