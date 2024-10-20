import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'


// https://vitejs.dev/config/
export default defineConfig({
  //可以更好的调试
  build: {
    sourcemap: true,
  },
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:{
    port:5173,

    //动态代理
    proxy:{
      '^/api':{
        //目标地址，这里应该填写网关地址，axios的baseUrl直接写一个/api就好了
        target:'http://127.0.0.1:8001',

        //允许跨域,也就是是否欺骗后台
        changeOrigin:true,

        //重写路径,\/api意思是转义/符号
        rewrite:(path)=>path.replace(/^\/api/,'')
      }
    }
  }
})
