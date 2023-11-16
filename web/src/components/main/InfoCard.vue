<template>
  <div class="container">
    <va-card class="card-container">
      <div class="row">
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
              <va-image
                :ratio="16 / 9"
                src="https://opendoodles.s3-us-west-1.amazonaws.com/sprinting.gif"
              />
              <va-card-title>lorem ipsum</va-card-title>
              <va-card-content> asd </va-card-content>
            </va-card>
          </div>
        </div>
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
              <va-image
                :ratio="16 / 9"
                src="https://opendoodles.s3-us-west-1.amazonaws.com/selfie.png"
              />
              <va-card-title>lorem ipsum</va-card-title>
              <va-card-content> asd </va-card-content>
            </va-card>
          </div>
        </div>
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
              <va-image
                :ratio="16 / 9"
                src="https://opendoodles.s3-us-west-1.amazonaws.com/clumsy.png"
              />
              <va-card-title>lorem ipsum</va-card-title>
              <va-card-content> asd </va-card-content>
            </va-card>
          </div>
        </div>
      </div>
    </va-card>
  </div>
</template>

<style lang="scss" scoped>
@import "vuestic-ui/styles/grid";

.item {
  border: 1.5px dashed;
  border-color: shadow black;
  text-align: center;
  margin: 5px;
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
