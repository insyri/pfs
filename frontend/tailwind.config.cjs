/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./index.html", "./src/**/*.{svelte,js,ts}"],
	theme: {
		extend: {
			colors: {
				greyscale: {
					light: "#323232",
					dark: "#131313",
				},
			},
		},
	},
	plugins: [],
}
