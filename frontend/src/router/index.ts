import AiPage from "@/views/AiPage.vue";
import HomePage from "@/views/HomePage.vue";
import LikePage from "@/views/LikePage.vue";
import CreateProfilePage from "@/views/Login/CreateProfilePage.vue";
import ForgetPasswordPage from "@/views/Login/ForgetPasswordPage.vue";
import LoginPage from "@/views/Login/LoginPage.vue";
import RegisterPage from "@/views/Login/RegisterPage.vue";
import ResetPasswordPage from "@/views/Login/ResetPasswordPage.vue";
import VerifyPage from "@/views/Login/VerifyPage.vue";
import EditPersonalDataPage from "@/views/User/EditPersonalDataPage.vue";
import FamilyPage from "@/views/User/FamilyPage.vue";
import LinkAccountPage from "@/views/User/LinkAccountPage.vue";
import PersonalDataPage from "@/views/User/PersonalDataPage.vue";
import QAPage from "@/views/User/QAPage.vue";
import SettingsPage from "@/views/User/SettingsPage.vue";
import UserPage from "@/views/UserPage.vue";
import NotifyPage from "@/views/NotifyPage.vue";
import {
  createRouter,
  createWebHistory,
  RouterView,
  type RouteRecordRaw,
} from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/home", // 重定向到登入頁
    meta: { needLogin: false },
  },
  {
    path: "/account",
    children: [
      {
        path: "login",
        name: "LoginPage",
        component: LoginPage,
        meta: { needLogin: false },
      },
      {
        path: "forgot-password",
        name: "ForgetPasswordPage",
        component: ForgetPasswordPage,
        meta: { needLogin: false },
      },
      {
        path: "reset-password",
        name: "ResetPasswordPage",
        component: ResetPasswordPage,
        meta: { needLogin: false },
      },
      {
        path: "register",
        name: "RegisterPage",
        component: RegisterPage,
        meta: { needLogin: false },
      },
      {
        path: "verify",
        name: "VerifyPage",
        component: VerifyPage,
        meta: { needLogin: false },
      },
      {
        path: "create-profile",
        name: "CreateProfilePage",
        component: CreateProfilePage,
        meta: { needLogin: false },
      },
    ],
    component: RouterView,
  },
  {
    path: "/user/settings",
    name: "Settings",
    component: SettingsPage,
    meta: { needLogin: false },
  },
  {
    path: "/user/edit-personal-data",
    name: "EditPersonalDataPage",
    component: EditPersonalDataPage,
    meta: { needLogin: true },
  },
  {
    path: "/user/qa",
    name: "QAPage",
    component: QAPage,
    meta: { needLogin: false },
  },
  {
    path: "/user/personal-data",
    name: "PersonalDataPage",
    component: PersonalDataPage,
    meta: { needLogin: true },
  },
  {
    path: "/user/family",
    name: "FamilyPage",
    component: FamilyPage,
    meta: { needLogin: true },
  },
  {
    path: "/user/link-account",
    name: "LinkAccountPage",
    component: LinkAccountPage,
    meta: { needLogin: false }, //記得改回true
  },
  // 最后定义父路径
  {
    path: "/user",
    name: "UserPage",
    component: UserPage,
  },
  {
    path: "/home",
    name: "HomePage",
    component: HomePage,
    meta: { needLogin: false },
  },
  {
    path: "/ai",
    name: "AiPage",
    component: AiPage,
    meta: { needLogin: false },
  },
  {
    path: "/favorites",
    name: "LikePage",
    component: LikePage,
    meta: { needLogin: true },
  },
  {
    path: "/notifications",
    name: "NotifyPage",
    component: NotifyPage,
    meta: { needLogin: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if (to.meta.needLogin) {
    if (localStorage.getItem("token")) {
      return true;
    } else {
      router.push("/account/login");
      return false;
    }
  } else {
    return true;
  }
});

export default router;
