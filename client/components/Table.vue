<template>
    <div class="w-full h-full p-5">
      <table class="w-full border-fuchsia-600 border-2">
        <thead>
          <tr class="text-white text-3xl font-semibold">
            <th v-for="(header, index) in headers" :key="index" class="border-fuchsia-600 border-2 px-4 py-2">{{ header }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(record, rowIndex) in records" :key="rowIndex">
            <td v-for="(value, colIndex) in record" :key="colIndex" class="text-white text-2xl font-semibold border-fuchsia-600 border-2 px-4 py-2">{{ value }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script setup lang="ts">
  import { defineProps } from 'vue';
  
  const props = defineProps({
    headers: {
      type: Array,
      required: true
    },
    records: {
      type: Array,
      required: true
    }
  });
  </script>

  