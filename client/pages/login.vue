<template>
  <div class="bg-violet-950 flex justify-center items-center w-full h-screen">
    <NuxtErrorBoundary>
      <LoginForm />
      <template #error="{error}">
      <div class="w-4/5 h-3/5 flex justify-center items-center bg-violet-800 border-2 border-fuchsia-600 shadow-md shadow-fuchsia-500">
        <div class="flex justify-center items-center w-full flex-col gap-5 mt-20">
          <h1 class="text-white text-4xl font-bold">Validation error!</h1>
          <h1 class="text-fuchsia-400 text-2xl ">{{ error }}</h1>
          <button @click="error.value = null" class="rounded p-3 bg-fuchsia-600 text-white hover:shadow-fuchsia-500 shadow-md">Go back</button>
        </div>
      </div>
    </template>
    </NuxtErrorBoundary>
   
  </div>
</template>

<script setup>
useHead({
  title: "ServU / Login",
});

onMounted(async()=>{
  const res = await fetch("http://localhost:8000/session", {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json" },
  });
  if(res.ok){
    navigateTo("/")
  }
  })
</script>
