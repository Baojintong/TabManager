import {createApp} from 'vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import infiniteScroll from 'vue-infinite-scroll'
import App from './App.vue'
const app = createApp(App)


app.use(Antd)
app.use(infiniteScroll)
app.mount('#app')
