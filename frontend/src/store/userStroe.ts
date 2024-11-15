import request from "@/axios";
import User from "@/model/user";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
    state: () => ({ user: new User() }),
    actions: {
        async fetchUser() {
            try {
                const result = await request.get("/users")
                this.user = result.data
            } catch (err: any) {
                localStorage.removeItem("token")
            }

        },
        async saveUserHanlder(user: string) {
            try {
                const result = await request.put("/users", user);
                this.user = result.data
            } catch (err: any) {

            }

        },
    },
})