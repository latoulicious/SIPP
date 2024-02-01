<script setup>
import TableTahun from "@/components/core/TableTahun.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Tahun Ajaran</va-card-title>
      <va-card-content>
        <TableTahun />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
