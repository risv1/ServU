<template>
  <div class="flex w-full h-full justify-center items-center">
    <div
      class="flex-col p-5 gap-5 w-2/3 h-3/5 border-fuchsia-600 shadow-md shadow-fuchsia-500 border-2 flex justify-center items-center"
    >
      <h1 class="text-4xl text-fuchsia-500 font-bold">Admin Console</h1>
      <p class="p-5 text-3xl text-fuchsia-300 font-semibold">
        Console to manage users and all incoming service requests
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: "dashboard",
});
</script>
