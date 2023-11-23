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

<style lang="scss" scoped>
.VaLayout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

#content {
  flex: 1;
}
</style>
