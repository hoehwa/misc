<h2 id="syntax"><a href="#syntax">Syntax</a></h2><div class="section-content"><div class="code-example"><pre class="brush: js notranslate"><code>x <span class="token operator">+</span> y
</code></pre></div></div><h2 id="examples"><a href="#examples">Examples</a></h2><div class="section-content"></div><h3 id="number_addition"><a href="#number_addition">Number addition</a></h3><div class="section-content"><div class="code-example"><pre class="brush: js notranslate"><code><span class="token comment">// Number + Number -&gt; addition</span>
<span class="token number">1</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span> <span class="token comment">// 3</span>

<span class="token comment">// Boolean + Number -&gt; addition</span>
<span class="token boolean">true</span> <span class="token operator">+</span> <span class="token number">1</span><span class="token punctuation">;</span> <span class="token comment">// 2</span>

<span class="token comment">// Boolean + Boolean -&gt; addition</span>
<span class="token boolean">false</span> <span class="token operator">+</span> <span class="token boolean">false</span><span class="token punctuation">;</span> <span class="token comment">// 0</span>
</code></pre></div></div><h3 id="bigint_addition"><a href="#bigint_addition">BigInt addition</a></h3><div class="section-content"><div class="code-example"><pre class="brush: js notranslate"><code><span class="token comment">// BigInt + BigInt -&gt; addition</span>
<span class="token number">1n</span> <span class="token operator">+</span> <span class="token number">2n</span><span class="token punctuation">;</span> <span class="token comment">// 3n</span>

<span class="token comment">// BigInt + Number -&gt; throws TypeError</span>
<span class="token number">1n</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span> <span class="token comment">// TypeError: Cannot mix BigInt and other types, use explicit conversions</span>

<span class="token comment">// To add a BigInt to a non-BigInt, convert either operand</span>
<span class="token number">1n</span> <span class="token operator">+</span> <span class="token function">BigInt</span><span class="token punctuation">(</span><span class="token number">2</span><span class="token punctuation">)</span><span class="token punctuation">;</span> <span class="token comment">// 3n</span>
<span class="token function">Number</span><span class="token punctuation">(</span><span class="token number">1n</span><span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span> <span class="token comment">// 3</span>
</code></pre></div></div><h3 id="string_concatenation"><a href="#string_concatenation">String concatenation</a></h3><div class="section-content"><div class="code-example"><pre class="brush: js notranslate"><code><span class="token comment">// String + String -&gt; concatenation</span>
<span class="token string">&#34;foo&#34;</span> <span class="token operator">+</span> <span class="token string">&#34;bar&#34;</span><span class="token punctuation">;</span> <span class="token comment">// &#34;foobar&#34;</span>

<span class="token comment">// Number + String -&gt; concatenation</span>
<span class="token number">5</span> <span class="token operator">+</span> <span class="token string">&#34;foo&#34;</span><span class="token punctuation">;</span> <span class="token comment">// &#34;5foo&#34;</span>

<span class="token comment">// String + Boolean -&gt; concatenation</span>
<span class="token string">&#34;foo&#34;</span> <span class="token operator">+</span> <span class="token boolean">false</span><span class="token punctuation">;</span> <span class="token comment">// &#34;foofalse&#34;</span>

<span class="token comment">// String + Number -&gt; concatenation</span>
<span class="token string">&#34;2&#34;</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span> <span class="token comment">// &#34;22&#34;</span>
</code></pre></div></div><h2 id="see_also"><a href="#see_also">See also</a></h2><div class="section-content"><ul>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Subtraction">Subtraction operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Division">Division operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Multiplication">Multiplication operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Remainder">Remainder operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Exponentiation">Exponentiation operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Increment">Increment operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Decrement">Decrement operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Unary_negation">Unary negation operator</a></li>
  <li><a href="/en-US/docs/Web/JavaScript/Reference/Operators/Unary_plus">Unary plus operator</a></li>
</ul></div><h2 id="description"><a href="#description">Description</a></h2><div class="section-content"><p>The <code>+</code> operator is overloaded for two distinct operations: numeric addition and string concatenation. When evaluating, it first <a href="/en-US/docs/Web/JavaScript/Data_structures#primitive_coercion">coerces both operands to primitives</a>. Then, the two operands&#39; types are tested:</p>
<ul>
  <li>If one side is a string, the other operand is also <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/String#string_coercion">converted to a string</a> and they are concatenated.</li>
  <li>If they are both <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigInt">BigInts</a>, BigInt addition is performed. If one side is a BigInt but the other is not, a <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/TypeError"><code>TypeError</code></a> is thrown.</li>
  <li>Otherwise, both sides are <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number#number_coercion">converted to numbers</a>, and numeric addition is performed.</li>
</ul>
<p>String concatenation is often thought to be equivalent with <a href="/en-US/docs/Web/JavaScript/Reference/Template_literals">template literals</a> or <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/concat"><code>String.prototype.concat()</code></a>, but they are not. Addition coerces the expression to a <em>primitive</em>, which calls <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/valueOf"><code>valueOf()</code></a> in priority; on the other hand, template literals and <code>concat()</code> coerce the expression to a <em>string</em>, which calls <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/toString"><code>toString()</code></a> in priority. If the expression has a <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/toPrimitive"><code>@@toPrimitive</code></a> method, string concatenation calls it with <code>&#34;default&#34;</code> as hint, while template literals use <code>&#34;string&#34;</code>. This is important for objects that have different string and primitive representations — such as <a href="https://github.com/tc39/proposal-temporal" class="external" target="_blank">Temporal</a>, whose <code>valueOf()</code> method throws.</p>
<div class="code-example"><pre class="brush: js notranslate"><code><span class="token keyword">const</span> t <span class="token operator">=</span> Temporal<span class="token punctuation">.</span>Now<span class="token punctuation">.</span><span class="token function">instant</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
<span class="token string">&#34;&#34;</span> <span class="token operator">+</span> t<span class="token punctuation">;</span> <span class="token comment">// Throws TypeError</span>
<span class="token template-string"><span class="token template-punctuation string">`</span><span class="token interpolation"><span class="token interpolation-punctuation punctuation">${</span>t<span class="token interpolation-punctuation punctuation">}</span></span><span class="token template-punctuation string">`</span></span><span class="token punctuation">;</span> <span class="token comment">// &#39;2022-07-31T04:48:56.113918308Z&#39;</span>
<span class="token string">&#34;&#34;</span><span class="token punctuation">.</span><span class="token function">concat</span><span class="token punctuation">(</span>t<span class="token punctuation">)</span><span class="token punctuation">;</span> <span class="token comment">// &#39;2022-07-31T04:48:56.113918308Z&#39;</span>
</code></pre></div>
<p>You are advised to not use <code>&#34;&#34; + x</code> to perform <a href="/en-US/docs/Web/JavaScript/Reference/Global_Objects/String#string_coercion">string coercion</a>.</p></div><thead><tr><th scope="col">Specification</th></tr></thead><tbody><tr><td><a href="https://tc39.es/ecma262/multipage/ecmascript-language-expressions.html#sec-addition-operator-plus">ECMAScript Language Specification<!-- --> <br/><small># <!-- -->sec-addition-operator-plus</small></a></td></tr></tbody>