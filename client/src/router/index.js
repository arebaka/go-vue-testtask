import { createRouter, createWebHistory } from 'vue-router'

import ProductList from '@/views/product/List'
import ProductCreate from '@/views/product/Create'
import ProductShow from '@/views/product/Show'
import ProductEdit from '@/views/product/Edit'

const routes = [
	{
		path: '/',
		name: 'ProductList',
		component: ProductList,
	},
	{
		path: '/add',
		name: 'ProductCreate',
		component: ProductCreate,
	},
	{
		path: '/show/:id',
		name: 'ProductShow',
		component: ProductShow,
	},
	{
		path: '/edit/:id',
		name: 'ProductEdit',
		component: ProductEdit,
	},
]

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes: routes,
})

export default router
