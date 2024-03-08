<template>
  <div
    class="lg:w-2/3 xl:w-2/3 sm:w-4/5 h-full mt-10 md:w-1/2 p-5 rounded flex flex-col justify-center gap-5 border-2 bg-black bg-opacity-50 items-center border-fuchsia-600 shadow-lg shadow-fuchsia-500"
  >
    <h1 class="text-fuchsia-600 text-4xl p-5 font-bold">Service</h1>
    <form @submit="onSubmit($event)" class="w-full">
      <div class="flex flex-col gap-3 w-full">
        <label for="problem" class="text-white">Problem Details</label>
        <textarea
          v-model="problem"
          id="problem"
          name="problem"
          rows="5"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        ></textarea>
      </div>
      <div class="flex flex-col gap-3 w-full">
        <label for="since_when" class="text-white">Since When</label>
        <input
          v-model="since_when"
          type="date"
          id="since_when"
          name="since_when"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        />
      </div>
      <div class="flex flex-col gap-3 w-full">
        <label for="phone_number" class="text-white">Phone Number</label>
        <input
          v-model="phone_number"
          type="tel"
          id="phone_number"
          name="phone_number"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        />
      </div>
      <div class="flex flex-col gap-3 w-full">
        <label for="severity" class="text-white">Severity</label>
        <select
          v-model="severity"
          id="severity"
          name="severity"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        >
          <option value="low">Low</option>
          <option value="medium">Medium</option>
          <option value="high">High</option>
        </select>
      </div>
      <div class="flex flex-col gap-3 w-full">
        <label for="category" class="text-white">Category</label>
        <input
          v-model="category"
          type="text"
          id="category"
          name="category"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        />
      </div>
      <button
        type="submit"
        class="bg-fuchsia-600 text-white font-bold py-2 px-4 rounded mt-5"
      >
        Submit
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { z, object } from "zod";

const problem = ref("");
const since_when = ref("");
const phone_number = ref("");
const severity = ref("");
const category = ref("");

const submitSchema = object({
  problem: z.string().min(10),
  since_when: z.string(),
  phone_number: z.string().min(10).max(10),
  severity: z.string(),
  category: z.string().min(3),
});

const onSubmit = async(event: Event) => {
  event.preventDefault();
  try {
    const validatedData = submitSchema.parse({
      problem: problem.value,
      since_when: since_when.value,
      phone_number: phone_number.value,
      severity: severity.value,
      category: category.value,
    });
    if (validatedData) {
      const res = await fetch("http://localhost:8000/request", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(validatedData),
      });
      if (!res.ok) {
        throw createError({
          message: "Failed to submit form",
          statusCode: res.status,
        });
      }
      console.log(
        "Form submitted successfully!",
        problem,
        since_when,
        phone_number,
        severity,
        category
      );
    } else {
      throw createError({
        message: "Invalid form data",
        statusCode: 400,
      });
    }
  } catch (error) {
    throw createError({
      message: "Unexpected error",
      statusCode: 500,
    });
  }
};
</script>
