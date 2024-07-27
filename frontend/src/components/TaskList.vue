<template>
  <div class="container mx-auto p-4">
    <h1 class="text-4xl font-bold mb-6 text-center">Task List</h1>
    <div v-if="loading" class="text-center text-lg">Loading...</div>
    <div v-else>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="task in tasks"
          :key="task.id"
          class="task bg-white p-4 rounded-lg shadow-lg flex flex-col"
        >
          <img
            :src="`http://localhost:8081/uploads/${task.image_url}`" 
            alt="Task Image"
            class="w-full h-40 object-cover mb-4 rounded-md"
          />
          <router-link :to="`/task-list/${task.id}`" class="flex-grow">
            <h2 class="text-2xl font-semibold text-gray-800 hover:text-blue-600 transition-colors mb-2">{{ task.title }}</h2>
            <p class="text-gray-600 mb-2">{{ task.description }}</p>
            <p class="text-gray-500 mb-2">Status: {{ task.status }}</p>
            <p class="text-gray-400 text-sm mb-2">Created at: {{ new Date(task.created_at).toLocaleString() }}</p>
            <p class="text-gray-400 text-sm">Updated at: {{ new Date(task.updated_at).toLocaleString() }}</p>
          </router-link>
          <div class="flex justify-between mt-4">
            <router-link :to="`/update-task/${task.id}`">
              <button class="btn btn-primary flex items-center">
                <i class="fas fa-edit mr-2"></i> Update
              </button>
            </router-link>
            <button
              @click="handleDelete(task.id)"
              class="btn btn-danger flex items-center"
            >
              <i class="fas fa-trash-alt mr-2"></i> Delete
            </button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="error" class="text-center text-red-500 mt-4">{{ error.message }}</div>
  </div>
</template>


<script>
import { ref, watchEffect,onMounted } from "vue";
import { useMutation, useQuery } from "@vue/apollo-composable";
import { gql } from "@apollo/client/core";

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
      image_url
    }
  }
`;
const DELETE_TASK = gql`
  mutation DeleteTask($id: Int!) {
    delete_tasks_by_pk(id: $id) {
      id
    }
  }
`;

export default {
  name: "Task List",
  setup() {
    const { result, loading, error, refetch } = useQuery(GET_TASKS);
    const tasks = ref([]);
    const { mutate: deleteTask } = useMutation(DELETE_TASK);

    const handleDelete = async (id) => {
      try {
        await deleteTask({ id });
        // Refresh the task list
        refetch();
      } catch (e) {
        console.error("Delete Error:", e);
      }
    };
    onMounted(()=>{
      refetch();
    });
    watchEffect(() => {
      if (error.value) {
        console.log("Error:", error.value);
        return;
      }
      if (result.value && result.value.tasks) {
        tasks.value = result.value.tasks;
      }
    });

    return {
      tasks,
      loading,
      error,
      handleDelete,
    };
  },
};
</script>

<style scoped>
.container {
  max-width: 1200px;
}

.btn {
  @apply flex items-center bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}

.btn-primary {
  @apply bg-blue-500 hover:bg-blue-700;
}

.btn-danger {
  @apply bg-red-500 hover:bg-red-700;
}

.task {
  @apply border border-gray-200;
  @apply transition-transform transform hover:scale-105;
}

.task img {
  @apply rounded-md;
}

.task h2 {
  @apply text-xl font-semibold;
}

.task p {
  @apply text-gray-600;
}
</style>
