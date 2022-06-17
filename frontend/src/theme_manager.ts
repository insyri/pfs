type Theme = "dark" | "light"
export default function getTheme(): Theme {
  if (!window) return "light"
  const systemPreference: Theme = window.matchMedia("(prefers-color-scheme: dark").matches ? "dark" : "light"

  let theme = localStorage.getItem("theme") as Theme | null
  if (!theme) {
    setTheme(systemPreference)
    theme = systemPreference
  }
  
  return theme
}

export function setTheme(t: Theme) {
  localStorage.setItem("theme", t)
}