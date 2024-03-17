import axios from 'axios'
import { createApp } from 'vue'

import 'bootstrap/dist/css/bootstrap.css'

import router from './router'

import App from './App.vue'

axios.defaults.baseURL = process.env.VUE_APP_API_URL

const app = createApp(App)
app.use(router)
app.config.devtools = process.env.NODE_ENV == 'development'
app.mount('#app')
