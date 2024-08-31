/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./internal/**/*.{templ,html,js,go}",
  ],
  theme: {
    container: {
      center: true,
    },
    fontFamily: {
      'sans': ['"Libre Franklin"', ...defaultTheme.fontFamily.sans],
      'serif': [...defaultTheme.fontFamily.serif],
      'mono': [...defaultTheme.fontFamily.mono],
    },
    extend: {
      textColor: {
        primary: "rgb(var(--background-color-900) / <alpha-value>)",
        secondary: "rgb(var(--background-color-700) / <alpha-value>)",
      },
      backgroundColor: {
        "surface": "rgb(var(--background-color-100) / <alpha-value>)",
        "surface-1": "rgb(var(--background-color-200) / <alpha-value>)",
        "surface-2": "rgb(var(--background-color-300) / <alpha-value>)",
        "surface-3": "rgb(var(--background-color-400) / <alpha-value>)",
        "surface-4": "rgb(var(--background-color-500) / <alpha-value>)",
      },
      borderColor: {
        "primary": "rgb(var(--background-color-400) / <alpha-value>)",
        "secondary": "rgb(var(--background-color-300) / <alpha-value>)",
      }
    },
  },
  darkMode: "selector",
  plugins: [
    // require('@tailwindcss/forms'),
    // require('@tailwindcss/typography'),
    // require('daisyui'),
  ],
};

