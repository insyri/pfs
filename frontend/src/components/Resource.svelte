<script lang="ts">
	export const method: HttpMethods = "GET"
	export const resource_url: string = "/somewhere"
	export const short_description: string = " Ex. Short description."
	export const description: string =
		"Example longer description. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
	export const parameters: Record<string, string>[] = [{}]

	import "../tailwind.css"
	import type { HttpMethods } from "../typings"
	import CodeBlock from "./CodeBlock.svelte"

	let showdetails = true // false
	function invertShowDetails() {
		console.log(!showdetails)
		showdetails = !showdetails
	}
</script>

<div
	class={`bg-blue-50 h-15 p-3 border border-blue-400 rounded-md shadow-lg hover:cursor-pointer ${
		showdetails ? "rounded-b-none" : ""
	}`}
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
	{short_description}
</div>
{#if showdetails}
	<div
		class="h64 border border-blue-400 rounded-md rounded-t-none shadow-lg border-t-0"
	>
		<span class="block p-6">{description}</span>
		<div class="h-28 border-b-2 bh-white p-6 bg-white text-bold text-xl">
			Parameters
		</div>
		<CodeBlock
			code="[
			&lbrace;
				'code': 0,
				'message': 'string',
				'reference': 'string'
			&rbrace;
		]"
		/>
	</div>
{/if}
