<template>
  <div class="bg-violet-950 flex justify-center items-center w-full h-full p-5">
    <NuxtErrorBoundary>
      <ServiceForm />
      <template #error="{error}">
      <div class="w-4/5 flex justify-center items-center bg-violet-800 border-2 border-fuchsia-600 shadow-md shadow-fuchsia-500">
        <div class="flex justify-center items-center w-full">
          <h1>Validation error!</h1>
          <h1>{{ error }}</h1>
          <button @click="error.value = null">Go back</button>
          <h1></h1>
        </div>
      </div>
    </template>
    </NuxtErrorBoundary>
  </div>
</template>

<script setup>
useHead({
  title: "ServU / Request",
});
</script>
