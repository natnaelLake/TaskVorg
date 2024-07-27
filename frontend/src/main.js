import "./style.css";
import App from "./App.vue";
import { apolloClient } from "./apollo";
import { createRouter, createWebHistory } from "vue-router";
import TaskList from "./components/TaskList.vue";
import CreateTask from "./components/CreateTask.vue";
import TaskDetails from "./components/TaskDetails.vue";
import UpdateTasks from "./components/UpdateTasks.vue";
import { createApp, provide, h } from "vue";
import { DefaultApolloClient } from "@vue/apollo-composable";
// Router setup
const routes = [
  { path: "/", name: "Task List", component: TaskList },
  { path: "/create", name: "Create Task", component: CreateTask },
  { path: "/task-list/:id", name: "Task Detail", component: TaskDetails },
  { path: "/update-task/:id", name: "Update Tasks", component: UpdateTasks },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },

  render: () => h(App),
});

app.use(router);
app.mount("#app");
