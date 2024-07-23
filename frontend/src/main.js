import './style.css';
import App from './App.vue';
import { apolloClient } from './apollo';
import { createRouter, createWebHistory } from 'vue-router';
import TaskList from './components/TaskList.vue';
import CreateTask from './components/CreateTask.vue';
import { createApp, provide, h } from 'vue'
import { DefaultApolloClient } from '@vue/apollo-composable'
// Router setup
const routes = [
  { path: '/', component: TaskList },
  { path: '/create', component: CreateTask },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp({
  setup () {
    provide(DefaultApolloClient, apolloClient)
  },

  render: () => h(App),
})

app.use(router);
app.mount('#app');