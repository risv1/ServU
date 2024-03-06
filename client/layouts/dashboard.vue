<template>
  <Navbar />
  <div class="w-screen h-screen bg-violet-950 flex">
    <div class="flex flex-row w-full gap-3">
      <div class="w-[25%] h-full flex flex-col gap-5 items-center justify-center">
        <h1 class="text-5xl font-bold text-fuchsia-600 pb-24">Dashboard</h1>
        <div
          v-for="(route, index) in routes"
          :key="index"
          @click="handleRoute(route.route)"
          class="cursor-pointer font-semibold w-4/5 mb-2 bg-violet-900 hover:shadow-fuchsia-500 shadow-md duration-150 ease-in-out h-24 rounded-lg flex flex-col justify-center items-center hover:bg-violet-800 text-white text-2xl"
        >
          <h1 class="text-3xl text-white font-bold">Module</h1>
          <h2 class="text-xl text-white font-semibold">{{ route.name }}</h2>
        </div>
      </div>
      <div class="w-[70%] h-full border-l-2 border-fuchsia-600 p-5">
        <slot />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const routes = [
  { name: "Home", route: "/admin" },
  { name: "Requests", route: "/admin/requests" },
  { name: "Users", route: "/admin/users" },
];

const handleRoute = (route: string) => {
  navigateTo(route);
};
</script>
