<script lang="ts">
	import { schema_des } from "../deserialize"

	import type { ResourceParams } from "./../types"

	export let method: ResourceParams["method"]
	export let short_description: ResourceParams["short_description"]
	export let description: ResourceParams["description"]
	export let url: ResourceParams["url"]
	export let request_schema: ResourceParams["request_schema"]
	export let request_schema_comments: string
	export let response_examples: ResourceParams["response_examples"]

	// let code: string = `{\n  ${request_schema.forEach((element) => {
	// 	schema(element)
	// })}\n}`

	let code = schema_des(request_schema, request_schema_comments)

	import "../tailwind.css"
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
	} select-none`}
	on:click={invertShowDetails}
>
	<span
		class="text-white bg-blue-400 rounded-md text-center w-20 inline-block font-bold p-2 text-sm"
	>
		{method}
	</span>
	<span class="text-gray-700 font-bold pl-2 pr-2 font-mono text-lg">
		{url}
	</span>
	{short_description}
</div>
{#if showdetails}
	<div
		class="h64 border border-blue-400 rounded-md rounded-t-none shadow-lg border-t-0"
	>
		<span class="block p-6">{description}</span>
		<div
			class="border-b-2 bh-white pl-6 p-3 bg-blue-50 text-xl font-bold text-gray-800"
		>
			Schema
		</div>
		<CodeBlock {code} />
		<div
			class="border-b-2 bh-white pl-6 p-3 mb-1 bg-blue-50 text-xl font-bold text-gray-800"
		>
			Responses
		</div>
		<table>
			<tr>
				<th class="pl-6 p-3 font-bold">Code</th>
				<th class="font-bold text-left">Response</th>
			</tr>
			<tr>
				<td class="pl-6 p-3 text-center">200</td>
				<td class="text-left">
					<CodeBlock code="things" />
				</td>
			</tr>
		</table>
	</div>
{/if}
