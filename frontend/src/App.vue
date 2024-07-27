<script>
import { RouterView } from "vue-router";
import Routes from "./components/Routes.vue";
export default {
  name: "App",
  components: {
    Routes,
  },
};
</script>

<template>
  <div class="container mx-auto p-4">
    <Routes />
    <!-- <RouterView /> -->
    <router-view />
  </div>
</template>
