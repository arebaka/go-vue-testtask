<template>
	<layout-div>
		<h2 class="text-center mt-5 mb-3"> Новый продукт </h2>
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
						<label htmlFor="name"> Название </label>
						<input
							v-model="product.name"
							type="text"
							class="form-control"
							id="name"
							name="name"/>
					</div>
					<div class="form-group">
						<label htmlFor="description"> Описание </label>
						<textarea
							v-model="product.description"
							class="form-control"
							id="description"
							rows="3"
							name="description"></textarea>
					</div>
					<div class="form-group">
						<label htmlFor="price"> Цена (в копейках) </label>
						<input type="number"
							v-model="product.price"
							class="form-control"
							id="price"
							name="price" />
					</div>
					<button
						@click="handleSave()"
						:disabled="isSaving"
						type="button"
						class="btn btn-outline-primary mt-3"
					> Сохранить продукт
					</button>
				</form>
			</div>
		</div>
	</layout-div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2'

import LayoutDiv from '../LayoutDiv.vue'

export default {
	name: 'ProductCreate',
	components: {
		LayoutDiv,
	},
	data() {
		return {
			product: {
				name: '',
				description: '',
				price: 0,
			},
			isSaving: false,
		}
	},
	methods: {
		handleSave() {
			this.isSaving = true
			axios.post('/', this.product)
			.then(response => {
				Swal.fire({
					icon: 'success',
					title: 'Продукт сохранён!',
					showConfirmButton: false,
					timer: 1500
				})
				this.isSaving = false
				this.product.name = ''
				this.product.description = ''
				this.product.price = 0
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
