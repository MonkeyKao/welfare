import AiPage from "@/views/AiPage.vue";
import CreateProfilePage from "@/views/CreateProfilePage.vue";
import FamilyPage from "@/views/FamilyPage.vue";
import ForgetPasswordPage from "@/views/ForgetPasswordPage.vue";
import HomePage from "@/views/HomePage.vue";
import LinkAccountPage from "@/views/LinkAccountPage.vue";
import LoginPage from "@/views/LoginPage.vue";
import PersonalDataPage from "@/views/PersonalDataPage.vue";
import QAPage from "@/views/QAPage.vue";
import RegisterPage from "@/views/RegisterPage.vue";
import ResetPasswordPage from "@/views/ResetPasswordPage.vue";
import Settings from "@/views/SettingsPage.vue";
import UserPage from "@/views/UserPage.vue";
import VerifyPage from "@/views/VerifyPage.vue";
import LikePage from "@/views/LikePage.vue";
import NotifyPage from "@/views/NotifyPage.vue";

import { createRouter, createWebHistory } from "vue-router";
const routes = [
  {
    path: "/",
    name: "Home", // 可以用來顯示首頁或重定向
    redirect: "/home", // 重定向到登入頁
    meta: {"needLogin":true}
  },
  {
    path: "/login",
    name: "LoginPage",
    component: LoginPage,
    meta: {"needLogin":false}
  },
  {
    path: "/forgot-password",
    name: "ForgetPasswordPage",
    component: ForgetPasswordPage,
    meta: {"needLogin":false}
  },
  {
    path: "/reset-password",
    name: "ResetPasswordPage",
    component: ResetPasswordPage,
    meta: {"needLogin":false}
  },
  {
    path: "/register",
    name: "RegisterPage",
    component: RegisterPage,
    meta: {"needLogin":false}
  },
  {
    path: "/verify",
    name: "VerifyPage",
    component: VerifyPage,
    meta: {"needLogin":false}
  },
  {
    path: "/create-profile",
    name: "CreateProfilePage",
    component: CreateProfilePage,
    meta: {"needLogin":false}

  },
  {
    path: "/home",
    name: "HomePage",
    component: HomePage,
    meta: {"needLogin":false}
  },
  {
    path: "/ai",
    name: "AiPage",
    component: AiPage,
    meta: {"needLogin":false}
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
    meta: {"needLogin":false}
  },
  {
    path: "/qa",
    name: "QAPage",
    component: QAPage,
    meta: {"needLogin":false}
  },
  {
    path: "/user",
    name: "UserPage",
    component: UserPage,
    meta: {"needLogin":false}
  },
  {
    path: "/personal-data",
    name: "PersonalDataPage",
    component: PersonalDataPage,
    meta: {"needLogin":true}
  },
  {
    path: "/family",
    name: "FamilyPage",
    component: FamilyPage,
    meta: {"needLogin":true}
  },
  {
    path: "/link-account",
    name: "LinkAccountPage",
    component: LinkAccountPage,
    meta: {"needLogin":true}
  },
  {
    path: "/favorites",
    name: "LikePage",
    component: LikePage,
    meta: {"needLogin":true}
  },
  {
    path: "/notifications",
    name: "NotifyPage",
    component: NotifyPage,
    meta: {"needLogin":true}
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if(to.meta.needLogin) {
    if(localStorage.getItem("token")){
      return true;
    }else {
      router.push("/login")
      return false;
    }
  }else {
    return true
  }
  
})

export default router;
