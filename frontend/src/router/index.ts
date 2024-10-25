import { createRouter, createWebHistory } from 'vue-router';
import SettingsPage from '@/views/SettingsPage.vue';
import HomePage from '@/views/HomePage.vue';
import AiPage from '@/views/AiPage.vue';

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: HomePage,
  },
  {
    path: '/ai',
    name: 'Ai',
    component: AiPage,
  },
  {
    path: '/settings',
    name: 'Settings',
    component: SettingsPage,
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

export default router;