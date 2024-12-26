import request from "@/axios";
import User from "@/model/user";
import dayjs from "dayjs";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
    state: () => ({ user: new User(), avatarBase64: "" }),
    actions: {
        async fetchUser() {
            try {
                const response = await request.get(`users/avatar`);
                this.avatarBase64 = response.data.avatar_base64;

                const result = await request.get("/users")
                
                this.user = result.data

            } catch (err: any) {
                localStorage.removeItem("token")
            }

        },
        async saveUserHanlder(user: string) {
            try {
                const result = await request.put("/users", user);
                await this.fetchUser()
            } catch (err: any) {

            }

        }
    },
})