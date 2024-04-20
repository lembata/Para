/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/*.html',
    './views/**/*.html', 
    './index.html'
  ],
  theme: {
    extend: {

    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

