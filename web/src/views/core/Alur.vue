<script setup lang="ts">
import TableAlur from "@/components/core/TableAlur.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Alur Tujuan Pembelajaran</va-card-title>
      <va-card-content>
        <TableAlur />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
