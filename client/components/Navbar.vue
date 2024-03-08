<template>
  <div class="h-20 w-screen bg-violet-900 flex">
    <div class="w-full flex justify-start items-center">
      <NuxtLink href="/" class="font-bold text-4xl text-white p-10">
        Serv<span class="text-fuchsia-500">U</span>
      </NuxtLink>
      <div class="flex flex-col ml-auto">
        <svg
          class="h-10 w-10 mr-5 sm:flex md:flex lg:hidden xl:hidden cursor-pointer"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="#ffffff"
          @click="toggle()"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h16"
          />
        </svg>
        <div
          class="flex flex-row gap-5 ml-auto font-semibold text-2xl text-white hidden lg:flex xl:flex"
        >
          <NuxtLink
            class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
            href="/"
            >Home</NuxtLink
          >
          <NuxtLink
            class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
            href="/request"
            >Request</NuxtLink
          >
          <NuxtLink
            v-show="admin"
            class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
            href="/admin"
            >Admin</NuxtLink
          >
          <NuxtLink
            v-show="login"
            class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
            href="/login"
            >Login</NuxtLink
          ><button
            v-show="logout"
            @click="logoutUser()"
            class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
            href="/login"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  </div>
  <div
    v-show="show"
    class="bg-violet-900 w-full justify-center p-3 flex flex-col gap-5 ml-auto font-semibold text-2xl text-white mr-5 lg:hidden xl:hidden"
  >
    <NuxtLink
      class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
      href="/"
      >Home</NuxtLink
    >
    <NuxtLink
      class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
      href="/request"
      >Request</NuxtLink
    >
    <NuxtLink
      class="hover:bg-violet-700 p-3 rounded duration-150 ease-in-out"
      href="/login"
      >Login</NuxtLink
    >
  </div>
</template>

<script setup>
const show = ref(false);
const admin = ref(false);
const login = ref(true);
const logout = ref(false);

const toggle = () => {
  show.value = !show.value;
};

const logoutUser = async() => {
  const res = await fetch("http://localhost:8000/logout", {
    method: "POST",
    credentials: "include",
    headers: { "Content-Type": "application/json" },
  });
  if (res.ok) {
    login.value = true;
    logout.value = false;
    admin.value = false;
    console.log("logged out");
  }
}

onMounted(async () => {
  const res = await fetch("http://localhost:8000/session", {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  if (res.ok) {
    login.value = false;
    logout.value = true;
    if (data.role === "admin") {
      admin.value = true;
      console.log("im admin");
    } else {
      console.log("not admin");
    }
  }
});
</script>
