<template>
  <div class="w-[100%] h-full flex justify-center items-cente">
    <div class="w-full h-full overflow-auto flex justify-center items-center">
      <Table
        :headers="[
          'ID',
          'Problem',
          'Date',
          'User',
          'Email',
          'Phone',
          'Severity',
          'Category',
        ]"
        :records="fetchedData"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: "dashboard",
});

const fetchedData = ref([]);

onMounted(async () => {
  try {
    const res = await fetch("http://localhost:8000/admin/requests", {
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
