<script setup>
import { ref, onMounted } from 'vue'

const axios = require('axios');
const res = ref()

onMounted(() => {
  axios.get('http://localhost:8080/ping')
  .then(function (response) {
    res.value = response.data;
  })
})
</script>

<template>
  <div name="app">{{ res }}</div>
</template>
