import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import * as path from "path";
import vueJsx from "@vitejs/plugin-vue-jsx";
import ViteComponents from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: '../www',
    emptyOutDir: true,
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        ws: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
  plugins: [
    vue(),
    vueJsx(),
    ViteComponents({
      resolvers: [
        NaiveUiResolver(),
      ],
    }),
  ],
  resolve: {
    alias: [
      {
        find: "@",
        replacement: path.resolve(__dirname, "src"),
      },
    ],
  },
});
