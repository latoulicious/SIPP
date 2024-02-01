<script setup lang="ts">
import TablePenilaian from "@/components/core/TablePenilaian.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Penilaian Ranah Pengetahuan</va-card-title>
      <va-card-content>
        <TablePenilaian />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
