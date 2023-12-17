<template>
  <CorePageLayout>
    <div class="centered-container">
      <div class="unauthorized-container">
        <img class="unauthorized" src="@/assets/401.png" alt="404" />
        <p class="text">
          Sorry, you don't have permission to access this page
          <br />
          <br />
          <a href="/">Go home?</a>
        </p>
      </div>
    </div>
  </CorePageLayout>
</template>

<script>
export default {
  name: "Unauthorized",
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

.unauthorized-container {
  text-align: center;
  margin-bottom: auto;
  margin-top: 20px;
}

.text {
  font-family: fantasy;
  font-size: 2em;
  margin-top: 20px;
}

.unauthorized {
  width: 65%; /* Adjusted to scale the image proportionally */
  margin-bottom: 20px; /* Adjusted margin for better spacing */
  margin-top: 20px;
}
</style>
