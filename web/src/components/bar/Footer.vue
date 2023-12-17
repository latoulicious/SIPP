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
            <a href="https://github.com/latoulicious/SIPP">
              <img src="@/assets/github.svg" alt="GitHub" class="logo" />
            </a>
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

<style lang="scss">
.footer-wrapper {
  position: fixed;
  bottom: 0;
  width: 100%;
}

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
  margin: 0;
}

.va-navbar-item {
  display: flex;
  align-items: center;
}

.center-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.placeholder {
  width: 18px;
  height: 18px;
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

/* Adjust the margin to fine-tune centering */
.center-content {
  margin-left: 35px;
  margin-right: auto;
}
</style>
