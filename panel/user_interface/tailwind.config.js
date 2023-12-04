/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      boxShadow: {
        btn: "0 0 3px 0 rgb(150, 150, 150);",
      },
    },
  },
  plugins: [],
};
