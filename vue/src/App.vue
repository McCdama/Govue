<template>
  <form >
    <div class="space-y-12 sm:space-y-16">
      <div>
        <h2 class="text-base font-semibold leading-7 text-gray-900">Gibberish Tool</h2>
        <p class="mt-1 max-w-2xl text-sm leading-6 text-gray-600">A hash tool to ensure immutability</p>
        <div class="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
          <label for="password" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5">Your
            Password</label>
          <div class="mt-2 sm:col-span-2 sm:mt-0">
            <input type="text" autocomplete="off" v-model="rawPassword"
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6" />
          </div>

          <!-- invisible -->
          <label v-if="hashedPassword && rawPassword" for="hashed-password" class="block text-sm font-medium leading-6 text-gray-900 sm:pt-1.5"> 
            Hashed Password</label>
          <div class="mt-2 sm:col-span-2 sm:mt-0" v-if="hashedPassword && rawPassword">
            <input type="text" autocomplete="off" v-model="hashedPassword"
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6" />
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6 flex items-center justify-end gap-x-6">
      <button type="button" @click="generateHash" class="inline-flex items-center gap-x-2 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
        <CheckCircleIcon class="-ml-0.5 h-5 w-5" aria-hidden="true" />
        GENERATE
      </button>
    </div>


    <div>
      <Bcrypt></Bcrypt>
    </div>
  </form>
</template>

<script setup>
import { CheckCircleIcon } from '@heroicons/vue/20/solid'
import axios from 'axios'
import Bcrypt from './components/Bcrypt.vue'

import { ref, onMounted } from 'vue'

const rawPassword = ref('')
const hashedPassword = ref('')

function generateHash() {

  axios({
  method: 'post',
  url: 'http://127.0.0.1:8080/hash',
  data: {
    rawPassword: rawPassword.value,
  }, 
    headers: {
      'Content-Type': 'multipart/form-data'
    }
}) .then(function (res) {
    hashedPassword.value = res.data.hash
  });
}

onMounted(() => {
  console.log("App mounted..")
})

</script>