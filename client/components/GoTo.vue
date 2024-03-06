<template>
  <div
    class="w-[45%] h-2/3 rounded-xl bg-black bg-opacity-50 border-2 border-fuchsia-600 p-5 shadow-md shadow-fuchsia-500 flex"
  >
    <div class="w-full flex-col flex justify-center items-center gap-3">
        <h1 class="text-4xl text-fuchsia-600 mb-5 font-bold">Go to</h1>
      <div
        v-for="(route, index) in routes"
        :key="index"
        @click="handleRoute(route.route)"
        class="cursor-pointer font-semibold w-4/5 mb-2 bg-violet-900 hover:shadow-fuchsia-500 shadow-md duration-150 ease-in-out h-24 rounded-lg flex justify-center items-center hover:bg-violet-800 text-white text-2xl"
      >
        {{ route.name }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const handleRoute = (route: string) => {
  navigateTo(route);
};
const routes = [
  {
    name: "Home",
    route: "/",
  },
  {
    name: "Request",
    route: "/request",
  },
];
</script>
