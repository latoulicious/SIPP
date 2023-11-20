<script setup lang="ts">
import Navbar from "@/components/bar/Navbar.vue";
import Sidebar from "@/components/bar/Sidebar.vue";
import Footer from "@/components/bar/Footer.vue";
</script>

<template>
  <VaLayout :left="{ absolute: true }">
    <template #top>
      <Navbar />
    </template>

    <template #left>
      <Sidebar />
    </template>

    <template #content>
      <main>
        <router-view />
      </main>
      <Footer />
    </template>
  </VaLayout>
</template>
