
//三大界面 
import AiPage from "@/views/AiPage.vue";
import HomePage from "@/views/HomePage.vue";
import UserPage from "@/views/UserPage.vue";

// 登錄注冊
import CreateProfilePage from "@/views/Login/CreateProfilePage.vue";
import ForgetPasswordPage from "@/views/Login/ForgetPasswordPage.vue";
import LoginPage from "@/views/Login/LoginPage.vue";
import RegisterPage from "@/views/Login/RegisterPage.vue";
import ResetPasswordPage from "@/views/Login/ResetPasswordPage.vue";
import VerifyPage from "@/views/Login/VerifyPage.vue";

// 首頁小頁面
import NotifyPage from "@/views/NotifyPage.vue";
import LikePage from "@/views/LikePage.vue";
import WelFarePage from "@/views/WelFarePage.vue";

// 個人資料頁面
import EditPersonalDataPage from "@/views/User/EditPersonalDataPage.vue";
import LinkAccountPage from "@/views/User/LinkAccountPage.vue";
import PersonalDataPage from "@/views/User/PersonalDataPage.vue";
import QAPage from "@/views/User/QAPage.vue";
import SettingsPage from "@/views/User/SettingsPage.vue";
import FamilyPage from "@/views/User/FamilyPage.vue";
import FamilyInf from "@/views/familyInf.vue";
import FamilySettingPage from "@/views/FamilySettingPage.vue";


// 老人界面
import ElderlySearchPage from "@/views/Older/ElderlySearchPage.vue";
import RegionSelectionPage from "@/views/Older/RegionSelectionPage.vue";
import ServiceSelectionPage from "@/views/Older/ServiceSelectionPage.vue";
import SearchResultPage from "@/views/Older/SearchResultPage.vue";
import OlderHomePage from "@/views/Older/OlderHomePage.vue";

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
    path: "/welfare/:id",
    name: "welfare",
    component: WelFarePage,
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
    path: "/user/family/:id",
    name: "FamilyInfPage",
    component: FamilyInf,
    meta: { needLogin: true }
  },
  {
    path: "/user/family",
    name: "FamilyPage",
    component: FamilyPage,
    meta: { needLogin: true },
  },
  {
    path: "/user/family/setting/:id",
    name: "FmailySettingPage",
    component: FamilySettingPage,
    meta: { needLogin: true }
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
    path: "/older-home",
    name: "OlderHomePage",
    component: OlderHomePage,
    meta: { needLogin: false },
  }, //老人首頁暫時放這
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
  {
    path: "/elderlysearch",
    name: "ElderlySearchPage",
    component: ElderlySearchPage,
    meta: { needLogin: false },
  },
  {
    path: "/regionselection",
    name: "RegionSelectionPage",
    component: RegionSelectionPage,
    meta: { needLogin: false },
  },
  {
    path: "/serviceselection",
    name: "ServiceSelectionPage",
    component: ServiceSelectionPage,
    meta: { needLogin: false },
  },
  {
    path: "/searchresult",
    name: "SearchResultPage",
    component: SearchResultPage,
    meta: { needLogin: false },
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
