<template>
	<div class="container">
		<h2 class="text-center mt-5 mb-3"> Редактировать продукт </h2>
		<div class="card">
			<div class="card-header">
				<router-link to="/"
					class="btn btn-outline-info float-right"
				> Все продукты
				</router-link>
			</div>
			<div class="card-body">
				<form>
					<div class="form-group">
						<label htmlFor="name"> Назавание </label>
						<input
							v-model="product.name"
							type="text"
							class="form-control"
							id="name"
							name="name"
						/>
					</div>
					<div class="form-group">
						<label htmlFor="description"> Описание </label>
						<textarea
							v-model="product.description"
							class="form-control"
							id="description"
							rows="3"
							name="description"
						></textarea>
					</div>
					<div class="form-group">
						<label htmlFor="price"> Цена (в копейках) </label>
						<input type="number"
							v-model="product.price"
							class="form-control"
							id="price"
							name="price"
						/>
					</div>
					<button
						@click="save"
						:disabled="isSaving"
						type="button"
						class="btn btn-outline-primary mt-3"
					> Сохранить продукт
					</button>
				</form>
			</div>
		</div>
	</div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
	name: 'ProductEdit',
	data() {
		return {
			product: {
				id: this.$route.params.id,
				name: '',
				description: '',
				price: 0,
			},
			isSaving: false,
		}
	},
	created() {
		axios.get(`/${this.product.id}`)
		.then(response => {
			this.product.name = response.data.name
			this.product.description = response.data.description
			this.product.price = response.data.price
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
	methods: {
		save() {
			this.isSaving = true
			axios.put(`/${this.product.id}`, this.product)
			.then(response => {
				Swal.fire({
					icon: 'success',
					title: 'Данные о продукте обновлены!',
					showConfirmButton: false,
					timer: 1500
				})
				this.isSaving = false
				this.$router.push({ name: 'ProductList' })
				return response
			})
			.catch(error => {
				this.isSaving = false
				Swal.fire({
					icon: 'error',
					title: 'Ошибка!',
					showConfirmButton: false,
					timer: 1500
				})
				return error
			})
		},
	},
}
</script>
