<template>
  <CorePageLayout>
    <div class="centered-container">
      <div class="not-found-container">
        <img class="not-found" src="@/assets/404.svg" alt="404" />
        <p class="text">
          <a href="/">Home</a>
        </p>
      </div>
    </div>
  </CorePageLayout>
</template>

<script>
export default {
  name: "NotFound",
};
</script>

<style scoped>
.centered-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh; /* 100% of the viewport height */
  max-width: 1600px; /* Set the maximum width to match the navbar container size */
  margin: 0 auto; /* Center the container horizontally */
}

.not-found-container {
  text-align: center;
  margin-bottom: auto;
  margin-top: 20px;
}

.text {
  font-family: fantasy;
  font-size: 1em;
  margin-top: 20px;
}

.not-found {
  width: 60%;
  margin-bottom: 20px; /* Adjusted margin for better spacing */
  margin-top: 20px;
}
</style>
