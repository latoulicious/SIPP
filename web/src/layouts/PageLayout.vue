<template>
  <div class="core-page-layout">
    <main class="scrollable-content">
      <router-view />
    </main>
  </div>
</template>

<style lang="scss">
.core-page-layout {
  display: flex;
  min-height: 100vh;
}

main {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  margin-left: 40px;
  margin-right: 40px;
}

.Footer {
  margin-top: auto;
}
</style>
