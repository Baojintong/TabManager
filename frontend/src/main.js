import {createApp} from 'vue'
import adv from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import App from './App.vue'
import store from './store'

const app = createApp(App)

app.use(adv)
app.use(store)
app.mount('#app')
