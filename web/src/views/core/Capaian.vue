<script setup lang="ts">
import TableCrud from "@/components/core/Table.vue";

import { ref } from "vue";
import { useModal } from "vuestic-ui";

const { confirm } = useModal();
const onButtonClick = () => {
  confirm({
    blur: true,
    title: "Confirm",
    message: "Are you sure you want to delete this item?",
  });
};

const basic = ref([]);
const value = ref("");
const showModal = ref(false);
const searchQuery = ref("");
</script>

<template>
  <div>
    <va-card outlined stripe stripe-color="danger" class="c-1">
      <va-card-title>Data Capaian Pendidikan</va-card-title>
      <div class="row">
        <div class="sb-1">
          <va-input
            v-model="searchQuery"
            placeholder="Search"
            preset="solid"
            loading
          >
            <template #prependInner>
              <va-icon name="search" color="secondary" />
            </template>
          </va-input>
        </div>
        <va-spacer />
        <div class="bg-1">
          <va-button-group
            icon-color="#000000"
            preset="secondary"
            border-color="transparent"
            size="small"
          >
            <va-button @click="showModal = !showModal" icon="add"
              >Add</va-button
            >
            <va-modal v-model="showModal" blur>
              <va-card :bordered="false" color="warning" gradient>
                <va-card-title>Input Data Capaian</va-card-title>
                <va-card-content>
                  <div>
                    <div class="container">
                      <div class="flex flex-col md6">
                        <va-input
                          v-model="value"
                          placeholder="nama penyusun"
                          label="Nama"
                          preset="bordered"
                        />
                      </div>
                      <div class="flex flex-col md6">
                        <va-input
                          v-model="value"
                          placeholder="nama mata pelajaran"
                          label="Mata Pelajaran"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <div class="container">
                    <div class="flex flex-col md6">
                        <va-input
                          v-model="value"
                          placeholder="nama penyusun"
                          label="Nama"
                          preset="bordered"
                        />
                      </div>
                      <div class="flex flex-col md6">
                        <va-input
                          v-model="value"
                          placeholder="nama mata pelajaran"
                          label="Mata Pelajaran"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <va-card>
                      <va-card-content>
                        <va-file-upload
                          v-model="basic"
                          dropzone
                          file-types="doc,docs,rtf,xls,xlsx,ppt,pptx,pdf,txt"
                        />
                        <va-spacer vertical />
                        <va-file-upload
                          v-model="basic"
                          dropzone
                          file-types="doc,docs,rtf,xls,xlsx,ppt,pptx,pdf,txt"
                        />
                      </va-card-content>
                    </va-card>
                  </div>
                </va-card-content>
              </va-card>
            </va-modal>
            <va-button icon="edit">Edit</va-button>
            <va-button @click="onButtonClick" type="delete" icon="delete"
              >Delete</va-button
            >
          </va-button-group>
        </div>
      </div>
      <va-card-content>
        <TableCrud :searchQuery="searchQuery" />
      </va-card-content>
    </va-card>
  </div>
</template>

<style>
.c-1 {
  margin-top: 10px;
  margin-left: 60px;
  margin-right: 60px;
}

.bg-1 {
  margin: 20px;
  margin-bottom: 10px;
  margin-top: 10px;
}

.sb-1 {
  margin: 20px;
  margin-bottom: 5px;
  margin-top: 5px;
}

.row {
  display: flex;
  align-items: center;
  size-adjust: auto;
}
</style>

<style scoped>
.container {
  display: flex;
  size-adjust: auto;
  margin-bottom: 20px;
  width: 100%;
}

</style>