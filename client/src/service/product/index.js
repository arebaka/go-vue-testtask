import axios from 'axios'

export function addProduct(details) {
	return axios.post(`/product/`, details)
	.then(response => {
		return response.data
	})
}

export function getProduct(id) {
	return axios.get(`/product/${id}`)
	.then(response => {
		return response.data
	})
}

export function updateProduct(id, details) {
	return axios.put(`/product/${id}`, details)
	.then(response => {
		return response.data
	})
}

export function deleteProduct(id) {
	return axios.delete(`/product/${id}`)
	.then(response => {
		return response.data
	})
}

export function getProductList() {
	return axios.get('/product/')
	.then(response => {
		return response.data
	})
}
