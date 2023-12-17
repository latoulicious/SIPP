import { defineStore } from "pinia";
import axios from "axios";
import { nextTick } from "vue";

export const useUserRoleStore = defineStore("userRole", {
  state: () => ({
    role: null as string | null,
    loading: true,
  }),
  actions: {
    async fetchUserRole() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          console.error("JWT token is missing.");
          return;
        }

        const response = await axios.get(
          "http://localhost:3000/api/get-user-role",
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
              "Content-Type": "application/json",
            },
          },
        );

        // Log the entire response for debugging
        console.log("Response from /api/get-user-role:", response);

        // Log the JSON data received from the response
        const jsonData = response.data;
        console.log("JSON Data from response:", jsonData);

        this.role = jsonData.role;

        // Set loading to false after the role has been updated
        this.loading = false;

        // Use nextTick to ensure that the component re-renders after the role is updated
        await nextTick();
      } catch (error) {
        console.error("Error fetching user role:", error);
      }
    },
  },
});
