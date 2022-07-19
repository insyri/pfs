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
			description="Some long description. Lorem ipsum dolor sit amet."
			request_schema={[
				{
					description: "The password needed to access the file",
					key_name: "password",
					optional: true,
					value_type: "string",
				},
			]}
			request_schema_comments="Only needed if the file needs a password. Body is not needed if the file is public."
		/>
		<!--
			response_examples={[
				{
					http_code: 200,
				},
			]}
		 -->
	</div>
</main>
