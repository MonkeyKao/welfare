import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./style/index.css";
import { createPinia } from 'pinia'
import validate from "validate.js";

import 'v-calendar/style.css';
import { VueQrcodeReader } from "vue-qrcode-reader";


validate.options = {format: "flat",fullMessages: false};
validate.validators.presence.options = {message: "不能為空"};
const app = createApp(App)
const pinia = createPinia();
app.use(VueQrcodeReader)
app.use(router)
app.use(pinia)



app.mount('#app')