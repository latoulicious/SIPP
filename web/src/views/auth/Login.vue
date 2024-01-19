<template>
  <div v-if="selectedTab === 'login'">
    <form @submit.prevent="login">
      <div class="form-group">
        <va-input
          v-model="username"
          class="mb-6"
          type="text"
          label="username"
          placeholder="Enter your username"
        >
          <template #prependInner>
            <va-icon name="account_box" color="secondary" />
          </template>
        </va-input>
      </div>

      <div class="form-group">
        <va-input
          v-model="password"
          class="mb-6"
          type="password"
          label="password"
          placeholder="Enter your password"
        >
          <template #prependInner>
            <va-icon name="mail_outline" color="secondary" />
          </template>
        </va-input>
      </div>

      <va-button color="info" gradient type="submit">Login</va-button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "Admin",
      password: "admin",
      selectedTab: "login", // Make sure selectedTab is defined
    };
  },
  methods: {
    async login() {
      try {
        const response = await this.$axios.post("/api/auth/login", {
          username: this.username,
          password: this.password,
        });

        // Store the JWT token in localStorage or a secure storage method
        const token = response.data.data; // Use response.data instead of response.data.data

        // Access localStorage directly without using `this`
        localStorage.setItem("jwtToken", token);

        // Redirect the user to the dashboard or another protected route
        this.$router.push({ name: "dashboard" });
      } catch (error) {
        console.error("Login failed:", error);
        // Handle login error (display error message, etc.)
      }
    },
  },
};
</script>

<style scoped>
.form-group {
  flex: auto;
  margin-bottom: 15px;
  margin-top: 30px;
}

.form-group input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}
</style>
