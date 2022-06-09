<script lang="ts">
	import Highlight from "svelte-highlight"
	import json from "svelte-highlight/languages/json"
	import githubdark from "svelte-highlight/styles/github-dark"

	export let code: string
</script>

<svelte:head>
	{@html githubdark}
</svelte:head>

<Highlight class="pt-3 pb-3 text-sm" language={json} {code} />
