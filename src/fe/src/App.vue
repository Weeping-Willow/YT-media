<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import TheNavBar from './components/Nav/NavBar.vue'
</script>

<template>
  <TheNavBar/>
  <RouterView class="mt-10"></RouterView>
</template>