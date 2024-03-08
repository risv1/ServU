<template>
  <div
    class="lg:w-1/2 xl:w-1/2 h-2/3 p-5 rounded flex flex-col justify-center gap-5 border-2 bg-black bg-opacity-50 items-center border-fuchsia-600 shadow-lg shadow-fuchsia-500"
  >
    <h1 class="text-fuchsia-600 text-4xl p-5 font-bold">Login</h1>
    <form @submit="onSubmit($event)">
      <div class="flex flex-col gap-5">
        <input
          v-model="email"
          type="email"
          class="p-3 rounded border-fuchsia-600 focus:outline-none bg-violet-950 border-2 text-white"
          placeholder="Email"
        />
        <input
          v-model="password"
          type="password"
          class="p-3 rounded border-fuchsia-600 focus:outline-none bg-violet-950 border-2 text-white"
          placeholder="Password"
        />
        <button
          type="submit"
          class="p-3 bg-fuchsia-700 rounded text-white font-bold hover:bg-fuchsia-600"
        >
          Login
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue';
import { z, object } from 'zod';

const email = ref('');
const password = ref('');

const submitSchema = object({
    email: z.string().email(),
    password: z.string().min(8),
})

const onSubmit = async(event: Event) => {
    event.preventDefault()
    try {
        const validatedData = submitSchema.parse({
            email: email.value,
            password: password.value
        })
        if (validatedData) {
          const res = await fetch("http://localhost:8000/login", {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              email: email.value,
              password: password.value,
            }),
          })
          if(!res.ok) {
            throw createError({
              message: "Unable to submit form!",
              statusCode: 400
            })
          }
            console.log("Form submitted successfully!", res.body);
            navigateTo("/");
        } else {
            throw createError({
                message: "Form validation failed!",
                statusCode: 400
            })
        }
    } catch (error) {
        throw createError({
            message: "Unable to submit form!",
            statusCode: 400
        })
    }
}

</script>
