<script setup lang="ts">
import TablePRP from "@/components/core/TablePRP.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Penilaian Ranah Pengetahuan</va-card-title>
      <va-card-content>
        <TablePRP />
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
