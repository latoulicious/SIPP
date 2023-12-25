<script>
import axios from "axios";

function decodeJwt(jwtToken) {
  try {
    const [, payloadBase64] = jwtToken.split(".");
    const payload = JSON.parse(atob(payloadBase64));
    return payload;
  } catch (error) {
    console.error("Error decoding JWT:", error);
    return null;
  }
}

function getUserId() {
  const jwtToken = localStorage.getItem("jwtToken");
  if (!jwtToken) {
    console.error("JWT token not available");
    return null;
  }

  const decodedToken = decodeJwt(jwtToken);
  return decodedToken && (decodedToken.user_id || decodedToken.ID);
}

export default {
  data() {
    return {
      currentPassword: "",
      newPassword: "",
      confirmNewPassword: "",
      currentUser: null,
      message: "Required field",
    };
  },
  methods: {
    async changePassword() {
      try {
        const userId = getUserId();
        if (!userId) {
          console.error("User ID not available");
          return;
        }

        const response = await axios.post(`http://localhost:3000/api/user/${userId}/change-password`, {
          currentPassword: this.currentPassword,
          newPassword: this.newPassword,
          confirmNewPassword: this.confirmNewPassword,
        });

        console.log(response.data);

        if (response.data.status === "success") {
          this.fetchData();
          this.resetForm();
        } else {
          console.error("Password change failed:", response.data.message);
          // Handle the failure, show a message to the user, etc.
        }
      } catch (error) {
        console.error("Error changing password:", error);
        // Handle the error, show a message to the user, etc.
      }
    },

    async fetchData() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");
        if (!jwtToken) {
          console.error("JWT token not available");
          return;
        }

        const userId = getUserId();
        if (!userId) {
          console.error("User ID not available");
          return;
        }

        const response = await this.$axios.get(`http://localhost:3000/api/user/${userId}`, {
          headers: {
            Authorization: `Bearer ${jwtToken}`,
          },
        });

        this.currentUser = response.data.data;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },

    resetForm() {
      this.currentPassword = "";
      this.newPassword = "";
      this.confirmNewPassword = "";
    },
  },
  created() {
    console.log("Fetching user data...");
    this.fetchData();
  },
};
</script>


<template>
  <div class="pw-container">
    <VaInput
      v-model="currentPassword"
      class="mb-6"
      :messages="message"
      label="Current Password"
      placeholder="Single-line message"
    />
    <VaInput
      v-model="newPassword"
      class="mb-6"
      :messages="message"
      label="New Password"
      placeholder="Single-line message"
    />
    <VaInput
      v-model="confirmNewPassword"
      class="mb-6"
      :messages="message"
      label="Confirm New Password"
      placeholder="Single-line message"
    />
    <va-button round color="info" gradient @click="changePassword">
      Change Password
    </va-button>
  </div>
</template>

<style scoped>
.pw-container {
  display: grid;
  flex-direction: column;
  align-items: center;
}

.va-input {
  display: block;
  margin-bottom: 10px;
}
</style>
