/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
const withMT = require('@material-tailwind/react/utils/withMT')

export default withMT({
  darkMode: 'class',
  content: [
    './src/**/*.{js,jsx,ts,tsx}', 
    'index.html', 
    "./node_modules/react-tailwindcss-datepicker/dist/index.esm.{js,ts}",
  ],
  theme: {
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px',
    },
    extend: {
      colors: {
        lime: colors.lime,
        white: colors.white,
        gray: colors.gray,
        blue: colors.blue,
        red: colors.red,
      },
      fontFamily: {
        sans: ['Inter', 'font-mono'],
      },
    },
    plugins: [
      require('@tailwindcss/typography'),
      require('@tailwindcss/forms'),
      require('@headlessui/tailwindcss'),
    ],
  },
})
