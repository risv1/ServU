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
        <label for="customer_name" class="text-white">Customer Name</label>
        <input
        v-model="customer_name"
          type="text"
          id="customer_name"
          name="customer_name"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        />
      </div>
      <div class="flex flex-col gap-3 w-full">
        <label for="contact_email" class="text-white">Contact Email</label>
        <input
        v-model="contact_email"
          type="email"
          id="contact_email"
          name="contact_email"
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
      <div class="flex flex-col gap-3 w-full">
        <label for="description" class="text-white">Description</label>
        <textarea
          v-model="description"
          id="description"
          name="description"
          rows="3"
          class="w-full bg-violet-900 focus:outline-none border-fuchsia-600 border-2 text-white rounded p-2"
        ></textarea>
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
import { ref } from 'vue';
import { z, object } from 'zod';

const problem = ref('');
const since_when = ref('');
const customer_name = ref('');
const contact_email = ref('');
const phone_number = ref('');
const severity = ref('');
const category = ref('');
const description = ref('');

const submitSchema = object({
  problem: z.string().min(10),
  since_when: z.string(),
  customer_name: z.string().min(3),
  contact_email: z.string().email(),
  phone_number: z.string().min(10).max(10),
  severity: z.string(),
  category: z.string().min(3),
  description: z.string().min(10),
});

const onSubmit = (event: Event) => {
  event.preventDefault();
  try {
    const validatedData = submitSchema.parse({
      problem: problem.value,
      since_when: since_when.value,
      customer_name: customer_name.value,
      contact_email: contact_email.value,
      phone_number: phone_number.value,
      severity: severity.value,
      category: category.value,
      description: description.value,
    });
    if (validatedData) {
      console.log('Form submitted successfully!', problem, since_when, customer_name, contact_email, phone_number, severity, category, description);
    } else {
      console.error('Validation error:', validatedData);
    }
  } catch (error) {
    console.error('Unexpected error:', error);
  }
};

</script>