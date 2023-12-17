<script setup>
import TableUser from "@/components/core/TableUser.vue";
</script>

<template>
  <CorePageLayout>
    <va-card class="scaled-card" stripe stripe-color="danger">
      <va-card-title>Manage Users</va-card-title>
      <va-card-content>
        <TableUser />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
