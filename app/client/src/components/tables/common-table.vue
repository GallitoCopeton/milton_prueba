<template>
    <table>
        <thead>
        <tr>
            <th v-for="column in columns">{{ column.title }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in rows" :key="index">
            <td v-for="column in columns" :key="column.property">{{ item[column.property] }}</td>
        </tr>
        </tbody>
    </table>
</template>

<script lang="ts" setup>
import {Column} from "../../types/Column.ts";
import {PropType} from 'vue';

defineProps({
    columns: {
        type: Array as PropType<Column[]>,
        required: true,
    },
    rows: {
        type: Array as PropType<Record<string, any>[]>,
        required: true,
    }
});
</script>