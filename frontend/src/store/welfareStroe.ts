import request from "@/axios";
import Welfare from "@/model/welfare";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { defineStore } from "pinia";

export const useWelfareStore = defineStore("welfare", {
    state: () => ({ welfares: new Array<Welfare> }),
    actions: {
        getWelfare(regions: Array<number>, services: Array<number>): Array<Welfare> {
            // 若地區或服務沒有選擇的話，分別將對應條件設為 true，使其不影響篩選
            const regionsText: Array<string> = regions.map(getTextByLocation);
            return this.welfares.filter((item) => {
                const matchesRegion = regions.length === 0 || regionsText.some(region => item.city.includes(region));
                const matchesService = services.length === 0 || services.some(service => item.category.includes(service));
                return matchesRegion && matchesService;
            });
        },

        async fetchWelfare() {
            const result = await request.get("/welfare")
            this.welfares = result.data
        }
    }
})