<script setup lang="ts">
import { useAppStore } from "@/stores/index";
import Dropdown from "@/components/main/Dropdown.vue";

const appStore = useAppStore();

const toggleSidebar = () => {
  appStore.toggleSidebar();
};
</script>

<template>
  <div class="navbar-wrapper">
    <VaNavbar class="a" color="#090F33">
      <template #left>
        <VaButton
          color="#ffffff"
          preset="plain"
          :icon="appStore.isSidebarVisible ? 'menu_open' : 'menu'"
          @click="toggleSidebar"
        />
      </template>
      <template #center>
        <va-navbar-item>
          <p class="logo-container">
            <img src="@/assets/logo.png" alt="Logo" class="logo" />
          </p>
        </va-navbar-item>
      </template>
      <template #right>
        <Dropdown />
      </template>
    </VaNavbar>
  </div>
</template>

<style lang="scss" scoped>
.a {
  padding-top: 0.3rem;
  padding-bottom: 0.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  height: 50px;
  max-height: 100px;
  width: auto;
  margin: 0; /* Remove any default margin */
}

.logo-container {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Set a minimum height for VaNavbar */
.va-navbar {
  min-height: 70px; /* Adjust this value as needed */
}

.va-navbar-item {
  display: flex;
  align-items: center;
}

/* If you want to center the Dropdown component in the #right section */
#right {
  display: flex;
  align-items: center;
}
</style>
