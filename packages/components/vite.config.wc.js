import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import dts from 'vite-plugin-dts'
import Components from 'unplugin-vue-components/vite'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import DefineOptions from 'unplugin-vue-define-options/vite'
import cssInjectedByJsPlugin from 'vite-plugin-css-injected-by-js'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

const filename = fileURLToPath(import.meta.url)
const dirname = path.dirname(filename)

export default defineConfig({
  server: {
    fs: {
      allow: ['..'],
    },
  },
  plugins: [
    vue({
      style: {
        filename: './style.css',
      },
      customElement: true,
    }),
    cssInjectedByJsPlugin(),
    Components({
      dirs: ['src'],
      extensions: ['vue'],
      directoryAsNamespace: true,
      globalNamespaces: ['global'],
      include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
      exclude: [/node_modules/, /\.git/],
      resolvers: [],
    }),
    dts({
      tsconfigPath: 'tsconfig.build.json',
      cleanVueFileName: true,
      exclude: ['src/test/**', 'src/**/story/**', 'src/**/*.story.vue'],
    }),
    DefineOptions(),
  ],
  build: {
    outDir: 'dist/web-components',
  },
  resolve: {
    alias: {
      '@': path.resolve(dirname, './src'),
      '@apis': path.resolve(dirname, './src/apis'),
      '@components': path.resolve(dirname, './src/components'),
      '@utils': path.resolve(dirname, './src/utils'),
      '@contracts': path.resolve(dirname, '../contracts/src/'),
    },
  },
})
