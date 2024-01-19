<script setup>
import TableMapel from "@/components/core/TableMapel.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Mata Pelajaran</va-card-title>
      <va-card-content>
        <TableMapel />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
