/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.templ"],
  theme: {
    extend: {
      colors: {
        surface: "#FFFEFF",
        base: "#F2F4F8",
        subtle: "#E5E5E6",
        red: "#FC5353",
        blue: "#888BFD",
        text: "#231F20",
      },
      backgroundSize: {
        "50%": "50%",
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
