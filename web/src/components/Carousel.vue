<template>
  <div class="carousel-wrapper">
    <va-carousel
      v-model="value"
      :items="items"
      :indicators="false"
      infinite
      height="350px"
      stateful
      autoscroll
      :autoscroll-interval="3000"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      value: 0,
      items: [
        "https://picsum.photos/1500",
        "https://picsum.photos/1501",
        "https://picsum.photos/1502",
        "https://picsum.photos/1503",
        "https://picsum.photos/1504",
      ],
    };
  },
};
</script>

<style scoped>
.carousel-wrapper {
  margin-top: 20px;
  margin-left: 80px;
  margin-right: 80px;
  margin-bottom: 20px;
}
</style>
