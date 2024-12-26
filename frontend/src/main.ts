import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";
import "v-calendar/style.css";
import "./style/index.css";
import { useFavoriteStore } from "./store/favoriteStore";
import { useUserStore } from "./store/userStore";
import { useFamilyStore } from "./store/familyStore";
import { isAxiosError } from "axios";

// 创建 Vue 实例
const app = createApp(App);

// 创建 Pinia 实例并注入
const pinia = createPinia();
app.use(pinia);

// 注册路由
app.use(router);
const userStore = useUserStore();
const favoriteStore = useFavoriteStore();
const familyStore = useFamilyStore();

(async () => {
    try {
        await userStore.fetchUser();
        await favoriteStore.fetchFavorite();
        await familyStore.fetchFamily();
    } catch (e) {
        if(isAxiosError(e)) {
            console.log(e.response?.data);
        }
    } finally {
        app.mount("#app");
    }
})();
