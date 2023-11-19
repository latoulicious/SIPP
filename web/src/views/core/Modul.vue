<script setup lang="ts">
import testtable from "@/components/core/testtable.vue";
</script>

<template>
  <div class="c-2">
    <va-card stripe stripe-color="danger">
      <va-card-title>Data Modul Ajar</va-card-title>
      <va-card-component>
        <testtable />
      </va-card-component>
    </va-card>
  </div>
</template>

<style>
.c-2 {
  margin-top: 10px;
  margin-left: 60px;
  margin-right: 60px;
}
</style>
