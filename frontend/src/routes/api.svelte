<script lang="ts">
	import "../tailwind.css"
	import Resource from "../components/Resource.svelte"
</script>

<main class="m-auto w-[900px]">
	<h1 class="p-10">pfs Api</h1>
	<div>
		<Resource
			method="GET"
			short_description="Returns the contents of the file"
			url="/:file"
			description="DEScc"
		/>
	</div>
</main>
