<script lang="ts">
	export const method: HttpMethods = "GET"
	export const resource_url: string = "/somewhere"
	export const description: string = "Example description."
	export const parameters: Record<string, string>[] = [{}]

	import "../tailwind.css"
	import type { HttpMethods } from "./typings"

	let showdetails = false
	function invertShowDetails() {
		console.log(!showdetails)
		showdetails = !showdetails
	}

	if (showdetails) {
	}
</script>

<div
	class="h-15 p-3 border border-blue-400 rounded-md shadow-lg hover:cursor-pointer"
	on:click={invertShowDetails}
>
	<span
		class="text-white bg-blue-400 rounded-md text-center w-20 inline-block font-bold p-2 text-sm"
	>
		{method}
	</span>
	<span class="text-gray-700 font-bold pl-2 pr-2 font-mono text-lg">
		{resource_url}
	</span>
	<span class="">{description}</span>
</div>
