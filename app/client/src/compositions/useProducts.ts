import {ref} from 'vue';
import axios from 'axios';
import Product from "../types/Product.ts";
import {BACKEND_V1_API_URL, HOST, PRODUCTS_API_URL} from "../constants.ts";

export function useProducts() {
    const products = ref<Product[]>([]);
    const API_URL = `${HOST}${BACKEND_V1_API_URL}${PRODUCTS_API_URL}`;

    async function fetchProducts() {
        try {
            const response = await axios.get(API_URL);
            products.value = response.data;
        } catch (error) {
            console.error('Error fetching products:', error);
        }
    }

    return {products, fetchProducts};
}