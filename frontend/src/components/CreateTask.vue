<template>
  <div>
    <h1 class="text-2xl font-bold">Create Task</h1>
    <VForm @submit="submitForm">
      <div class="mb-4">
        <label for="title" class="block text-gray-700">Title</label>
        <Field id="title" name="title" as="input" class="mt-1 block w-full" :rules="yup.string().required()" />
        <ErrorMessage name="title" />
      </div>

      <div class="mb-4">
        <label for="description" class="block text-gray-700">Description</label>
        <Field id="description" name="description" as="textarea" class="mt-1 block w-full" :rules="yup.string().required()" />
        <ErrorMessage name="description" />
      </div>

      <div class="mb-4">
        <label for="status" class="block text-gray-700">Status</label>
        <Field id="status" name="status" as="select" class="mt-1 block w-full" :rules="yup.string().required()">
          <option value="pending">Pending</option>
          <option value="completed">Completed</option>
        </Field>
        <ErrorMessage name="status" />
      </div>

      <button type="submit" class="btn">Create Task</button>
    </VForm>
  </div>
</template>

<script>
import { useMutation } from '@vue/apollo-composable';
import { gql } from '@apollo/client/core';
import { ref } from 'vue';
import { Field, ErrorMessage, Form as VForm } from 'vee-validate';
import * as yup from 'yup';

const CREATE_TASK = gql`
  mutation CreateTask($title: String!, $description: String!, $status: String!, $userId: String!) {
    insert_tasks(objects: { title: $title, description: $description, status: $status, userId: $userId }) {
      returning {
        id
        title
        description
        status
        created_at
        updated_at
      }
    }
  }
`;

export default {
  setup() {
    const userId = ref('user-123'); // Replace with actual user ID
    const { mutate } = useMutation(CREATE_TASK);

    const submitForm = async (values) => {
      await mutate({ ...values, userId: userId.value });
      // Redirect or notify user after success
    };

    return {
      submitForm,
    };
  },
  components: {
    Field,
    ErrorMessage,
    VForm,
  },
};
</script>

<style scoped>
.btn {
  @apply bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}
</style>
