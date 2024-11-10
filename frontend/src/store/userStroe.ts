import request from "@/axios";
import User from "@/model/user";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
    state: () => ({ user: new User() }),
    actions: {
        async fetchUser() {
            const result = await request.get("/users")
            this.user = result.data
        },
        async saveUserHanlder(user: string) {
            try{
                const result = await request.post("/users/updateuser", user);
                this.user = result.data
            }catch{

            }
            
        },
    },
})