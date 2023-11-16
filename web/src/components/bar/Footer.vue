<template>
  <div class="footer-wrapper">
    <VaNavbar class="a">
      <template #left>
        <!-- Placeholder for left content -->
        <div class="placeholder"></div>
      </template>
      <template #center>
        <va-navbar-item>
          <div class="center-content">
            <p>
              <a href="https://github.com/latoulicious/SIPP">
                <img src="@/assets/github.svg" alt="GitHub" class="logo" />
              </a>
            </p>
            <p>&copy; latoulicious &bull; 2023</p>
          </div>
        </va-navbar-item>
      </template>
      <template #right>
        <!-- Placeholder for right content -->
        <div class="placeholder2"></div>
      </template>
    </VaNavbar>
  </div>
</template>

<style lang="scss" scoped>
.a {
  padding-top: 0.3rem;
  padding-bottom: 0.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  height: 20px;
  max-height: 20px;
  width: auto;
  margin: 0; /* Remove any default margin */
}

.va-navbar-item {
  display: flex;
  align-items: center;
}

.center-content {
  display: flex;
  flex-direction: column; /* Align items in a column */
  align-items: center;
}

/* If you want to center the placeholder elements in the #left and #right sections */
.placeholder {
  width: 18px; /* Adjust the width as needed */
  height: 18px; /* Adjust the height as needed */
}

.placeholder2 {
  width: 54.4333px;
  height: 20px;
}

#right,
#left {
  display: flex;
  align-items: center;
}
</style>
