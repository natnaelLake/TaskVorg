<template>
  <div class="container mx-auto p-6">
    <h1 class="text-3xl font-bold mb-6 text-center text-gray-800">Update Task</h1>
    <form @submit.prevent="handleUpdate" class="max-w-lg mx-auto bg-white p-6 rounded-lg shadow-md">
      <div class="mb-4">
        <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
        <input
          v-model="title"
          id="title"
          type="text"
          class="form-input"
          placeholder="Enter task title"
          required
        />
      </div>
      <div class="mb-4">
        <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
        <textarea
          v-model="description"
          id="description"
          class="form-input"
          placeholder="Enter task description"
          rows="4"
          required
        ></textarea>
      </div>
      <div class="mb-4">
        <label for="status" class="block text-sm font-medium text-gray-700">Status</label>
        <input
          v-model="status"
          id="status"
          type="text"
          class="form-input"
          placeholder="Enter task status"
          required
        />
      </div>
      <button type="submit" class="btn btn-primary w-full">Update Task</button>
    </form>
    <div v-if="error" class="mt-4 text-center text-red-500">{{ error.message }}</div>
  </div>
</template>

<script>
import { ref, watchEffect } from "vue";
import { useMutation, useQuery } from "@vue/apollo-composable";
import { gql } from "@apollo/client/core";
import { useRoute, useRouter } from "vue-router";

const GET_TASK_BY_ID = gql`
  query GetTaskById($id: Int!) {
    tasks_by_pk(id: $id) {
      id
      title
      description
      status
    }
  }
`;

const UPDATE_TASK = gql`
  mutation UpdateTask(
    $id: Int!
    $title: String!
    $description: String!
    $status: String!
  ) {
    update_tasks_by_pk(
      pk_columns: { id: $id }
      _set: { title: $title, description: $description, status: $status }
    ) {
      id
    }
  }
`;

export default {
  name: "UpdateTask",
  setup() {
    const route = useRoute();
    const router = useRouter();
    const taskId = route.params["id"];
    const { result, loading, error } = useQuery(GET_TASK_BY_ID, { id: taskId });
    const { mutate: updateTask } = useMutation(UPDATE_TASK);
    const title = ref("");
    const description = ref("");
    const status = ref("");

    watchEffect(() => {
      if (result.value && result.value.tasks_by_pk) {
        const task = result.value.tasks_by_pk;
        title.value = task.title;
        description.value = task.description;
        status.value = task.status;
      }
    });

    const handleUpdate = async () => {
      try {
        await updateTask({
          id: taskId,
          title: title.value,
          description: description.value,
          status: status.value,
        });
        router.push("/");
      } catch (e) {
        console.error("Update Error:", e);
      }
    };

    return {
      title,
      description,
      status,
      handleUpdate,
      error,
      loading,
    };
  },
};
</script>

<style scoped>
.container {
  max-width: 800px;
}

.form-input {
  @apply border border-gray-300 p-3 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent;
}

.btn-primary {
  @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}
</style>