<template>
  <div class="dropdown-wrapper">
    <va-button-dropdown
      class="absolute"
      :label="isLoggedIn ? username : 'User'"
      icon="person"
      icon-color="#ffffff"
      right-icon
      size="medium"
      preset="plain"
      color="#ffffff"
    >
      <div class="menu-item" v-if="isLoggedIn">
        <router-link :to="getSettingsLink()">Settings</router-link>
      </div>
      <div class="menu-item" v-if="isLoggedIn">
        <router-link :to="getUpdatePasswordLink()">Change Password</router-link>
      </div>
      <div class="menu-item">
        <router-link :to="getLogoutLink()">{{
          isLoggedIn ? "Logout" : "Login"
        }}</router-link>
      </div>
    </va-button-dropdown>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: false,
      username: "",
    };
  },
  mounted() {
    // Fetch user information from the JWT payload
    const payload = this.decodeJWT();
    if (payload) {
      this.isLoggedIn = true;
      this.username = payload.name;
    }
  },
  methods: {
    decodeJWT() {
      // Decode the JWT to access the payload
      const token = localStorage.getItem("jwtToken"); // Replace with your actual JWT key
      if (token) {
        const [, payloadBase64] = token.split(".");
        const payload = JSON.parse(atob(payloadBase64));
        return payload;
      }
      return null;
    },
    getSettingsLink() {
      return this.isLoggedIn ? "/settings" : "#"; // Replace with the correct route path or name
    },
    getUpdatePasswordLink() {
      return this.isLoggedIn ? "/change-password" : "#"; // Replace with the correct route path or name
    },
    getLogoutLink() {
      return this.isLoggedIn ? "/login" : "/login"; // Replace with the correct route path or name
    },
  },
};
</script>

<style scoped>
.menu-item {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>
