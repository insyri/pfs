<script lang="ts">
	import "../tailwind.css"
	import Resource from "$lib/Resource.svelte"
</script>

<main class="m-auto w-[960px]">
	<h1 class="p-10">pfs Api</h1>
	<div>
		<Resource
			method="GET"
			description="Returns the contents of the file"
			parameters={[]}
			resource_url={"/:file"}
		/>
	</div>
</main>
