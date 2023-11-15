import { defineStore } from "pinia";

export const useAppStore = defineStore("app", {
  state: () => ({
    isSidebarVisible: true, // or any default value
  }),
  actions: {
    toggleSidebar() {
      this.isSidebarVisible = !this.isSidebarVisible;
    },
  },
});
