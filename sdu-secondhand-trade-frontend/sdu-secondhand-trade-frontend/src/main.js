
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'
import { lazyPlugin } from './utils/imgLazy'
import { componentPlugin } from './components'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import '@/styles/common.scss'

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
const pinia = createPinia()

pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(router)
//全局导入懒加载组件
app.use(lazyPlugin)
//全局导入通用组件库
app.use(componentPlugin)

app.mount('#app')
