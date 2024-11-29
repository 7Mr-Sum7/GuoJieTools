<template>
  <div>
    <NavSide class="nav-side"/>
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import NavSide from './components/NavSide.vue'
</script>

<style scoped>
.nav-side {
  width: 100%;
  margin-bottom: 1%;
}
</style>
