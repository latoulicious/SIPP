<script setup>
import TablePengayaan from "@/components/core/TablePengayaan.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Data Pengayaan</va-card-title>
      <va-card-content>
        <TablePengayaan />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
