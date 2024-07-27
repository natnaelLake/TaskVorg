<template>
  <div class="container mx-auto p-6">
    <h1 class="text-3xl font-bold mb-6 text-center text-gray-800">
      Create Task
    </h1>
    <form
      @submit.prevent="submitForm"
      class="max-w-lg mx-auto bg-white p-6 rounded-lg shadow-md"
    >
      <div class="mb-4">
        <label for="title" class="block text-sm font-medium text-gray-700"
          >Title</label
        >
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
        <label for="description" class="block text-sm font-medium text-gray-700"
          >Description</label
        >
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
        <label for="status" class="block text-sm font-medium text-gray-700"
          >Status</label
        >
        <input
          v-model="status"
          id="status"
          type="text"
          class="form-input"
          placeholder="Enter task status"
          required
        />
      </div>
      <div class="mb-4">
        <label for="image" class="block text-sm font-medium text-gray-700"
          >Image</label
        >
        <input
          id="image"
          type="file"
          @change="handleFileChange"
          class="form-input"
        />
      </div>
      <button type="submit" class="btn btn-primary w-full">Create Task</button>
    </form>
    <div v-if="error" class="mt-4 text-center text-red-500">
      {{ error.message }}
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import { gql } from "@apollo/client/core";
import axios from "axios";
import { useRouter } from 'vue-router';

const CREATE_TASK = gql`
  mutation CreateTask(
    $title: String!
    $description: String!
    $status: String!
    $image_url: String!
    $user_id: Int!
  ) {
    insert_tasks(
      objects: {
        title: $title
        description: $description
        status: $status
        user_id: $user_id
        image_url: $image_url
      }
    ) {
      returning {
        id
        title
        description
        status
        created_at
        updated_at
        image_url
      }
    }
  }
`;

export default {
  setup() {
    const router = useRouter();
    const title = ref("");
    const description = ref("");
    const status = ref("");
    const user_id = ref(1);
    const image = ref(null);
    const error = ref(null);
    const { mutate } = useMutation(CREATE_TASK);

    const handleFileChange = (event) => {
      image.value = event.target.files[0];
    };

    const submitForm = async () => {
      try {
        let imageUrl = "";
        if (image.value) {
          const formData = new FormData();
          formData.append("file", image.value);
          const response = await axios.post(
            "http://localhost:8081/uploads",
            formData,
            {
              headers: {
                "Content-Type": "multipart/form-data",
              },
            }
          );
          imageUrl = response.data.url;
        }
        console.log("+++++++++++++++++++++", imageUrl);
        await mutate({
          title: title.value,
          description: description.value,
          status: status.value,
          image_url: imageUrl,
          user_id: user_id.value
        });
        router.push('/');
      } catch (e) {
        error.value = e.message;
        console.error("Submit Error:", e);
      }
    };

    return {
      title,
      description,
      status,
      handleFileChange,
      submitForm,
      error,
    };
  },
};
</script>

<style scoped>
.form-input {
  @apply border border-gray-300 p-3 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent;
}
.btn-primary {
  @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}
</style>
