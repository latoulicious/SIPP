<script setup lang="ts">
import TableAlur from "@/components/core/TableAlur.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Alur Tujuan Pembelajaran</va-card-title>
      <va-card-content>
        <TableAlur />
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
