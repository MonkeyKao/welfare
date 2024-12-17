import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";
import "v-calendar/style.css";
import "./style/index.css";

// 创建 Vue 实例
const app = createApp(App);

// 创建 Pinia 实例并注入
const pinia = createPinia();
app.use(pinia);

// 注册路由
app.use(router);

// 挂载应用到 DOM
app.mount("#app");