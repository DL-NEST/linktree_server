import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
const path = require('path')

// https://vitejs.dev/config/
export default defineConfig({
  server:{
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8000',
        changeOrigin: true,
      }
    }
  },
  plugins: [react()],
  resolve:{
    alias:{
      '@': path.resolve(__dirname, './src')
    }
  },
  css:{
    preprocessorOptions:{
      scss:{
        additionalData: '@import "style/var.scss";', // 添加公共样式
      }
    }
  }
})
