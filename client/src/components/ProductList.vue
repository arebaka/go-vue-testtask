<template>
  <div class="container">
    <h2 class="text-center mt-5 mb-3">
      Продукты
    </h2>
    <div class="card overflow-hidden">
      <div class="card-header row">
        <div class="col-8">
          <router-link
            to="/add"
            class="btn btn-outline-primary"
          >
            Новый продукт
          </router-link>
        </div>
        <div class="col-4">
          <input
            v-model="search"
            type="text"
            class="form-control"
            placeholder="Поиск по названию..."
            @input="searchByName"
          >
        </div>
      </div>
      <div class="card-body">
        <table class="table table-bordered">
          <thead>
            <tr>
              <th> Название </th>
              <th> Описание </th>
              <th> Цена </th>
              <th> Создан </th>
              <th width="240px">
                Действие
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="product in products"
              :key="product.id"
            >
              <td>{{ product.name }}</td>
              <td>{{ product.description }}</td>
              <td>{{ product.price.toFixed(2) }}</td>
              <td>{{ product.created_at.toLocaleString() }}</td>
              <td>
                <router-link
                  :to="`/show/${product.id}`"
                  class="btn btn-outline-info mx-1"
                >
                  Показать
                </router-link>
                <router-link
                  :to="`/edit/${product.id}`"
                  class="btn btn-outline-success mx-1"
                >
                  Редактировать
                </router-link>
                <button
                  class="btn btn-outline-danger mx-1"
                  @click="deleteProduct(product.id)"
                >
                  Удалить
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Swal from 'sweetalert2'

import { getProductList, deleteProduct, searchProductsByName } from '@/service/product'

export default {
  name: 'ProductList',
  data() {
    return {
      products: [],
      nTotal: 0,
      search: '',
    }
  },
  created() {
    this.fetchProductList()
  },
  methods: {
    fetchProductList() {
      getProductList()
      .then(data => {
        for (let product of data.products) {
          product.price /= 100
          product.created_at = new Date(Date.parse(product.created_at))
        }
        this.products = data.products
        this.nTotal = data.n_total
      })
      .catch(() => {
        Swal.fire({
          icon: 'error',
          title: 'Ошибка получения списка продуктов!',
          showConfirmButton: false,
          timer: 1500
        })
      })
    },
    deleteProduct(id) {
      Swal.fire({
        title: 'Удалить продукт?',
        text: 'Это действие необратимо!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Да, удалить!',
        cancelButtonText: 'Отмена',
      }).then(result => {
        if (result.isConfirmed) {
          deleteProduct(id)
          .then(() => {
            Swal.fire({
              icon: 'success',
              title: 'Продукт удалён!',
              showConfirmButton: false,
              timer: 1500
            })
            this.fetchProductList()
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
    },
    searchByName() {
      if (this.search == '') {
		this.fetchProductList()
		return
      }
      searchProductsByName(this.search)
      .then(data => {
        for (let product of data.products) {
          product.price /= 100
          product.created_at = new Date(Date.parse(product.created_at))
        }
        this.products = data.products
        this.nTotal = data.n_total
      })
      .catch(() => {
        Swal.fire({
          icon: 'error',
          title: 'Ошибка получения списка продуктов!',
          showConfirmButton: false,
          timer: 1500
        })
      })
      return
    }
  },
}
</script>
