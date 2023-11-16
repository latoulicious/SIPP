<script setup lang="ts">
import TableModul from "@/components/core/TableModul.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Modul Ajar</va-card-title>
      <va-card-content>
        <TableModul />
      </va-card-content>
    </va-card>
  </div>
</template>

<style>
.container {
  margin-top: 10px;
  margin-left: 60px;
  margin-right: 60px;
}
</style>
