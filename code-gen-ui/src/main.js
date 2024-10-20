import './assets/css/normalize.css'
import './assets/css/base.scss'
import './assets/css/main.scss'

//引入nprogress的css
import 'nprogress/nprogress.css';


import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)


//引入elementUI图标和样式
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import "element-plus/dist/index.css"
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}


import pinia from "@/stores/index.js";
app.use(pinia)
app.use(router)

app.mount('#app')
