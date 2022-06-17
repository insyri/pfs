<script export="module" lang="ts">
	import "../tailwind.css"

	let text = "Enter your text here."

	function upload() {
		console.log(text)
	}
</script>

<main class="index flex justify-center m-8 bg-black">
	<div
		class="py-8 px-4 border border-indigo-50 shadow-lg w-5/6 h-96 bg-slate-100"
	>
		<form on:submit|preventDefault={upload}>
			<textarea
				value={text}
				class="border border-indigo-100 shadow-lg h-72 w-5/6"
			/>
			<button type="submit" class="bg-indigo-100 p-2 rounded-full"
				>Upload!</button
			>
		</form>
	</div>
	<div />
</main>
<footer class="border-t-[1px] border-gray-700 px-40 bg-black text-white" />
