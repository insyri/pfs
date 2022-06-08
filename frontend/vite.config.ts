import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  // resolve: { alias: { '@': (await import("path")).resolve(__dirname, 'src') } },
  plugins: [svelte()],
  server: {
    host: true,
    port: 3000
  }
})
