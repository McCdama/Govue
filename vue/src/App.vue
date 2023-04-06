<template>
    <div>
      <HashPass></HashPass>
      <Bcrypt></Bcrypt>
      <UUID />
    </div>
</template>

<script setup>

import Bcrypt from './components/Bcrypt.vue'
import HashPass from './components/HashPass.vue'
import UUID from './components/UUID.vue'

import { onMounted } from 'vue'

onMounted(() => {
    console.log("Main App mounted..")
  })

</script>