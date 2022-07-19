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
			console.log("Request failed.  Returned status of " + xhr.status)
		}
	}

	xhr.send(JSON.stringify(data))

	console.log(JSON.stringify(data))
}
</script>

<main>
	<div class="h-20 border-b-[1px] border-[#323232] bg-black md:h-40" />
	<form
		class="m-4 rounded-lg border-[1px] border-[#323232] bg-black p-4"
		on:submit|preventDefault={onSubmit}
	>
		<input
			class="m-4 rounded-lg border-[1px] border-[#323232] bg-[#1d1d1d] p-4 text-white"
			type="text"
			id="text"
			name="text"
			placeholder="Enter your text here."
		/>
		<input
			class="duration-250 rounded-lg border-[1px] border-[#323232] bg-white p-4 transition ease-in-out hover:cursor-pointer hover:bg-black hover:text-white"
			type="submit"
		/>
	</form>
</main>
