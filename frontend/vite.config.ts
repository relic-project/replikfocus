import { fileURLToPath, URL } from 'node:url'
import { VitePWA } from 'vite-plugin-pwa'
import { defineConfig } from 'vite'

import vue from '@vitejs/plugin-vue'
import icons from './manifest.icons'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    VitePWA({
      registerType: 'autoUpdate',
      devOptions: {
        enabled: true
      },
      manifest: {
        name: 'ReplikFocus',
        short_name: 'ReplikFocus',
        theme_color: '#ffffff',
        icons: icons,
        display: 'fullscreen',
      }
    }),

  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      '/ws': {
        target: 'ws://localhost:1323',
        changeOrigin: true,
        ws: true
      }
    }
  },

})
