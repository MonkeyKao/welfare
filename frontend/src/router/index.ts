import CreateProfilePage from "@/views/CreateProfilePage.vue";
import ForgetPasswordPage from "@/views/ForgetPasswordPage.vue";
import HomePage from "@/views/HomePage.vue";
import LoginPage from "@/views/LoginPage.vue";
import RegisterPage from "@/views/RegisterPage.vue";
import ResetPasswordPage from "@/views/ResetPasswordPage.vue";
import AiPage from "@/views/AiPage.vue"
import VerifyPage from "@/views/VerifyPage.vue";
import Settings from "@/views/SettingsPage.vue"
import LikePage from "@/views/LikePage.vue";
import NotifyPage from "@/views/NotifyPage.vue";

import { createRouter, createWebHistory } from "vue-router";
const routes = [
  {
    path: "/",
    name: "Home", // 可以用來顯示首頁或重定向
    redirect: "/login", // 重定向到登入頁
  },
  {
    path: "/login",
    name: "LoginPage",
    component: LoginPage,
  },
  {
    path: "/forgot-password",
    name: "ForgetPasswordPage",
    component: ForgetPasswordPage,
  },
  {
    path: "/reset-password",
    name: "ResetPasswordPage",
    component: ResetPasswordPage,
  },
  {
    path: "/register",
    name: "RegisterPage",
    component: RegisterPage,
  },
  {
    path: "/verify",
    name: "VerifyPage",
    component: VerifyPage,
  },
  {
    path: "/create-profile",
    name: "CreateProfilePage",
    component: CreateProfilePage,
  },
  {
    path: "/home",
    name: "HomePage",
    component: HomePage,
  },
  {
    path: "/ai",
    name: "AiPage",
    component: AiPage,
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
  },
  {
    path: "/favorites",
    name: "LikePage",
    component: LikePage,
  },
  {
    path: "/notifications",
    name: "NotifyPage",
    component: NotifyPage,
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
