<script setup lang="ts">
import ChangePassword from "@/components/other/Password.vue";
</script>

<template>
  <CorePageLayout>
    <div class="container">
      <va-card stripe stripe-color="danger">
        <va-card-title>Create New Password</va-card-title>
        <va-card-content>
          <ChangePassword />
        </va-card-content>
      </va-card>
    </div>
  </CorePageLayout>
</template>

<style scoped>
.container {
  display: grid;
  margin: 10px auto; /* Set auto for left and right margins to center the container */
  width: 50%;
}
</style>
