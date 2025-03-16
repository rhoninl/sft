const { heroui } = require("@heroui/react");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  plugins: [
    heroui({
      themes: {
        cyber: {
          colors: {
            background: "#0d0d0d",
            foreground: "#a40aff",
            primary: {
              foreground: "#ffffff",
              DEFAULT: "#a40aff",
            },
            content1: {
              DEFAULT: "#73d18d",
            },
          },
        },
        light: {
          colors: {
            background: "#FFFFFF",
            foreground: "#11181C",
            primary: {
              foreground: "#FFFFFF",
              DEFAULT: "#006FEE",
            },
          },
        },
        dark: {
          colors: {
            background: "#2f2f36",
            foreground: "#DCEDEE",
            primary: {
              foreground: "#FFFFFF",
              DEFAULT: "#006FEE",
            },
            content1: {
              DEFAULT: "#3f3f46",
            },
          },
        },
      },
      defaultTheme: "light",
    }),
  ],
};
