module.exports = {
  purge: {
    mode: "all",
    content: [
      "./components/**/*.{js,ts,jsx,tsx}",
      "./pages/**/*.{js,ts,jsx,tsx}",
      "./layout/**/*.{js,ts,jsx,tsx}",
      "./node_modules/antd/**/*.{js,ts,jsx,tsx,css,less}",
    ],
  },
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  theme: {},
  variants: {},
  plugins: [],
};
