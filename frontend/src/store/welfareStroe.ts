import request from "@/axios";
import Welfare from "@/model/welfare";
import { defineStore } from "pinia";

export const useWelfareStore = defineStore("welfare", {
    state: () => ({ welfares: new Array<Welfare> }),
    actions: {
        async fetchWelfare() {
            const result = await request.get("/welfare")
            this.welfares = result.data
        }
    }
})