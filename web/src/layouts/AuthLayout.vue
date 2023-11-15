<template>
  <div class="auth-layout">
    <va-card class="auth-card" stripe="primary">
      <div class="content">
        <div class="logo-container">
          <img src="@/assets/login.png" alt="Logo" class="logo" />
        </div>
        <div class="form-container">
          <div class="tabs"></div>
          <Login :selectedTab="selectedTab" />
        </div>
      </div>
    </va-card>
  </div>
</template>

<script>
import Login from "@/views/auth/Login.vue";

export default {
  data() {
    return {
      selectedTab: "login",
    };
  },
  methods: {
    selectTab(tab) {
      this.selectedTab = tab;
    },
  },
  components: {
    Login, // Register the Login component
  },
};
</script>

<style scoped>
.auth-layout {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #f0f0f0;
}

.logo-container {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.logo {
  transform: scale(0.85); /* Adjust the scale factor as needed */
}

.auth-card {
  width: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.content {
  display: flex;
}

.form-container {
  flex: 1;
  padding-left: 30px; /* Add padding between logo and form fields */
}

.tabs {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
  font-size: 25px;
}

.tabs button {
  padding: 10px;
  border: none;
  background-color: transparent;
  cursor: pointer;
}

.tabs button.active {
  font-weight: bold;
}

.content {
  padding: 10px;
}
</style>
