<script setup lang="ts">
import { useAppStore } from "@/stores/index";
const appStore = useAppStore();
</script>

<template>
  <VaSidebar v-model="appStore.isSidebarVisible" color="#AFC1EE">
    <router-link to="/dashboard">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="home" />
          <VaSidebarItemTitle>Dashboard</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/capaian">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Capaian Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/alurtujuan">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Alur Tujuan Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/modulajar">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Modul Ajar</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/penilaian">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Penilaian Ranah Pengetahuan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/soal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/banksoal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder" />
          <VaSidebarItemTitle>Bank Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
  </VaSidebar>
</template>
