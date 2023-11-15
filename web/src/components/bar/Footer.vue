<template>
  <va-footer>
    <div class="footer-content">
      <div class="footer-text">
        <p>
          <span class="name">placeholder</span>
          <span class="divider"> • </span>
          <a
            href="https://github.com/latoulicious/SIPP"
            target="_blank"
            rel="noopener noreferrer"
          >
            <img src="@/assets/github.svg" alt="GitHub" class="github-icon" />
          </a>
        </p>
        <p class="year">&copy; 2023</p>
      </div>
    </div>
  </va-footer>
</template>

<style scoped>
.footer-content {
  position: fixed;
  bottom: 0;
  width: 100%;
  text-align: center;
  padding: 10px;
  background-color: var(--va-background-primary);
  color: var(--va-text-primary);
}

.footer-text {
  line-height: 1.5; /* Adjust as needed for spacing between lines */
}

.name,
.year {
  font-size: 16px; /* Adjust font size as needed */
  line-height: 20px; /* Adjust line height to match the height of the GitHub icon */
  vertical-align: middle; /* Align text vertically with the GitHub icon */
}

.divider {
  margin: 0 10px; /* Adjust as needed for spacing around the divider */
}

.github-icon {
  height: 20px; /* Set the desired height */
  vertical-align: middle; /* Align the GitHub icon vertically with the text */
}
</style>
