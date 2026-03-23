<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Go Testing Guide</title>

<style>
    body {
        font-family: Arial, sans-serif;
        line-height: 1.6;
        margin: 40px;
        background-color: #0f172a;
        color: #e2e8f0;
    }

    h1, h2, h3 {
        color: #38bdf8;
    }

    code {
        background-color: #1e293b;
        padding: 2px 6px;
        border-radius: 4px;
        color: #f8fafc;
    }

    pre {
        background-color: #020617;
        padding: 15px;
        border-radius: 8px;
        overflow-x: auto;
        border: 1px solid #334155;
    }

    blockquote {
        border-left: 4px solid #38bdf8;
        padding-left: 15px;
        color: #94a3b8;
        margin-left: 0;
    }

    table {
        border-collapse: collapse;
        width: 100%;
        margin-top: 10px;
    }

    th, td {
        border: 1px solid #334155;
        padding: 8px;
        text-align: left;
    }

    th {
        background-color: #1e293b;
    }

    .section {
        margin-bottom: 40px;
    }
</style>
</head>

<body>

<h1>🧪 Testing Guide (Go)</h1>

<div class="section">
<h2>📌 Philosophy</h2>
<p>In Go, tests are written and executed <strong>per package</strong>, not per file.</p>
<blockquote>
File-level testing is an anti-pattern. Use naming and filtering instead.
</blockquote>
</div>

<div class="section">
<h2>🚀 Run Tests</h2>

<h3>Run all tests in project</h3>
<pre><code>go test ./...</code></pre>

<h3>Run tests in a specific package</h3>
<pre><code>go test ./examples -v</code></pre>
</div>

<div class="section">
<h2>🎯 Run Specific Tests</h2>

<h3>By test name</h3>
<pre><code>go test -run TestMap1Demo</code></pre>

<h3>Exact match</h3>
<pre><code>go test -run ^TestMap1Demo_Basic$</code></pre>

<h3>By pattern (regex)</h3>
<pre><code>go test -run Map1Demo</code></pre>
</div>

<div class="section">
<h2>🧱 Test Naming Convention</h2>
<p>Group tests by function:</p>

<pre><code>func TestMap1Demo_Basic(t *testing.T) {}
func TestMap1Demo_Empty(t *testing.T) {}
func TestMap1Demo_Invalid(t *testing.T) {}</code></pre>

<p>👉 This allows filtering:</p>

<pre><code>go test -run Map1Demo</code></pre>
OR
<pre><code>go test ./examples -run 'TestMakeMultiplier|TestMakeAdder|TestApplyTwice' -v</code></pre>
OR
<pre><code>go test ./examples -run 'TestMake.*' -v</code></pre>
</div>

<div class="section">
<h2>🧪 Subtests (Recommended)</h2>
<p>Use subtests for better structure:</p>

<pre><code>func TestMap1Demo(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// test logic
	})

	t.Run("empty", func(t *testing.T) {
		// test logic
	})
}</code></pre>

<h3>Run specific subtest</h3>
<pre><code>go test -run Map1Demo/basic</code></pre>
</div>

<div class="section">
<h2>📊 Table-Driven Tests (Best Practice)</h2>

<pre><code>func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 0, 0},
		{"negative", -1, -1, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a + tt.b
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}</code></pre>
</div>

<div class="section">
<h2>⚠️ What NOT to Do</h2>

<h3>❌ Running a single test file</h3>
<pre><code>go test closure_test.go</code></pre>

<p>Problems:</p>
<ul>
<li>Ignores other files in the package</li>
<li>Breaks dependencies</li>
<li>Leads to inconsistent results</li>
</ul>

<h3>❌ Mixing test logic with production code</h3>
<p>Keep tests in <code>_test.go</code> files only.</p>

<h3>❌ Global state in tests</h3>
<p>Avoid shared mutable state.</p>
</div>

<div class="section">
<h2>🧠 Best Practices</h2>
<ul>
<li>Test behavior, not implementation</li>
<li>Keep tests deterministic</li>
<li>Use small, focused test cases</li>
<li>Prefer table-driven tests for scalability</li>
<li>Use <code>-run</code> instead of isolating files</li>
</ul>
</div>

<div class="section">
<h2>⚡ Quick Cheat Sheet</h2>

<table>
<tr>
<th>Task</th>
<th>Command</th>
</tr>
<tr>
<td>Run all tests</td>
<td><code>go test ./...</code></td>
</tr>
<tr>
<td>Run package tests</td>
<td><code>go test ./examples</code></td>
</tr>
<tr>
<td>Run one test</td>
<td><code>go test -run TestName</code></td>
</tr>
<tr>
<td>Run group</td>
<td><code>go test -run Pattern</code></td>
</tr>
<tr>
<td>Run subtest</td>
<td><code>go test -run Test/subtest</code></td>
</tr>
</table>
</div>

<div class="section">
<h2>🔚 Summary</h2>
<ul>
<li>Go tests are <strong>package-oriented</strong></li>
<li>Use <strong>naming + regex</strong> for control</li>
<li>Avoid file-level execution hacks</li>
<li>Structure tests for scalability</li>
</ul>
</div>

</body>
</html>