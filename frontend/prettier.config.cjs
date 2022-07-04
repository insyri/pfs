/** @type {import('prettier').Config} */
module.exports = {
	useTabs: true,
	trailingComma: "es5",
	tabWidth: 4,
	semi: false,
	singleQuote: false,
	svelteStrictMode: false,
	svelteAllowShorthand: true,
	svelteIndentScriptAndStyle: false,
	plugins: [require("prettier-plugin-tailwindcss")],
}
