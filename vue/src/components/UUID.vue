<template>
    {{ id }}
  </template>
  
  <script setup>
  import axios from 'axios'
  
  import { onBeforeMount, ref } from 'vue'
  
  const id = ref('')

onBeforeMount(() => {
    getUUID()
  })

  
  function getUUID() {
    axios({
    method: 'post',
    url: 'http://127.0.0.1:8080/uuid',
    data: {
        numberRequested: 1,
    }, 
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
  }) .then(function (res) {
    id.value = res.data.uuid
    });
  }
  
  </script>