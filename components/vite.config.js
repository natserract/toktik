import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import dts from 'vite-plugin-dts'
import Components from 'unplugin-vue-components/vite'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import DefineOptions from 'unplugin-vue-define-options/vite'

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
    Components({
      dirs: ['src'],
      extensions: ['vue'],
      directoryAsNamespace: true,
      globalNamespaces: ['global'],
      include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
      exclude: [/node_modules/, /\.git/],
      resolvers: [],
    }),
    dts(),
    DefineOptions(),
  ],
  build: {
    outDir: 'dist',
    lib: {
      entry: path.resolve(dirname, 'src/index.ts'),
      name: 'components',
      formats: ['es', 'cjs'],
    },
    sourcemap: true,
    rollupOptions: {
      // manualChunks: (id) => {
      //   if (id.includes('node_modules')) {
      //     return 'vendor';
      //   }
      //   // return path.parse(id).name;
      // },
      external: ['vue'],
      input: ['src/index.ts'],
      output: {
        // minifyInternalExports: false,
        // chunkFileNames: '[name].js',
        // Provide global variables to use in the UMD build
        // for externalized deps
        globals: {
          vue: 'Vue',
        },
      },
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(dirname, './src'),
      '@components': path.resolve(dirname, './src/components'),
    },
  },
})
