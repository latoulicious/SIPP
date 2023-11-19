<script setup lang="ts">
/* import Table from "@/components/core/Table.vue"; */
import tabletest from "@/components/core/tabletest.vue";

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
const input = ref("");
</script>

<template>
  <div>
    <va-card outlined stripe stripe-color="danger" class="c-1">
      <va-card-title>Data Capaian Pendidikan</va-card-title>
      <div class="row">
        <div class="sb-1">
          <va-input v-model="input" placeholder="Filter..." class="w-full" />
          <!-- <va-input
            v-model="searchQuery"
            placeholder="Search"
            preset="solid"
            loading
          >
            <template #prependInner>
              <va-icon name="search" color="secondary" />
            </template>
          </va-input> -->
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
            <va-modal v-model="showModal" blur size="large" fixed-layout>
              <va-card :bordered="false" stripe>
                <va-card-title>Input Data Capaian</va-card-title>
                <va-card-content>
                  <div>
                    <div class="container">
                      <div>
                        <va-input
                          v-model="value"
                          placeholder="nama penyusun"
                          label="Nama"
                          preset="bordered"
                          style="width: 100%"
                        />
                      </div>
                      <div style="margin-top: 10px">
                        <va-input
                          v-model="value"
                          placeholder="nama mata pelajaran"
                          label="Mata Pelajaran"
                          preset="bordered"
                          style="width: 100%"
                        />
                      </div>
                    </div>
                    <div class="txt flex justify-between">
                      <div
                        class="flex flex-col md6"
                        style="margin-right: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Judul Elemen"
                          placeholder="Judul mengenai elemen pemahaman suatu mata pelajaran"
                          preset="bordered"
                        />
                      </div>
                      <div
                        class="flex flex-col md6"
                        style="margin-left: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Elemen"
                          placeholder="Menjelaskan inti dari judul elemen tersebut"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <div class="txt flex justify-between">
                      <div
                        class="flex flex-col md6"
                        style="margin-right: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Mengamati"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                          preset="bordered"
                        />
                      </div>
                      <div
                        class="flex flex-col md6"
                        style="margin-left: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Mempertanyakan"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <div class="txt flex justify-between">
                      <div
                        class="flex flex-col md6"
                        style="margin-right: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Merencanakan"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                          preset="bordered"
                        />
                      </div>
                      <div
                        class="flex flex-col md6"
                        style="margin-left: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Memproses"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <div class="txt flex justify-between">
                      <div
                        class="flex flex-col md6"
                        style="margin-right: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Mengevaluasi"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajarann"
                          preset="bordered"
                        />
                      </div>
                      <div
                        class="flex flex-col md6"
                        style="margin-left: 10px; width: 100%"
                      >
                        <va-textarea
                          v-model="value"
                          label="Keterangan Proses Mengkomunikasikan"
                          placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                          preset="bordered"
                        />
                      </div>
                    </div>
                    <va-card :bordered="false" stripe disabled>
                      <va-card-title>Upload Data Capaian</va-card-title>
                      <va-card-content>
                        <va-file-upload
                          v-model="basic"
                          dropzone
                          undo
                          undo-button-text="Restore"
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
        <tabletest />
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
  margin-bottom: 10px;
}

.txt {
  display: flex;
  justify-content: space-between;
}

.va-textarea {
  width: 100%;
  box-sizing: border-box;
  margin-bottom: 20px;
}
</style>
