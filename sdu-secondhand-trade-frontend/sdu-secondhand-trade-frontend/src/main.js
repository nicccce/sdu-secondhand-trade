
import { createApp } from 'vue'
import { createPinia } from 'pinia'


import App from './App.vue'
import router from './router'
import { lazyPlugin } from './utils/imgLazy'
import { componentPlugin } from './components'

import '@/styles/common.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)
//全局导入懒加载组件
app.use(lazyPlugin)
//全局导入通用组件库
app.use(componentPlugin)

app.mount('#app')
