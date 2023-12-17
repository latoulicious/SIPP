<script setup lang="ts">
import TableCapaian from "@/components/core/TableCapaian.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Capaian Pembelajaran</va-card-title>
      <va-card-content>
        <TableCapaian />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
