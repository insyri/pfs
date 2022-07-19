<script lang="ts">
import { schema_des } from "../deserialize"

import type { ResourceParams } from "./../types"

export let method: ResourceParams["method"]
export let short_description: ResourceParams["short_description"]
export let description: ResourceParams["description"]
export let url: ResourceParams["url"]
export let request_schema: ResourceParams["request_schema"]
export let request_schema_comments: string
// export let response_examples: ResourceParams["response_examples"]

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
	class={`h-15 rounded-md border border-blue-400 bg-blue-50 p-3 shadow-lg hover:cursor-pointer ${
		showdetails ? "rounded-b-none" : ""
	} select-none`}
	on:click={invertShowDetails}
>
	<span
		class="inline-block w-20 rounded-md bg-blue-400 p-2 text-center text-sm font-bold text-white"
	>
		{method}
	</span>
	<span class="pl-2 pr-2 font-mono text-lg font-bold text-gray-700">
		{url}
	</span>
	{short_description}
</div>
{#if showdetails}
	<div
		class="h64 rounded-md rounded-t-none border border-t-0 border-blue-400 shadow-lg"
	>
		<span class="block p-6">{description}</span>
		<div
			class="bh-white border-b-2 bg-blue-50 p-3 pl-6 text-xl font-bold text-gray-800"
		>
			Schema
		</div>
		<CodeBlock {code} />
		<div
			class="bh-white mb-1 border-b-2 bg-blue-50 p-3 pl-6 text-xl font-bold text-gray-800"
		>
			Responses
		</div>
		<table>
			<tr>
				<th class="p-3 pl-6 font-bold">Code</th>
				<th class="text-left font-bold">Response</th>
			</tr>
			<tr>
				<td class="p-3 pl-6 text-center">200</td>
				<td class="text-left">
					<CodeBlock code="things" />
				</td>
			</tr>
		</table>
	</div>
{/if}
