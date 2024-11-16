import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./style/index.css";
import { createPinia } from 'pinia'
import validate from "validate.js";

validate.options = {format: "flat",fullMessages: false};
validate.validators.presence.options = {message: "不能為空"};
const app = createApp(App)
const pinia = createPinia();
app.use(router)
app.use(pinia)
app.mount('#app')