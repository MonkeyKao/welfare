import { createApp } from "vue";
import App from "./App.vue";
import "./assets/main.css";
import router from "./router";
import "./style/index.css";
import { createPinia } from 'pinia'

const app = createApp(App)
const pinia = createPinia();
app.use(router)
app.use(pinia)
app.mount('#app')