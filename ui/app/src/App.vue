<script setup>
import { RouterLink, RouterView } from 'vue-router'
//import HelloWorld from './components/HelloWorld.vue'
//import Sidebar from './components/Sidebar.vue'
import AppLayout from '@/layout/AppLayout.vue'
</script>

<template>
   <!--
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>


         <Sidebar />
         <RouterView class="w-full sm:w-2/3 md:w-3/4 pt-1 px-2" />

   <div class="container mx-auto">
      <section class="main">
         <AppLayout />
      </section>
   </div>
  -->
   <router-view />
</template>


<style>
.main {
   @apply flex flex-row flex-wrap py-4;
}
</style>
