/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')
const plugin = require('tailwindcss/plugin')

module.exports = {
  content: [
    "./internal/**/*.{templ,html,js,go}",
  ],
  theme: {
    container: {
      center: true,
      padding: {
        'lg': '2.5rem',
        'xl': '8rem',
        '2xl': '15rem',
      },
    },
    fontFamily: {
      'sans': ['"Libre Franklin"', ...defaultTheme.fontFamily.sans],
      'serif': [...defaultTheme.fontFamily.serif],
      'mono': [...defaultTheme.fontFamily.mono],
    },
    extend: {
      textColor: {
        "primary": "rgb(var(--background-color-900) / <alpha-value>)",
        "secondary": "rgb(var(--background-color-700) / <alpha-value>)",
        "hyperlink": "rgb(var(--hyperlink-color) / <alpha-value>)",
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
      },
    },
    // for prose class, applied to blog markdowns
    // theme call must put inside "": https://github.com/tailwindlabs/tailwindcss/issues/9143#issuecomment-1849773132
    typography: (theme) => ({
      DEFAULT: {
        css: {
          color: "theme('backgroundColor.surface')",
          h1: {
            fontSize: "theme('fontSize.3xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '1.75rem',
            paddingBottom: '1.4rem',
          },
          h2: {
            fontSize: "theme('fontSize.2xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '1.1rem',
            paddingBottom: '0.9rem',
          },
          h3: {
            fontSize: "theme('fontSize.xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '0.8rem',
            paddingBottom: '0.75rem',
          },
          h4: {
            fontSize: "theme('fontSize.lg')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '0.6rem',
            paddingBottom: '0.5rem',
          },
          p: {
            marginTop: '0.35rem',
            marginBottom: '0.35rem',
          },
          a: {
            color: "theme('textColor.hyperlink')",
            whiteSpace: 'nowrap',
            textDecoration: 'underline',
          },
          li: {
            listStyleType: 'disc',
            listStylePosition: 'inside',
          },
          figure: {
            width: '100%',
            display: 'flex',
            flexDirection: 'column',
            gap: '1rem',
            paddingTop: '1.5rem',
            paddingBottom: '1.5rem',
            justifyContent: 'center',
          },
          figcaption: {
            fontSize: "theme('fontSize.sm')",
            textAlign: 'center',
            fontStyle: 'italic',
          },
          img: {
            boxShadow: "theme('boxShadow.lg')",
            marginLeft: 'auto',
            marginRight: 'auto',
          },
        },
      },
    })
  },
  darkMode: "selector",
  plugins: [
    plugin(function ({ addBase, theme }) {
      addBase({
        'p': {
          marginTop: '0.35rem',
          marginBottom: '0.35rem',
        },
      })
    }),
    // require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    // require('daisyui'),
  ],
};

