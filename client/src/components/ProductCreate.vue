<template>
  <div class="container">
    <h2 class="text-center mt-5 mb-3">
      Новый продукт
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
            <label htmlFor="name"> Название </label>
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
            class="btn btn-outline-primary mt-3"
            :disabled="isSaving"
            type="button"
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

import { addProduct } from '@/service/product'

export default {
  name: 'ProductCreate',
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
    save() {
      this.isSaving = true
      addProduct(this.product)
      .then(() => {
        Swal.fire({
          icon: 'success',
          title: 'Продукт сохранён!',
          showConfirmButton: false,
          timer: 1500
        })
        this.isSaving = false
        this.$router.push({ name: 'ProductList' })
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
