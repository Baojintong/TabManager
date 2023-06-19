import {createApp} from 'vue'
import adv from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import App from './App.vue'
import store from './store'
import collect from './components/Collect.vue'
import { createRouter, createWebHistory } from 'vue-router'

const app = createApp(App)


const routes = [
    { path: '/collect', component: collect },
]
const router = createRouter({
    history: createWebHistory(),
    routes,
})

app.use(adv)
app.use(store)
app.use(router)
app.mount('#app')
