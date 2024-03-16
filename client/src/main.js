import axios from 'axios'
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import 'bootstrap/dist/css/bootstrap.css'

import App from './App.vue'

import ProductList from './components/pages/ProductList'
import ProductCreate from './components/pages/ProductCreate'
import ProductShow from './components/pages/ProductShow'
import ProductEdit from './components/pages/ProductEdit'

axios.defaults.baseURL = process.env.VUE_APP_API_URL

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: '/', component: ProductList },
		{ path: '/add', component: ProductCreate },
		{ path: '/edit/:id', component: ProductEdit },
		{ path: '/show/:id', component: ProductShow },
	],
})

createApp(App).use(router).mount('#app')
