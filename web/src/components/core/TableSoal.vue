<script>
import { defineComponent } from "vue";
import { useModal } from "vuestic-ui";

export default defineComponent({
  data() {
    const { confirm } = useModal();

    const generateItems = (count) => {
      const users = [
        {
          id: 1,
          name: "Leanne Graham",
          mapel: "Kimia",
          kelas: "10",
          semester: "Ganjil",
        },
        {
          id: 2,
          name: "Ervin Howell",
          mapel: "Kimia",
          kelas: "11",
          semester: "Genap",
        },
        {
          id: 3,
          name: "Clementine Bauch",
          mapel: "Kimia",
          kelas: "12",
          semester: "Ganjil & Genap",
        },
        // Add more users if needed
      ];

      // Generate user Array
      return new Array(count).fill(null).map((_, idx) => {
        const user = { ...users[idx % users.length] };
        user.id = idx + 1;
        return user;
      });
    };

    const onButtonClick = () => {
      confirm({
        blur: true,
        title: "Confirm",
        message: "Are you sure you want to delete this item?",
      });
    };

    const columns = [
      { key: "id", sortable: false },
      { key: "name", sortable: false },
      { key: "mapel", sortable: false },
      { key: "kelas", sortable: false },
      { key: "semester", sortable: false },
    ];

    const items = generateItems(50); // Adjust the count as needed
    const filtered = items.map((item) => item.id);

    return {
      items,
      columns,
      perPage: 5,
      currentPage: 1,
      isTableStriped: true,
      animated: true,
      selectedRows: [],
      rowEventType: "",
      rowId: "",
      filtered,
      loading: true,
      showModal: false,
      input: "",
      onButtonClick,
      isDataLoaded: false,
    };
  },

  computed: {
    pages() {
      return this.perPage && this.perPage !== 0
        ? Math.ceil(this.filtered.length / this.perPage)
        : this.filtered.length;
    },
  },

  methods: {
    async fetchData() {
      // Simulate a delay and then set loading to false
      await new Promise((resolve) => setTimeout(resolve, 600));
      this.loading = false;
      this.isDataLoaded = true; // Set the flag to indicate data is loaded
    },

    async printTable() {
      // Show loading spinner or any indication that the content is being loaded
      this.loading = true;

      // Wait for the data to be loaded
      await this.fetchData();

      // Hide loading spinner
      this.loading = false;

      // Check if data is loaded
      if (this.isDataLoaded) {
        const tableElement = document.querySelector(".table-container table"); // Adjust the selector based on your actual HTML structure
        if (tableElement) {
          // Clone the table to avoid modifying the original
          const clonedTable = tableElement.cloneNode(true);

          // Remove elements you don't want in the print version
          const elementsToRemove = [
            'input[type="checkbox"]',
            ".pagination-container",
          ];
          elementsToRemove.forEach((selector) => {
            const elements = clonedTable.querySelectorAll(selector);
            elements.forEach((element) => element.remove());
          });

          const printWindow = window.open("", "_blank");
          printWindow.document.open();
          printWindow.document.write(`
        <html>
          <head>
            <title>Print Table</title>
            <style>
              body {
                font-family: Arial, sans-serif;
              }
              table {
                width: 100%;
                border-collapse: collapse;
                margin-bottom: 20px;
              }
              th, td {
                border: 1px solid #ddd;
                padding: 8px;
                text-align: left;
              }
              th {
                background-color: #f2f2f2;
              }
            </style>
          </head>
          <body>
            ${clonedTable.outerHTML}
          </body>
        </html>
      `);
          printWindow.document.close();
          printWindow.print();
        } else {
          console.error("Table element not found. Unable to print.");
        }
      } else {
        console.error("Data not loaded. Unable to print.");
      }
    },

    // ... (your other methods)
  },

  watch: {
    currentPage(newPage, oldPage) {
      // Watch for changes in currentPage and update loading state
      if (newPage !== oldPage) {
        this.loading = true;

        // Simulate fetching data and set loading to false after a delay
        this.fetchData();
      }
    },
  },

  mounted() {
    // Initial data fetching
    this.loading = true;
    this.fetchData();
  },
});
</script>

<template>
  <div class="table-container">
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
        <va-button @click="showModal = !showModal" icon="add">Add</va-button>
        <va-button @click="showModal = !showModal" icon="edit">Edit</va-button>
        <va-button @click="printTable" icon="print">Print</va-button>
        <va-button @click="onButtonClick" type="delete" icon="delete"
          >Delete</va-button
        >
      </va-button-group>
    </div>
    <va-data-table
      ref="printTable"
      :items="items"
      :columns="columns"
      :striped="isTableStriped"
      :current-page="currentPage"
      :per-page="perPage"
      selectable
      :animated="animated"
      :delay="500"
      :loading="loading"
    >
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
    </va-data-table>
  </div>
  <!-- Modal Content -->
  <va-modal v-model="showModal" blur size="large" fixed-layout>
    <va-card :bordered="false" stripe>
      <va-card-title>Input Data Alur Tujuan Pembelajaran</va-card-title>
      <va-card-content>
        <div>
          <div
            class="modal-container"
            style="display: grid; grid-template-columns: 1fr 1fr; gap: 10px"
          >
            <div>
              <va-input
                v-model="value"
                placeholder="Nama Penyusun"
                label="Nama Penyusun"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div>
              <va-select
                v-model="value"
                :options="options"
                label="Mata Pelajaran"
                placeholder="Pilih Mata Pelajaran"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div>
              <va-select
                v-model="value"
                :options="options"
                label="Semester"
                placeholder="Pilih Semester"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div>
              <va-select
                v-model="value"
                :options="options"
                label="Kelas"
                placeholder="Pilih Kelas"
                preset="bordered"
                style="width: 100%"
              />
            </div>
          </div>
          <va-card :bordered="false" stripe disabled>
            <va-card-title>Upload Data Soal</va-card-title>
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
</template>

<style>
.table-container {
  border: shadow black;
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
</style>
