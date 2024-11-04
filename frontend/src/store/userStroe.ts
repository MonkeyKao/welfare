import request from "@/axios";
import User from "@/model/user";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
    state: () => ({user:new User()}),
    actions: {
        async fetchUser() {
            const result = await request.get("/users")
            this.user = result.data
        }
    }
})