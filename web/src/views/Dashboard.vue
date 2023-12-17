// Dashboard.vue

<script setup lang="ts">
import Carousel from "@/components/main/Carousel.vue";
import InfoCard from "@/components/main/InfoCard.vue";
</script>

<template>
  <CorePageLayout>
    <main class="dashboard-container">
      <va-card class="dashboard-card">
        <Carousel />
        <va-divider />
        <InfoCard />
      </va-card>
    </main>
  </CorePageLayout>
</template>

<style scoped>
.dashboard-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin: 2vh 5vw; /* Adjust top and bottom margins based on your preference */
}

.dashboard-card {
  width: calc(95vw - 2 * 5vw); /* Adjust width based on your preference */
  max-width: 100vw; /* Set a maximum width if needed */
  padding: 20px; /* Adjust padding based on your preference */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); /* Optional: Add a box shadow for visual effect */
}
</style>
