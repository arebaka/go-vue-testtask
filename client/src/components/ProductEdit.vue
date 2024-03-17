<template>
  <div class="container">
    <h2 class="text-center mt-5 mb-3">
      Редактировать продукт
    </h2>
    <div class="card">
      <div class="card-header">
        <router-link
          to="/"
          class="btn btn-outline-info float-right"
        >
          Все продукты
        </router-link>
      </div>
      <div class="card-body">
        <form>
          <div class="form-group">
            <label htmlFor="name"> Назавание </label>
            <input
              id="name"
              v-model="product.name"
              type="text"
              class="form-control"
              name="name"
            >
          </div>
          <div class="form-group">
            <label htmlFor="description"> Описание </label>
            <textarea
              id="description"
              v-model="product.description"
              class="form-control"
              name="description"
              rows="3"
            />
          </div>
          <div class="form-group">
            <label htmlFor="price"> Цена (в копейках) </label>
            <input
              id="price"
              v-model="product.price"
              type="number"
              class="form-control"
              name="price"
            >
          </div>
          <button
            :disabled="isSaving"
            type="button"
            class="btn btn-outline-primary mt-3"
            @click="save"
          >
            Сохранить продукт
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Swal from 'sweetalert2'

import { getProduct, updateProduct } from '@/service/product'

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
    getProduct(this.product.id)
    .then(data => {
      this.product.name = data.name
      this.product.description = data.description
      this.product.price = data.price
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
      updateProduct(this.product.id, this.product)
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
