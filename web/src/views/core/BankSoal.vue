<script setup lang="ts">
import TableBankSoal from "@/components/core/TableBankSoal.vue";
</script>

<template>
  <div class="container">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Bank Soal</va-card-title>
      <va-card-content>
        <TableBankSoal />
      </va-card-content>
    </va-card>
  </div>
</template>

<style>
.container {
  margin-top: 10px;
  margin-left: 60px;
  margin-right: 60px;
}
</style>
