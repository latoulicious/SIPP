<script setup lang="ts">
import TableSoal from "@/components/core/TableSoal.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Soal</va-card-title>
      <va-card-content>
        <TableSoal />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
