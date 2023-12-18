<script setup lang="ts">
import TableBankSoal from "@/components/core/TableBankSoal.vue";
</script>

<template>
  <CorePageLayout>
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Bank Soal</va-card-title>
      <va-card-content>
        <TableBankSoal />
      </va-card-content>
    </va-card>
  </CorePageLayout>
</template>
