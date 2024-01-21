/// <reference types="vite" />
import { defineConfig } from "vite";

export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        main: "src/main.ts",
        "components/lazy_comp_one": "src/components/lazy_comp_one.ts",
        "components/lazy_comp_two": "src/components/lazy_comp_two.ts",
      },
      output: {
        entryFileNames: "[name].js",
        dir: "./static/dist",
      },
    },
  },
});
