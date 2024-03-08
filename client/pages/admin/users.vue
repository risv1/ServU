<template>
  <div class="overflow-auto">
    <Table
      :headers="['ID', 'Name', 'Email', 'Password', 'Role']"
      :records="fetchedData"
    />
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: "dashboard",
});

const fetchedData = ref([]);

onMounted(async () => {
  try {
    const res = await fetch("http://localhost:8000/admin/users", {
      method: "GET",
      credentials: "include",
      headers: { "Content-Type": "application/json" },
    });
    if (res.status === 200) {
      const data = await res.json();
      console.log(data);
      fetchedData.value = data;
    }
  } catch (error) {
    console.log(error);
  }
});
</script>
