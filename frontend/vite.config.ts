import vue from "@vitejs/plugin-vue";
import path from "node:path";
import { defineConfig } from "vite";

const resolve = (dir: string) => path.resolve(__dirname, dir);

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      script: {
        defineModel: true,
      },
    }),
  ],
  resolve: {
    alias: {
      "@": resolve("src"),
    },
  },
});
