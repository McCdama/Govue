<template>
  <div class="bg-white py-16 sm:py-24 lg:py-32">
    <div class="mx-auto max-w-7xl px-6 lg:px-8">
      <div class="max-w-2xl text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
        <h2 class="inline sm:block">Want secure your password?</h2>
      </div>
      <form class="mt-10 max-w-md">

        <div class="flex gap-x-4">
          <label for="password" class="sr-only">Your Password</label>
          <input v-model="password" id="password" name="password" type="text" autocomplete="off" class="min-w-0 flex-auto rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" placeholder="Enter your Password" />
         
          <label for="rounds" class="sr-only">Rounds</label>
          <input v-model="rounds" id="rounds" name="rounds" type="number" autocomplete="off" class="min-w-0 flex-auto rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" placeholder="Round"  min="10" max="31" />
        
          <button @click="generateBcrypt" type="button" class="flex-none rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Hash it!</button>
        </div>
        
        <div>
            <label for="bcryptedPass" class="mt-5 block text-sm font-medium leading-6 text-gray-900" v-if="bcryptedPass">Bcrypted Password</label>
            <div class="mt-2">
                <textarea  v-if="bcryptedPass" v-model="bcryptedPass" rows="4" name="bcryptedPass" id="bcryptedPass" class="block w-full rounded-md border-0 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:py-1.5 sm:text-sm sm:leading-6" />
            </div>
        </div>

      </form>
    </div>
  </div>
</template>

<script setup>
import axios from 'axios'
import { ref } from 'vue'

const password = ref('')
const rounds = ref(10)
const bcryptedPass = ref('')

function generateBcrypt() {

  axios({
  method: 'post',
  url: 'http://127.0.0.1:8080/bcrypt',
  data: {
    password: password.value,
    rounds: rounds.value,
  }, 
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
}) .then(function (res) {
    bcryptedPass.value = res.data.hash
  });
}

</script>