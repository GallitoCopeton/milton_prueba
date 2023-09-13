<template>
    <common-table :rows="products" :columns="columns"></common-table>
</template>

<script setup lang="ts">

import CommonTable from "../components/tables/common-table.vue";
import {Column} from "../types/Column.ts";
import {onMounted} from "vue";
import {useProducts} from "../compositions/useProducts.ts";

const {products, fetchProducts} = useProducts();

onMounted(async () => {
    await fetchProducts()
})

const columns: Column[] = [
    {
        title: "ID",
        property: "id",
    },
    {
        title: "Title",
        property: "name",
    },
    {
        title: "Price",
        property: "price",
    },
    {
        title: "Rank",
        property: "rank",
    },

]

</script>


<style scoped>

</style>