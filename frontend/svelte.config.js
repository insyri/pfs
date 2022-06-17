import sveltePreprocess from "svelte-preprocess"

/** @type {import('@sveltejs/kit').Config} */
export default {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: sveltePreprocess(),
	kit: {
		vite: {
			optimizeDeps: {
				include: ["highlight.js", "highlight.js/lib/core"],
			},
		},
	},
}
