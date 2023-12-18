<script setup lang="ts">
import TableAsesmen from "@/components/core/TableAsesmen.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Asesmen</va-card-title>
      <va-card-content>
        <TableAsesmen />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
