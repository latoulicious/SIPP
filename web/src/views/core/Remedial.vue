<script setup lang="ts">
import TableRemedial from "@/components/core/TableRemedial.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Remedial</va-card-title>
      <va-card-content>
        <TableRemedial />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
