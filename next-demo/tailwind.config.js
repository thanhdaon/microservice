module.exports = {
  purge: {
    mode: "all",
    content: [
      "./components/**/*.{js,ts,jsx,tsx}",
      "./pages/**/*.{js,ts,jsx,tsx}",
      "./node_modules/antd/**/*.{js,ts,jsx,tsx}",
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
