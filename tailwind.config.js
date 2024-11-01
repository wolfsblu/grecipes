/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    extend: {
      colors: {
        'moonstone': '#0FA3B1',
        'uranian-blue': '#B5E2FA',
        'baby-powder': '#F9F7F3',
        'vanilla': '#EDDEA4',
        'tangerine': '#F7A072'
      }
    }
  },
  plugins: [],
}

