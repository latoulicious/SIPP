import { defineStore } from "pinia";

export const useFilterStore = defineStore("filter", {
  state: () => ({
    filterValue: "",
  }),
  actions: {
    setFilterValue(value) {
      this.filterValue = value;
    },
  },
});
