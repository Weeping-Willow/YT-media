<template>
    <nav class="fixed top-0 left-0 right-0 z-50 bg-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex-shrink-0">
           <h1 class="text-lg">Media Magic</h1>
          </div>
          <div class="hidden md:block">
            <div class="ml-10 flex items-baseline space-x-4">
              <NavLink to="/" label="Home" />
              <NavLink to="/upload" label="Upload" />
              <NavLink to="/statistics" label="Statistics" />
            </div>
          </div>
        </div>
      </div>
    </nav>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import NavLink from './NavLink.vue';
  
  export default defineComponent({
    name: 'MainNav',
    components: {
      NavLink,
    },
  });
  </script>