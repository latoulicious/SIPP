<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const jwtToken = ref(localStorage.getItem("jwtToken"));
const router = useRouter();

onMounted(() => {
  // Check if the user is authenticated when the component is mounted
  if (!jwtToken.value && router.currentRoute.value.name !== "login") {
    router.push({ name: "login" });
  }
});
</script>

<template>
  <RouterView />
</template>

<style scoped></style>
