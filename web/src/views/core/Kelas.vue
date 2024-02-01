<script setup>
import TableKelas from "@/components/core/TableKelas.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Kelas</va-card-title>
      <va-card-content>
        <TableKelas />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
