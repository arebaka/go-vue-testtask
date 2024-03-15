<template>
	<layout-div>
		<div class="container">
			<h2 class="text-center mt-5 mb-3"> Продукты </h2>
			<div class="card">
				<div class="card-header">
					<router-link to="/add"
						class="btn btn-outline-primary"
					> Новый продукт
					</router-link>
				</div>
				<div class="card-body">
					<table class="table table-bordered">
						<thead>
							<tr>
								<th> Название </th>
								<th> Описание </th>
								<th> Цена </th>
								<th> Создан </th>
								<th width="240px"> Действие </th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="product in products" :key="product.id">
								<td>{{product.name}}</td>
								<td>{{product.description}}</td>
								<td>{{product.price.toFixed(2)}}</td>
								<td>{{product.created_at.toLocaleString()}}</td>
								<td>
									<router-link :to="`/show/${product.id}`" class="btn btn-outline-info mx-1"> Показать </router-link>
									<router-link :to="`/edit/${product.id}`" class="btn btn-outline-success mx-1"> Редактировать </router-link>
									<button
										@click="handleDelete(product.id)"
										className="btn btn-outline-danger mx-1"
									> Удалить
									</button>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</layout-div>
</template>

<script>
import axios from 'axios'
import LayoutDiv from '../LayoutDiv.vue'
import Swal from 'sweetalert2'

export default {
	name: 'ProductList',
	components: {
		LayoutDiv,
	},
	data() {
		return {
			products: []
		}
	},
	created() {
		this.fetchProductList()
	},
	methods: {
		fetchProductList() {
			axios.get('/')
			.then(response => {
				for (let product of response.data) {
					product.price /= 100
					product.created_at = new Date(Date.parse(product.created_at))
				}
				this.products = response.data
				return response
			})
			.catch(error => {
				return error
			})
		},
		handleDelete(id) {
			Swal.fire({
				title: 'Удалить продукт?',
				text: "Это действие необратимо!",
				icon: 'warning',
				showCancelButton: true,
				confirmButtonColor: '#3085d6',
				cancelButtonColor: '#d33',
				confirmButtonText: 'Да, удалить!',
				cancelButtonText 'Отмена',
			}).then(result => {
				if (result.isConfirmed) {
					axios.delete(`/${id}`)
					.then( response => {
						Swal.fire({
							icon: 'success',
							title: 'Продукт удалён!',
							showConfirmButton: false,
							timer: 1500
						})
						this.fetchProductList()
						return response
					})
					.catch(error => {
						Swal.fire({
							icon: 'error',
							title: 'Ошибка!',
							showConfirmButton: false,
							timer: 1500
						})
						return error
					})
				}
			})
		}
	},
}
</script>
