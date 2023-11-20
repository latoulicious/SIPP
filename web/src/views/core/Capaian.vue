<script setup lang="ts">
import TableCapaian from "@/components/core/TableCapaian.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Capaian Pembelajaran</va-card-title>
      <va-card-content>
        <TableCapaian />
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
