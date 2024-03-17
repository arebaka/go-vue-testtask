<template>
  <div class="container">
    <h2 class="text-center mt-5 mb-3">
      Продукт
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
        <b class="text-muted"> Название: </b>
        <p>{{ product.name }}</p>
        <b class="text-muted"> Описание: </b>
        <p>{{ product.description }}</p>
        <b class="text-muted"> Цена: </b>
        <p>{{ product.price.toFixed(2) }}</p>
        <b class="text-muted"> Создан: </b>
        <p>{{ product.created_at.toLocaleString() }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import Swal from 'sweetalert2'

import { getProduct } from '@/service/product'

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
    getProduct(this.product.id)
    .then(data => {
      this.product.name = data.name
      this.product.description = data.description
      this.product.price = data.price / 100
      this.product.created_at = new Date(Date.parse(data.created_at))
    })
    .catch(() => {
      Swal.fire({
        icon: 'error',
        title: 'Ошибка получения данных о продукте!',
        showConfirmButton: false,
        timer: 1500
      })
    })
  },
  methods: {},
}
</script>
