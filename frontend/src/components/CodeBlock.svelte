<script>
	import Highlight from "svelte-highlight"
	import typescript from "svelte-highlight/languages/typescript"
	import github from "svelte-highlight/styles/github"

	const code = "const add = (a: number, b: number) => a + b;"
</script>

<svelte:head>
	{@html github}
</svelte:head>

<Highlight language={typescript} {code} />
