import request from "@/axios";
import Family from "@/model/family";
import { defineStore } from "pinia";

export const useFamilyStore = defineStore('family', {
    state: () => ({
        families: new Array<Family>
    }),
    actions: {
        async fetchFamily() {
            const result = await request.get("/family")
            this.families = JSON.parse(result.data);
        },
        async joinFamily(code:string) {
            try {
                await request.post("/family/join/" + code)
            } catch (err: any) {
                throw (err.response.data)
            }
        },
        async createFamily(name: string) {
            try {
                const result = await request.post("/family/" + name)
                this.families.push(new Family(result.data.FamilyName, result.data.ID, []))
            } catch (err: any) {
                throw (err.response.data)
            }

        },
        async deleteFamily(id: number) {
            try {
                await request.delete("/family/" + id)
                this.families = this.families.filter((family) => family.familyId !== id)
            } catch (err: any) {
                throw (err.response.data)
            }
        },
        async leaveFamily(id:number) {
            try {
                await request.delete("/family/leave/"+id)
                this.families = this.families.filter((family) => family.familyId !== id)
            } catch(err:any) {
                throw(err.response.data)
            }
        }
        
    }
})