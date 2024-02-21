<template>
  <div class="carousel-wrapper">
    <va-carousel
      v-model="value"
      :items="items.length > 0 ? items : ['https://picsum.photos/1500']"
      :indicators="false"
      infinite
      height="400px"
      stateful
      autoscroll
      :autoscroll-interval="3000"
    >
      <template v-slot:item="{ item }">
        <VaSkeleton variant="squared" height="10rem" v-if="!item.src" />
        <img v-else :src="item.src" alt="" />
      </template>
    </va-carousel>
  </div>
</template>

<script>
export default {
  data() {
    return {
      value: 0,
      items: [],
    };
  },
  mounted() {
    // Simulate asynchronous loading of carousel items
    setTimeout(() => {
      this.items = [
        { src: "https://picsum.photos/1500" },
        { src: "https://picsum.photos/1501" },
        { src: "https://picsum.photos/1502" },
        { src: "https://picsum.photos/1503" },
        { src: "https://picsum.photos/1504" },
      ];
    }, 2000);
  },
};
</script>

<style scoped>
.carousel-wrapper {
  margin-top: 20px;
  margin-left: 60px;
  margin-right: 60px;
  margin-bottom: 20px;
}
</style>
