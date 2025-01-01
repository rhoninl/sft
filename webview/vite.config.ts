import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      proto: path.resolve(__dirname, "../proto"),
    },
  },
  server: {
    port: 3000,
  },
});
