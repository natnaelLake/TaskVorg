<template>
  <nav class="bg-gray-800 p-4 text-white">
    <ul class="flex space-x-4">
      <li>
        <router-link to="/" class="hover:underline">Task List</router-link>
      </li>
      <li>
        <router-link to="/create" class="hover:underline">Create Task</router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'Routes',
  setup(props) {
    // You can use Composition API here if needed
  }
};
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
