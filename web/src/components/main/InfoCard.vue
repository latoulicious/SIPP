<template>
  <div class="container">
    <va-card color="primary" gradient>
      <div class="row">
        <div class="flex flex-col md12">
          <div class="item">
            <va-card>
                <va-card-title>lorem ipsum</va-card-title>
                <va-card-content>
                    asd
                </va-card-content>
            </va-card>
          </div>
        </div>
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
                <va-card-title>lorem ipsum</va-card-title>
                <va-card-content>
                    asd
                </va-card-content>
            </va-card>
          </div>
        </div>
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
                <va-card-title>lorem ipsum</va-card-title>
                <va-card-content>
                    asd
                </va-card-content>
            </va-card>
          </div>
        </div>
        <div class="flex flex-col md4">
          <div class="item">
            <va-card>
                <va-card-title>lorem ipsum</va-card-title>
                <va-card-content>
                    asd
                </va-card-content>
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
  border: 10px solid;
  border-color: black;
  text-align: center;
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
</style>
