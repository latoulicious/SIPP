<script setup>
import TableJurusan from "@/components/core/TableJurusan.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Jurusan</va-card-title>
      <va-card-content>
        <TableJurusan />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
