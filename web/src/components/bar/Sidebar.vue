<script setup lang="ts">
import { useAppStore } from "@/stores/index";
import { useUserRoleStore } from "@/stores/userRoleStore";
import { onMounted, computed } from "vue";

const appStore = useAppStore();
const userRoleStore = useUserRoleStore();

onMounted(async () => {
  // Fetch user role when the component is mounted
  await userRoleStore.fetchUserRole();
});

const showManageUsers = computed(
  () => !userRoleStore.loading && userRoleStore.role === "Admin",
);
</script>

<template>
  <VaSidebar
    v-model="appStore.isSidebarVisible"
    color="#AFC1EE"
    class="custom-sidebar"
  >
    <router-link to="/dashboard">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="dashboard" />
          <VaSidebarItemTitle>Dashboard</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/capaian">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Capaian Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/alurtujuan">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Alur Tujuan Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/modulajar">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Modul Ajar</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/penilaian">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assignment" />
          <VaSidebarItemTitle>Penilaian Ranah Pengetahuan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/asesmen">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assignment" />
          <VaSidebarItemTitle>Asesmen</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/soal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assessment" />
          <VaSidebarItemTitle>Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/banksoal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assessment" />
          <VaSidebarItemTitle>Bank Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link v-show="showManageUsers" to="/user">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="people" />
          <VaSidebarItemTitle>Manage Users</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
  </VaSidebar>
</template>

<style scoped>
.custom-sidebar {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow-y: auto; /* Add this to enable scrolling if the content exceeds the sidebar height */
}

.custom-sidebar .VaSidebarItems {
  flex-grow: 1;
}

.custom-sidebar .VaScrollbar {
  display: none;
}
</style>
