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
        'xl': '6rem',
        '2xl': '12rem',
      },
    },
    fontFamily: {
      'sans': ["Libre Franklin", ...defaultTheme.fontFamily.sans],
      'serif': [...defaultTheme.fontFamily.serif],
      'mono': [...defaultTheme.fontFamily.mono],
    },
    extend: {
      textColor: {
        "primary": "rgb(var(--background-color-800) / <alpha-value>)",
        "secondary": "rgb(var(--background-color-600) / <alpha-value>)",
        "hyperlink": "rgb(var(--obj-hyperlink-color) / <alpha-value>)",
      },
      backgroundColor: {
        "surface-0": "rgb(var(--background-color-50) / <alpha-value>)",
        "surface": "rgb(var(--background-color-100) / <alpha-value>)",
        "surface-1": "rgb(var(--background-color-200) / <alpha-value>)",
        "surface-2": "rgb(var(--background-color-300) / <alpha-value>)",
        "surface-3": "rgb(var(--background-color-400) / <alpha-value>)",
        "surface-4": "rgb(var(--background-color-500) / <alpha-value>)",

        "obj-primary": "rgb(var(--obj-primary-color) / <alpha-value>)",
        "obj-secondary": "rgb(var(--obj-secondary-color) / <alpha-value>)",
        "obj-tertiary": "rgb(var(--obj-tertiary-color) / <alpha-value>)",
        "obj-success": "rgb(var(--obj-success-color) / <alpha-value>)",
        "obj-warning": "rgb(var(--obj-warning-color) / <alpha-value>)",
        "obj-error": "rgb(var(--obj-error-color) / <alpha-value>)",
      },
      borderColor: {
        "primary": "rgb(var(--background-color-400) / <alpha-value>)",
        "secondary": "rgb(var(--background-color-300) / <alpha-value>)",
        "third": "rgb(var(--background-color-200) / <alpha-value>)",

        "obj-primary": "rgb(var(--obj-primary-color) / <alpha-value>)",
        "obj-secondary": "rgb(var(--obj-secondary-color) / <alpha-value>)",
        "obj-tertiary": "rgb(var(--obj-tertiary-color) / <alpha-value>)",
        "obj-success": "rgb(var(--obj-success-color) / <alpha-value>)",
        "obj-warning": "rgb(var(--obj-warning-color) / <alpha-value>)",
        "obj-error": "rgb(var(--obj-error-color) / <alpha-value>)",
      },
    },
    // for prose class, applied to blog markdowns
    // theme call must put inside "": https://github.com/tailwindlabs/tailwindcss/issues/9143#issuecomment-1849773132
    typography: (theme) => ({
      DEFAULT: {
        css: {
          color: "theme('textColor.primary')",
          h1: {
            fontSize: "theme('fontSize.3xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '1.75rem',
            paddingBottom: '1rem',
          },
          h2: {
            fontSize: "theme('fontSize.2xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '1.5rem',
            paddingBottom: '0.25rem',
          },
          h3: {
            fontSize: "theme('fontSize.xl')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '1rem',
            paddingBottom: '0.25rem',
          },
          h4: {
            fontSize: "theme('fontSize.lg')",
            fontWeight: "theme('fontWeight.bold')",
            paddingTop: '0.5rem',
            paddingBottom: '0.25rem',
          },
          p: {
            marginTop: '0.5rem',
            marginBottom: '0.5rem',
          },
          a: {
            color: "theme('textColor.hyperlink')",
            whiteSpace: 'nowrap',
            textDecoration: 'underline',
          },
          ul: {
            listStyleType: 'disc',
            marginTop: '0.2rem',
            marginBottom: '0.5rem',
          },
          ol: {
            listStyleType: 'decimal',
            marginTop: '0.5rem',
            marginBottom: '0.5rem',
          },
          li: {
            marginTop: '0.2rem',
            marginLeft: '1rem',
            paddingLeft: '0.75rem',
          },
          figure: {
            width: '100%',
            display: 'flex',
            flexDirection: 'column',
            gap: '0.75rem',
            paddingTop: '1.5rem',
            paddingBottom: '1.5rem',
            alignItems: 'center',
          },
          figcaption: {
            // fontSize: "theme('fontSize.sm')",
            textAlign: 'center',
            fontStyle: 'italic',
          },
          img: {
            boxShadow: "theme('boxShadow.lg')",
            marginLeft: 'auto',
            marginRight: 'auto',
          },
          blockquote: {
            backgroundColor: 'theme("backgroundColor.surface-1")',
            borderTopWidth: '0.1rem',
            borderLeftWidth: '0.3rem',
            borderTopColor: 'theme("borderColor.secondary")',
            borderLeftColor: 'theme("backgroundColor.obj-primary")',
            paddingLeft: '1.5rem',
            marginTop: '1rem',
            marginBottom: '1rem',
            paddingTop: '0.4rem',
            paddingBottom: '0.4rem',
            fontSize: '0.9rem',
            fontStyle: 'italic',
            color: 'theme("textColor.primary")',
          },
          pre: {
            marginTop: "1rem",
            marginBottom: "1.5rem",
            overflowX: 'auto',
          },
          code: {
            display: 'inline-block',
            borderColor: 'theme("borderColor.primary")',
            borderWidth: '0.2rem',
            borderRadius: '0.25rem',
            padding: '0.25rem',
            fontSize: '0.9rem',
            lineHeight: '1.4',
            letterSpacing: '-0.01rem',
          },
          hr: {
            borderColor: 'theme("borderColor.primary")',
            borderWidth: '0.1rem',
            marginTop: '1rem',
            marginBottom: '1rem',
          },
          table: {
            width: 'auto',
            marginLeft: 'auto',
            marginRight: 'auto',
            borderCollapse: 'collapse',
            marginTop: '1rem',
            marginBottom: '1rem',
          },
          tr: {
            borderWidth: '0.1rem',
            borderColor: 'theme("borderColor.primary")',
          },
          th: {
            paddingTop: '0.5rem',
            paddingBottom: '0.5rem',
            paddingLeft: '1rem',
            paddingRight: '1rem',
            borderWidth: '0.1rem',
            borderColor: 'theme("borderColor.primary")',
          },
          td: {
            paddingTop: '0.5rem',
            paddingBottom: '0.5rem',
            paddingLeft: '1rem',
            paddingRight: '1rem',
            borderWidth: '0.1rem',
            borderColor: 'theme("borderColor.primary")',
            fontSize: '0.9rem',
          },
          input: {
            marginRight: '0.5rem',
          }
        },
      },
    })
  },
  darkMode: "selector",
  plugins: [
    plugin(function ({ addBase, theme }) {
      addBase({
        'html': {
          fontSize: '16px',
          fontWeight: '450',
          lineHeight: 1.55,
        },
        'p': {
          marginTop: '0.5rem',
          marginBottom: '0.5rem',
        },
      })
    }),
    // require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    // require('daisyui'),
  ],
};

