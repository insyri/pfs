<script export="module" lang="ts">
	import "../tailwind.css"

	let text = "Enter your text here."

	function onSubmit(e) {
		const formData = new FormData(e.target)

		const data = {}
		for (let field of formData) {
			const [key, value] = field
			data[key] = value
		}

		const date = new Date()

		data["created_at"] = Date.now()
		// 604800 = seconds in 1 week
		data["expires_at"] = Date.now() + 604800 * 3 * 1000

		console.log(data)
	}
</script>

<main>
	<div class="bg-black h-20 md:h-60 border-b-[1px] border-[#323232]" />
	<form
		class="bg-gray-400 p-4"
		action="/api/upload/paste"
		method="post"
		on:submit|preventDefault={onSubmit}
	>
		<input type="text" id="text" name="text" value={text} />
		<input type="" />
		<input class="hover:cursor-pointer" type="submit" />
	</form>
</main>

<!-- <main class="index flex justify-center m-8">
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
</main> -->
