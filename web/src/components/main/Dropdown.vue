<template>
  <div class="dropdown-wrapper">
    <va-button-dropdown
      class="absolute"
      label="User"
      icon="face"
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
        <router-link to="/login">Logout</router-link>
      </div>
    </va-button-dropdown>
  </div>
</template>

<style scoped>
.menu-item {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>
