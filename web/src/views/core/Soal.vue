<script setup lang="ts">
import TableSoal from "@/components/core/TableSoal.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Soal</va-card-title>
      <va-card-content>
        <TableSoal />
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
