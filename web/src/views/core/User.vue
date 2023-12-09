<script setup>
import TableUser from "@/components/core/TableUser.vue";
</script>

<template>
  <va-card class="container" stripe stripe-color="danger">
    <va-card-title>Manage Users</va-card-title>
    <va-card-content>
      <TableUser />
    </va-card-content>
  </va-card>
</template>

<style>
.container {
  margin-left: 60px;
  margin-right: 60px;
  margin-bottom: 10px;
  margin-top: 10px;
  border: shadow black;
}
</style>
