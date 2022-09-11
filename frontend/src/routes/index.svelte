<script export="module" lang="ts">
import "../tailwind.css"

// let text = "Enter your text here."

function onSubmit(e: { target: any }) {
	const formData = new FormData(e.target)

	const data = {}
	for (let field of formData) {
		const [key, value] = field
		data[key] = value
	}

	// const date = new Date()

	data["created_at"] = Date.now()
	// 604800 = seconds in 1 week
	data["expires_at"] = Date.now() + 604800 * 3 * 1000

	let xhr = new XMLHttpRequest()
	xhr.open("POST", "http://localhost:8080/api/upload/paste", true)

	xhr.setRequestHeader("Accept", "application/json")
	xhr.setRequestHeader("Content-Type", "application/json")
	xhr.setRequestHeader("Access-Control-Allow-Origin", "*")

	xhr.onload = () => {
		if (xhr.status === 200) {
			console.log(xhr.responseText)
		} else {
			console.log("Request failed. Returned status of " + xhr.status)
		}
	}

	xhr.send(JSON.stringify(data))

	console.log("Text:" + data["text"])
}
</script>

<main class="max-w-8xl m-4 bg-white shadow-lg shadow-gray-700">
	<div class="m-auto h-[96vh] w-[80%] max-w-4xl">
		<form on:submit={onSubmit}>
			<!-- File upload button -->
			<!-- Textbox -->
			<textarea
				required
				id="text"
				class="mt-20 h-[60vh] w-full rounded-md border border-slate-600 bg-slate-200 p-2 text-sm"
				placeholder="Enter your text here."
			/>
			<!-- Other buttons + submit -->
			<!-- <input type="submit" value="Submit"> -->
			<input
				class="w-35 float-right h-10 rounded-md bg-slate-500 font-semibold text-slate-100"
				type="submit"
			/>
		</form>
	</div>
</main>
