import request from "@/axios";
import type Welfare from "@/model/welfare";
import { defineStore } from "pinia";

export const useFavoriteStore = defineStore("favorite", {
    state: () => ({ favorites: new Array<Welfare> }),
    actions: {
        async fetchFavorite() {
            try {
                const result = await request.get("/favorite")
                this.favorites = result.data
            } catch (err: any) {

            }

        },
        async createFavoriteHanlder(welfare:Welfare) {
            try {
                await request.post("/favorite/" + welfare.id)
                this.favorites.push(welfare)
            } catch (err: any) {
                throw(err.response.data.error)
            }

        },
        async deleteFavoriteHandler(id:number) {
            try {
                const result = await request.delete("/favorite/" + id)
                this.favorites = this.favorites.filter((favorite: Welfare) => favorite.id !== id);
                // 從factorite刪除
            } catch (err: any) {

            }
        }
    },
})