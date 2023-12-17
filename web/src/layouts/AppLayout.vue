<script setup lang="ts">
import Navbar from "@/components/bar/Navbar.vue";
import Sidebar from "@/components/bar/Sidebar.vue";
import Footer from "@/components/bar/Footer.vue";
</script>

<template>
  <div class="app-layout">
    <VaLayout :left="{ absolute: true }">
      <template #top>
        <Navbar />
      </template>

      <template #left>
        <Sidebar />
      </template>

      <template #content>
        <main class="scrollable-content">
          <router-view />
        </main>
      </template>
    </VaLayout>

    <Footer />
  </div>
</template>

<style lang="scss">
.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

main {
  flex: 1;
  display: flex;
  flex-direction: column;
  margin: 20px 60px 0; /* top margin is 20px, left and right margins are 40px, bottom margin is 0 */
  margin-bottom: 60px; /* Adjust this value to the height of your footer */
}

.Footer {
  margin-top: auto; /* Make the footer stay at the bottom of the viewport */
}
.custom-sidebar {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* Make the sidebar extend to the bottom of the viewport */
  overflow-y: auto; /* Add this to enable scrolling if the content exceeds the sidebar height */
  flex: 0 0 200px; /* Fixed width for the sidebar */
}

.custom-sidebar .VaSidebarItems {
  flex-grow: 1;
}

.custom-sidebar .VaScrollbar {
  display: none;
}
</style>
