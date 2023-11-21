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
      <div class="menu-item">
        <router-link to="/settings">Settings</router-link>
      </div>
      <div class="menu-item">
        <router-link v-if="isLoggedIn" to="/logout">Logout</router-link>
        <router-link v-else to="/login">Login</router-link>
      </div>
    </va-button-dropdown>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: false, // Set this to true if the user is logged in
      username: "", // Set the username when the user is logged in
    };
  },
};
</script>

<style scoped>
.menu-item {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>
