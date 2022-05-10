<script context="module">
	export const prerender = true;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section class="container">
	<h1>Raul Grell Cartolano Leal</h1>
	<h2 class="subtitle">IT Consultant - Software Developer</h2>

	<h3>Projects:</h3>
	<div>
		<a href="https://github.com/raulgrell/zox">Zox</a> - A programming language
		implemented with a bytecode compiler
	</div>
	<div>
		<a href="https://github.com/raulgrell/tetris">Tetris</a> - A WASM port of
		andrewrk/tetris, written in Zig
	</div>
	<div>
		<a href="https://github.com/raulgrell/advent">Advent</a> - Solving the advent
		of code in Zig
	</div>
	<div>
		<a href="https://github.com/raulgrell/tick">Tick</a> - An quick foray into
		designing a game development library
	</div>
	<div>
		<a href="https://buas.itch.io/warfair">WarFair</a> - A turn-based strategy
		game featuring a creazy clowns
	</div>
	<div>
		<a href="https://github.com/raulgrell/raulgrell.github.io">This</a> - Using
		Svelte as a static site generator
	</div>

	<h3>Interests:</h3>
	<div><a href="https://github.com/ziglang/zig">Zig</a></div>
	<div><a href="https://github.com/ziglang/zig">WASM</a></div>
	<div><a href="https://github.com/ziglang/zig">Dat</a></div>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		flex: 1;
	}

	h1 {
		width: 100%;
	}
</style>
