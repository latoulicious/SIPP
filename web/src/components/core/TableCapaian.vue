<script>
import { defineComponent } from "vue";
import { useModal } from "vuestic-ui";
import axios from "axios";

export default defineComponent({
  data() {
    const { confirm } = useModal();

    const onButtonClick = () => {
      confirm({
        blur: true,
        title: "Confirm",
        message: "Are you sure you want to delete this item?",
      });
    };

    const columns = [
      { key: "Name", sortable: false },
      { key: "Mapel", sortable: false },
      { key: "Kelas", sortable: false },
      { key: "Tahun", sortable: false },
      { key: "actions", width: 80 },
    ];

    return {
      items: [],
      createdItem: {
        name: "",
        mapel: "",
        kelas: null,
        tahun: "",
        judul_elemen: "",
        ket_elemen: "",
        ket_proses_mengamati: "",
        ket_proses_mempertanyakan: "",
        ket_proses_merencanakan: "",
        ket_proses_memproses: "",
        ket_proses_mengevaluasi: "",
        ket_proses_mengkomunikasikan: "",
      },
      columns,
      options: ["10", "11", "12"],
      perPage: 10,
      currentPage: 1,
      isTableStriped: true,
      animated: true,
      selectedRows: [],
      rowEventType: "",
      rowId: "",
      showModal: false,
      input: "",
      onButtonClick,
    };
  },

  computed: {
    pages() {
      return this.items &&
        this.items.length &&
        this.perPage &&
        this.perPage !== 0
        ? Math.ceil(this.items.length / this.perPage)
        : this.items
          ? this.items.length
          : 0;
    },
  },

  methods: {
    async fetchData() {
      try {
        const response = await axios.get("http://localhost:3000/api/capaian");

        // Check if response.data.data is an array
        if (Array.isArray(response.data.data)) {
          this.items = response.data.data;
        } else {
          console.error(
            "Response data.data is not an array:",
            response.data.data,
          );
        }

        this.loading = false;
      } catch (error) {
        console.error("Error fetching data:", error);
        this.loading = false;
      }
    },

    async addNewItem() {
      try {
        // Prepare the payload to send to the backend
        const payload = {
          ...this.createdItem,
          Kelas: parseInt(this.createdItem.kelas), // Convert to integer if needed
          Tahun: parseInt(this.createdItem.tahun), // Convert to integer if needed
        };

        console.log("Payload:", payload);

        const response = await axios.post(
          "http://localhost:3000/api/capaian",
          payload,
        );

        console.log("Response:", response.data);

        // Assuming the server returns the newly created item
        const newItem = response.data.data;

        // Add the new item to the items array
        this.items.push(newItem);

        // Reset the createdItem for the next input
        this.resetCreatedItem();

        this.fetchData();
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },
    resetCreatedItem() {
      // Reset the createdItem to its initial state
      this.createdItem = {
        name: "",
        mapel: "",
        kelas: null,
        tahun: "",
        judul_elemen: "",
        ket_elemen: "",
        ket_proses_mengamati: "",
        ket_proses_mempertanyakan: "",
        ket_proses_merencanakan: "",
        ket_proses_memproses: "",
        ket_proses_mengevaluasi: "",
        ket_proses_mengkomunikasikan: "",
      };
    },

    toggleAddModal() {
      this.showModal = !this.showModal;
      if (!this.showModal) {
        this.addNewItem();
        this.resetCreatedItem();
      }
    },
  },

  mounted() {
    this.fetchData();
  },
});
</script>

<template>
  <div
    class="header-container"
    style="display: flex; justify-content: space-between; align-items: center"
  >
    <va-input v-model="input" placeholder="Search"></va-input>
    <va-button-group
      icon-color="#000000"
      preset="secondary"
      border-color="bordered"
    >
      <va-button @click="toggleAddModal" icon="add">Add</va-button>
    </va-button-group>
  </div>
  <va-modal v-model="showModal" blur size="large" fixed-layout @ok="addNewItem">
    <va-card :bordered="false" stripe>
      <va-card-title>Input Data Capaian Pembelajaran</va-card-title>
      <va-card-content>
        <div>
          <div class="modal-container">
            <div>
              <va-input
                v-model="createdItem.name"
                placeholder="Nama Penyusun Capaian Pembelajaran"
                label="Nama Penyusun"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div style="margin-top: 10px">
              <va-input
                v-model="createdItem.mapel"
                label="Mata Pelajaran"
                placeholder="Mata Pelajaran"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div style="margin-top: 10px">
              <va-select
                v-model.number="createdItem.kelas"
                :options="options"
                label="Kelas"
                placeholder="Pilih Kelas"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div style="margin-top: 10px">
              <va-input
                v-model.number="createdItem.tahun"
                label="Tahun Ajaran"
                placeholder="Tahun Ajar"
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
                v-model="createdItem.judul_elemen"
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
                v-model="createdItem.ket_elemen"
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
                v-model="createdItem.ket_proses_mengamati"
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
                v-model="createdItem.ket_proses_mempertanyakan"
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
                v-model="createdItem.ket_proses_merencanakan"
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
                v-model="createdItem.ket_proses_memproses"
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
                v-model="createdItem.ket_proses_mengevaluasi"
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
                v-model="createdItem.ket_proses_mengkomunikasikan"
                label="Keterangan Proses Mengkomunikasikan"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                preset="bordered"
              />
            </div>
          </div>
          <va-card :bordered="false" stripe disabled>
            <va-card-title>Upload Data Capaian Pembelajaran</va-card-title>
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
  <div class="table-container">
    <vaDataTable
      :items="items"
      :columns="columns"
      :striped="isTableStriped"
      :current-page="currentPage"
      :per-page="perPage"
      selectable
      :animated="animated"
      :delay="500"
    >
      <template #cell(actions)>
        <div class="action-buttons">
          <va-button preset="plain" icon="remove_red_eye" @click="" />
          <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          />
          <va-button
            preset="plain"
            icon="delete"
            @click="deleteItemById(rowIndex)"
          />
        </div>
      </template>
      <template #bodyAppend>
        <tr>
          <td colspan="6">
            <div class="flex justify-center mt-4">
              <div class="pagination-container">
                <va-pagination v-model="currentPage" :pages="pages" />
              </div>
            </div>
          </td>
        </tr>
      </template>
      <template #bodyCellCheckbox="{ value }">
        <input type="checkbox" v-model="selectedRows" :value="value" />
      </template>
    </vaDataTable>
  </div>
</template>

<style>
.table-container {
  border: solid black;
}

.action-buttons {
  display: flex;
  gap: 8px;
  /* Adjust the gap to your preference */
}

.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 10px;
}

.header-container {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>

<style scoped>
.modal-container {
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
