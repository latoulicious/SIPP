<script>
export default {
  data() {
    return {
      information: [
        {
          text: "Perayaan Hari Guru Nasional (HGN) di SMAN 3 Subang.",
        },
        {
          text: "Juara Harapan 1 Pasanggiri Gamelan Degung Klasik Antar SMK/SMK Tingkat Propinsi.",
        },
        {
          text: "LDKS SMAN 3 Subang 2023.",
        },
        {
          text: "Siswi SMAN 3 Subang Aisyah Ditasya Putri Wakili Indonesia di event Internasional",
        },
      ],
    };
  },
};
</script>

<template>
  <div class="container">
    <VaCard stripe stripe-color="primary">
      <VaCardTitle>Informasi Umum</VaCardTitle>
      <va-card-content>
        <va-list fit>
          <VaListItem
            v-for="(information, index) in information"
            :key="index"
            class="list__item"
          >
            <VaListItemSection avatar>
              <VaSkeleton variant="circle" height="4rem" />
            </VaListItemSection>

            <VaListItemSection>
              <VaListItemLabel
                caption
                :lines="index + 1"
                class="text-black text-justify"
              >
                {{ information.text }}
              </VaListItemLabel>
            </VaListItemSection>
          </VaListItem>
        </va-list>
      </va-card-content>
    </VaCard>
  </div>
</template>

<style lang="scss" scoped>
@import "vuestic-ui/styles/grid";

.item {
  border: 1.5px;
  border-color: shadow black;
  margin: 5px;
}

.text-black {
  color: black;
}

.text-justify {
  text-align: justify;
}
</style>

<style lang="scss">
.list__item + .list__item {
  margin-top: 20px;
}
</style>

<style>
.container {
  display: grid;
  margin-left: 60px;
  margin-right: 60px;
  margin-bottom: 10px;
  margin-top: 10px;
}

.card-container {
  border: shadow black;
}
</style>
