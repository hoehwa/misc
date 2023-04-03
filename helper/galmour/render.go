package galmour

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)


func Render() {
	in := `
	<div class="section-content"><p>You can use the comma operator when you want to include multiple expressions in a location that requires a single expression. The most common usage of this operator is to supply multiple updaters in a <code>for</code> loop.</p>
<p>Because all expressions except the last are evaluated and then discarded, these expressions must have side effects to be useful. Common expressions that have side effects are assignments, function calls, and <a href="/en-US/docs/Web/JavaScript/Reference/Operators/Increment"><code>++</code></a> and <a href="/en-US/docs/Web/JavaScript/Reference/Operators/Decrement"><code>--</code></a> operators. Others may also have side effects if they invoke <a href="/en-US/docs/Web/JavaScript/Reference/Functions/get">getters</a> or trigger <a href="/en-US/docs/Web/JavaScript/Data_structures#type_coercion">type coercions</a>.</p>
<p>The comma operator has the lowest <a href="/en-US/docs/Web/JavaScript/Reference/Operators/Operator_Precedence">precedence</a> of all operators. If you want to incorporate a comma-joined expression into a bigger expression, you must parenthesize it.</p>
<p>The comma operator is completely different from commas used as syntactic separators in other locations, which include:</p>
<ul>
  <li>Elements in array initializers (<code>[1, 2, 3]</code>)</li>
  <li>Properties in <a href="/en-US/docs/Web/JavaScript/Reference/Operators/Object_initializer">object initializers</a> (<code>{ a: 1, b: 2 }</code>)</li>
  <li>Parameters in <a href="/en-US/docs/Web/JavaScript/Reference/Statements/function">function declarations</a>/expressions (<code>function f(a, b) { … }</code>)</li>
  <li>Arguments in function calls (<code>f(1, 2)</code>)</li>
  <li>Binding lists in <a href="/en-US/docs/Web/JavaScript/Reference/Statements/let"><code>let</code></a>, <a href="/en-US/docs/Web/JavaScript/Reference/Statements/const"><code>const</code></a>, or <a href="/en-US/docs/Web/JavaScript/Reference/Statements/var"><code>var</code></a> declarations (<code>const a = 1, b = 2;</code>)</li>
  <li>Import lists in <a href="/en-US/docs/Web/JavaScript/Reference/Statements/import"><code>import</code></a> declarations (<code>import { a, b } from "c";</code>)</li>
  <li>Export lists in <a href="/en-US/docs/Web/JavaScript/Reference/Statements/export"><code>export</code></a> declarations (<code>export { a, b };</code>)</li>
</ul>
<p>In fact, although some of these places accept almost all expressions, they don't accept comma-joined expressions because that would be ambiguous with the syntactic comma separators. In this case, you must parenthesize the comma-joined expression. For example, the following is a <code>const</code> declaration that declares two variables, where the comma is not the comma operator:</p>
<div class="code-example"><pre class="brush: js notranslate"><code><span class="token keyword">const</span> a <span class="token operator">=</span> <span class="token number">1</span><span class="token punctuation">,</span> b <span class="token operator">=</span> <span class="token number">2</span><span class="token punctuation">;</span>
</code></pre><button type="button" class="icon copy-icon"><span class="visually-hidden">Copy to Clipboard</span></button><span class="copy-icon-message visually-hidden" role="alert" style="top: 52px;"></span></div>
<p>It is different from the following, where <code>b = 2</code> is an <a href="/en-US/docs/Web/JavaScript/Reference/Operators/Assignment">assignment expression</a>, not a declaration. The value of <code>a</code> is <code>2</code>, the return value of the assignment, while the value of <code>1</code> is discarded:</p>
<div class="code-example"><pre class="brush: js notranslate"><code><span class="token keyword">const</span> a <span class="token operator">=</span> <span class="token punctuation">(</span><span class="token number">1</span><span class="token punctuation">,</span> b <span class="token operator">=</span> <span class="token number">2</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
</code></pre><button type="button" class="icon copy-icon"><span class="visually-hidden">Copy to Clipboard</span></button><span class="copy-icon-message visually-hidden" role="alert" style="top: 52px;"></span></div>
<p>Comma operators cannot appear as <a href="/en-US/docs/Web/JavaScript/Reference/Trailing_commas">trailing commas</a>.</p></div>
	`

	out, _ := glamour.Render(in, "dark")
	fmt.Print(out)
}
