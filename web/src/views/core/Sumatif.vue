<script setup>
import TableSumatif from "@/components/core/TableSumatif.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Data Asesmen Sumatif</va-card-title>
      <va-card-content>
        <TableSumatif />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
