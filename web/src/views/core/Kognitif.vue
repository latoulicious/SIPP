<script setup lang="ts">
import TableKognitif from "@/components/core/TableKognitif.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Asesmen Kognitif</va-card-title>
      <va-card-content>
        <TableKognitif />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
