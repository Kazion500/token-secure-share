import { axiosInstance } from "@/config/axios";
import { defineStore } from "pinia";

export interface Link {
  id: string;
  link: string;
  content: string;
  createdAt: string;
}

export const useLinkStore = defineStore("link", {
  state: () => ({
    link: null as Link | null,
    loading: false,
    error: null as Error | null,
  }),

  // getters
  //   getters: {
  //     link: (state) => state.link,
  //     loading: (state) => state.loading,
  //     error: (state) => state.error,
  //   },

  // optional actions
  actions: {
    async createLink(link: string) {
      this.loading = true;
      this.error = null;
      return new Promise(async (resolve, reject) => {
        try {
          const response = await axiosInstance.post("/links", { data: link });
          this.link = response.data.data;
          resolve(response.data);
        } catch (error: any) {
          this.error = error;
          reject(error);
        } finally {
          this.loading = false;
        }
      });
    },

    async getLink(id: string) {
      this.loading = true;
      this.error = null;
      return new Promise(async (resolve, reject) => {
        try {
          const response = await axiosInstance.get(`/links/${id}`);
          this.link = response.data.data;
          resolve(response.data);
        } catch (error: any) {
          this.error = error;
          reject(error);
        } finally {
          this.loading = false;
        }
      });
    },
  },
});
