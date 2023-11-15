<template>
  <div v-if="selectedTab === 'login'">
    <form @submit.prevent="login">
      <div class="form-group">
        <va-input
          v-model="value"
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
          v-model="value"
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

      <va-button color="info" gradient type="submit" href="/dashboard"
        >Login</va-button
      >
    </form>
  </div>
</template>

<script>
export default {
  props: {
    selectedTab: String,
  },
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    login() {
      console.log("Logging in...", this.username, this.password);
      // Perform login logic here
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
