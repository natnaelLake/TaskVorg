<template>
  <div>
    <h1 class="text-2xl font-bold">Task List</h1>
    <div class="my-4">
      <router-link to="/create" class="btn btn-primary">Create Task</router-link>
    </div>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <div v-for="task in tasks" :key="task.id" class="task">
        <h2 class="text-xl font-semibold">{{ task.title }}</h2>
        <p>{{ task.description }}</p>
        <p>Status: {{ task.status }}</p>
        <p>Created at: {{ new Date(task.created_at).toLocaleString() }}</p>
        <p>Updated at: {{ new Date(task.updated_at).toLocaleString() }}</p>
      </div>
    </div>
    <div v-if="error">{{ error.message }}</div>
  </div>
</template>

<script>
import { ref, watchEffect } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import { gql, } from '@apollo/client/core';

const GET_TASKS = gql`
  query GetTasks {
    tasks(order_by: { created_at: desc }) {
      id
      title
      description
      status
      created_at
      updated_at
      user_id
    }
  }
`;

export default {
  name:'App',
  setup() {
    const { result, loading, error } = useQuery(GET_TASKS);
    const tasks = ref([]);

    watchEffect(() => {
      if (error.value) {
        console.log('Error:', error.value);
        return;
      }
      if (result.value && result.value.tasks) {
        console.log('-----------------')
        tasks.value = result.value.tasks;
      }
    });

    return {
      tasks,
      loading,
      error,
    };
  },
};
</script>

<style scoped>
.btn {
  @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}
.task {
  @apply border p-4 mb-4 rounded shadow;
}
</style>