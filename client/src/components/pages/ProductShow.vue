<template>
	<div class="container">
		<h2 class="text-center mt-5 mb-3"> Продукт </h2>
		<div class="card">
			<div class="card-header">
				<router-link to="/"
					class="btn btn-outline-info float-right"
				> Все продукты
				</router-link>
			</div>
			<div class="card-body">
				<b class="text-muted"> Название: </b>
				<p>{{product.name}}</p>
				<b class="text-muted"> Описание: </b>
				<p>{{product.description}}</p>
				<b class="text-muted"> Цена: </b>
				<p>{{product.price.toFixed(2)}}</p>
				<b class="text-muted"> Создан: </b>
				<p>{{product.created_at.toLocaleString()}}</p>
			</div>
		</div>
	</div>
 </template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
	name: 'ProductShow',
	data() {
		return {
			product: {
				id: this.$route.params.id,
				name: '',
				description: '',
				price: 0,
				created_at: '',
			},
			isSaving: false,
		}
	},
	created() {
		axios.get(`/${this.product.id}`)
		.then(response => {
			this.product.name = response.data.name
			this.product.description = response.data.description
			this.product.price = response.data.price / 100
			this.product.created_at = new Date(Date.parse(response.data.created_at))
			return response
		})
		.catch(error => {
			Swal.fire({
				icon: 'error',
				title: 'Ошибка получения данных о продукте!',
				showConfirmButton: false,
				timer: 1500
			})
			return error
		})
	},
	methods: {},
}
</script>
