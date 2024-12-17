
//三大界面 
import MouInput from "@/components/MouInput.vue";
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
import WelFareInfPage from "@/views/Home/WelFareInfPage.vue";
import NotifyPage from "@/views/Home/NotifyPage.vue";
import LikePage from "@/views/Home/LikePage.vue";

// 個人資料頁面
import EditPersonalDataPage from "@/views/User/EditPersonalDataPage.vue";
import LinkAccountPage from "@/views/User/LinkAccountPage.vue";
import PersonalDataPage from "@/views/User/PersonalDataPage.vue";
import QAPage from "@/views/User/QAPage.vue";
import SettingsPage from "@/views/User/SettingsPage.vue";

// 家庭頁面
import FamilyInfPage from "@/views/Family/FamilyInfPage.vue";
import FamilyHomePage from "@/views/Family/FamilyHomePage.vue";
import FamilySettingPage from "@/views/Family/FamilySettingPage.vue";

// 老人界面
import OlderHomePage from "@/views/Older/OlderHomePage.vue";
import OlderSearchHomePage from "@/views/Older/OlderSearchHomePage.vue";
import OlderRegionSelectionPage from "@/views/Older/OlderRegionSelectionPage.vue";
import OlderServiceSelectionPage from "@/views/Older/OlderServiceSelectionPage.vue";

import {
  createRouter,
  createWebHistory,
  RouterView,
  type RouteRecordRaw,
} from "vue-router";

// 老人路由
const OlderRoutes: Array<RouteRecordRaw> = [
  {
    path: "/accessibility/home",
    name: "OlderHomePage",
    component: OlderHomePage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/accessibility/search",
    name: "OlderSearchHomePage",
    component: OlderSearchHomePage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/accessibility/region",
    name: "OlderRegionSelectionPage",
    component: OlderRegionSelectionPage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/accessibility/service",
    name: "OlderServiceSelectionPage",
    component: OlderServiceSelectionPage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/accessibility/result",
    name: "OlderSearchHomePage",
    component: OlderSearchHomePage,
    meta: { needLogin: false, showBottomBar: false },
  },
]

const UserRoutes: Array<RouteRecordRaw> = [
  {
    path: "/user/settings",
    name: "Settings",
    component: SettingsPage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/user/edit-personal-data",
    name: "EditPersonalDataPage",
    component: EditPersonalDataPage,
    meta: { needLogin: true, showBottomBar: false },
  },
  {
    path: "/user/qa",
    name: "QAPage",
    component: QAPage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/user/personal-data",
    name: "PersonalDataPage",
    component: PersonalDataPage,
    meta: { needLogin: true, showBottomBar: false },
  },
  {
    path: "/user/family/:id",
    name: "FamilyInfPage",
    component: FamilyInfPage,
    meta: { needLogin: true, showBottomBar: false }
  },
  {
    path: "/user/family",
    name: "FamilyHomePage",
    component: FamilyHomePage,
    meta: { needLogin: true, showBottomBar: false },
  },
  {
    path: "/user/family/setting/:id",
    name: "FmailySettingPage",
    component: FamilySettingPage,
    meta: { needLogin: true, showBottomBar: false }
  },
  {
    path: "/user/link-account",
    name: "LinkAccountPage",
    component: LinkAccountPage,
    meta: { needLogin: false, showBottomBar: false }, //記得改回true
  },
  // 最后定义父路径
  {
    path: "/user",
    name: "UserPage",
    component: UserPage,
  },
]

const HomeRoutes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/home",
    meta: { needLogin: false },
  },
  {
    path: "/welfare/:id",
    name: "welfare",
    component: WelFareInfPage,
    meta: { needLogin: false, showBottomBar: false },
  },
  {
    path: "/home",
    name: "HomePage",
    component: HomePage,
    meta: { needLogin: false },
  },
  {
    path: "/mou",
    name: "MouInput",
    component: MouInput,
    meta: { needLogin: false },
  },
  {
    path: "/favorites",
    name: "LikePage",
    component: LikePage,
    meta: { needLogin: true, showBottomBar: false },
  },
  {
    path: "/notifications",
    name: "NotifyPage",
    component: NotifyPage,
    meta: { needLogin: true, showBottomBar: false },
  },
]

const AccountRoutes: Array<RouteRecordRaw> = [
  {
    path: "/account",
    children: [
      {
        path: "login",
        name: "LoginPage",
        component: LoginPage,
        meta: { needLogin: false, showBottomBar: false },
      },
      {
        path: "forgot-password",
        name: "ForgetPasswordPage",
        component: ForgetPasswordPage,
        meta: { needLogin: false, showBottomBar: false },
      },
      {
        path: "reset-password",
        name: "ResetPasswordPage",
        component: ResetPasswordPage,
        meta: { needLogin: false, showBottomBar: false },
      },
      {
        path: "register",
        name: "RegisterPage",
        component: RegisterPage,
        meta: { needLogin: false, showBottomBar: false },
      },

      {
        path: "verify",
        name: "VerifyPage",
        component: VerifyPage,
        meta: { needLogin: false, showBottomBar: false },
      },
      {
        path: "create-profile",
        name: "CreateProfilePage",
        component: CreateProfilePage,
        meta: { needLogin: false, showBottomBar: false },
      },
    ],
    component: RouterView,
  }
];

const routes: Array<RouteRecordRaw> = [
  ...OlderRoutes,
  ...UserRoutes,
  ...AccountRoutes,
  ...HomeRoutes
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
