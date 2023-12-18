<script setup lang="ts">
import TablePRP from "@/components/core/TablePRP.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Penilaian Ranah Pengetahuan</va-card-title>
      <va-card-content>
        <TablePRP />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
