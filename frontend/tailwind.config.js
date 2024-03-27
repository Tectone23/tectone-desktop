/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "node_modules/flowbite-react/lib/esm/**/*.js",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
      },
      colors: {
        "tectone-blue-black": "#001939",
        "tectone-purple": "#8A25C0",
        "tectone-cyan": "#0092FF",
        "tectone-light-gray": "#F2F2F2",
        "tectone-teal": "#00A1AF",
        "tectone-pink-dark": "#F52C53",
        "tectone-gray-dark": "##6B758D",
        "tectone-paragraph-black": "#1F324A",
      },
    },
  },
  plugins: ["flowbite/plugin"],
};
