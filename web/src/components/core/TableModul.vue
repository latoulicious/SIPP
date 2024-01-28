<script>
import { defineComponent } from "vue";
import { ref } from "vue";
import pdfMake from "pdfmake/build/pdfmake";
import pdfFonts from "pdfmake/build/vfs_fonts";
import axios from "axios";

pdfMake.vfs = pdfFonts.pdfMake.vfs;

const defaultItem = {
  Elemen: "", // Change to match the server property name
  LingkupMateri: "", // Change to match the server property name
  TujuanPembelajaran: "", // Change to match the server property name
  KodeTP: "", // Change to match the server property name
  AlokasiWaktu: "", // Change to match the server property name
  SumberBelajar: "", // Change to match the server property name
  ProjekPPancasila: "", // Change to match the server property name
};

const displayNames = {
  Elemen: "Elemen", // Change to match the server property name
  LingkupMateri: "Lingkup Materi", // Change to match the server property name
  TujuanPembelajaran: "Tujuan Pembelajaran", // Change to match the server property name
  KodeTP: "Kode TP", // Change to match the server property name
  AlokasiWaktu: "Alokasi Waktu", // Change to match the server property name
  SumberBelajar: "Sumber Belajar", // Change to match the server property name
  ProjekPPancasila: "Projek Profile Pancasila", // Change to match the server property name
};

function imageFileToDataUrl(imageFile) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = function (event) {
      resolve(event.target.result);
    };

    reader.onerror = function (error) {
      reject(error);
    };

    reader.readAsDataURL(imageFile);
  });
}

export default defineComponent({
  data() {
    const columns = [
      { key: "KodeTP", sortable: false },
      { key: "Elemen", sortable: false },
      { key: "LingkupMateri", sortable: false },
      { key: "AlokasiWaktu", sortable: false },
      { key: "actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: { ...defaultItem },
      input: ref({ value: "" }),
      items: [],
      showModal: false,
      viewModalVisible: false,
      detailItem: null,
      detailModalVisible: false,
      displayNames,
    };
  },

  computed: {
    filteredDetailFields() {
      return Object.keys(this.detailItem).filter(
        (key) =>
          !["created_at", "updated_at", "deleted_at", "id"].includes(key),
      );
    },

    filteredDisplayNames() {
      return Object.fromEntries(
        Object.entries(this.displayNames).filter(
          ([key]) =>
            !["created_at", "updated_at", "deleted_at", "id"].includes(key),
        ),
      );
    },

    filteredInputFields() {
      return this.inputFields.filter((key) => key);
    },

    isViewModalVisible() {
      return this.viewModalVisible;
    },

    filteredItems() {
      const searchTerms = this.input.value.toLowerCase().trim().split(/\s+/);

      if (searchTerms.length === 0) {
        return this.items; // No search term, return all items
      } else {
        return this.items.filter((item) => {
          return searchTerms.every((term) => {
            return Object.values(item).some((value) => {
              // Check if any property value contains the search term
              return String(value).toLowerCase().includes(term);
            });
          });
        });
      }
    },

    isNewData() {
      return Object.keys(this.createdItem).every(
        (key) => !!this.createdItem[key],
      );
    },

    inputFields() {
      return Object.keys(this.createdItem);
    },
  },

  methods: {
    async fetchData() {
      try {
        const response = await axios.get("http://localhost:3000/api/capaian");

        if (!response.data.data) {
          console.error("No data received from the server:", response);
          return;
        }

        this.items = response.data.data.map((alur) => ({
          ...alur,
          id: alur.ID,
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },

    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        const response = await axios.post(
          "http://localhost:3000/api/capaian",
          this.createdItem,
        );

        if (!response.data.data) {
          console.error("No data received from the server:", response);
          return;
        }

        this.items.push({
          ...this.createdItem,
          id: response.data.data.ID,
        });

        this.resetCreatedItem();
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },

    async editItem() {
      try {
        // Create a deep copy of the edited item
        const editedData = JSON.parse(JSON.stringify(this.editedItem));

        const response = await axios.put(
          `http://localhost:3000/api/capaian/${this.editedItem.id}`,
          editedData,
        );

        if (!response.data.data) {
          console.error("No data received from the server:", response);
          return;
        }

        this.items = this.items.map((item) =>
          item.id === this.editedItem.id ? { ...item, ...editedData } : item,
        );

        this.resetEditedItem();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          console.log("Before axios.delete");
          await axios.delete(`http://localhost:3000/api/capaian/${id}`);
          console.log("After axios.delete");

          this.items = this.items.filter((item) => item.id !== id);
        } catch (error) {
          console.error("Error deleting item:", error);
        }
      }
    },

    resetEditedItem() {
      this.editedItem = null;
      this.editedItemId = null;
    },

    resetCreatedItem() {
      this.createdItem = { ...defaultItem };
      this.showModal = false;
    },

    openModalToEditItemById(id) {
      this.editedItemId = id;
      this.editedItem = { ...this.items[id], id: this.items[id].id };
    },

    toggleAddModal() {
      this.showModal = !this.showModal;
      if (!this.showModal) {
        this.resetCreatedItem();
      }
    },

    openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      axios
        .get(`http://localhost:3000/api/capaian/${selectedItemId}`)
        .then((response) => {
          const data = response.data.data;
          if (data) {
            // Include only the fields you want from the server response
            this.detailItem = {
              Elemen: data.Elemen,
              LingkupMateri: data.LingkupMateri,
              TujuanPembelajaran: data.TujuanPembelajaran,
              KodeTP: data.KodeTP,
              AlokasiWaktu: data.AlokasiWaktu,
              SumberBelajar: data.SumberBelajar,
              ProjekPPancasila: data.ProjekPPancasila,
              // Add more fields as needed
              id: selectedItemId,
            };
            this.detailModalVisible = true;
            console.log("Detail modal data:", this.detailItem);
          } else {
            console.error("No data received from the server");
          }
        })
        .catch((error) => {
          console.error("Error fetching data for the detail modal:", error);
        });
    },

    resetDetailItem() {
      this.detailItem = null;
      this.detailModalVisible = false;
    },

    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;

      try {
        // Fetch the necessary data directly from the server
        const response = await axios.get(
          `http://localhost:3000/api/capaian/${selectedItemId}`,
        );
        const data = response.data.data;

        if (data) {
          console.log("Printing row with ID:", selectedItemId);

          const imagePath = "../../src/assets/login.png";
          const imageFile = await fetch(imagePath).then((response) =>
            response.blob(),
          );
          const logoDataUrl = await imageFileToDataUrl(imageFile);

          const docDefinition = {
            footer: function (currentPage, pageCount) {
              return [
                {
                  text: currentPage.toString() + " of " + pageCount,
                  alignment: "center", // Align the text to the center
                  fontSize: 8, // Use a smaller font size
                  margin: [10, 10, 10, 0], // Add custom margins (top, right, bottom, left)
                },
              ];
            },
            // header: function (currentPage) {
            //   return [
            //     {
            //       stack: [
            //         {
            //           image: logoDataUrl,
            //           fit: [80, 80], // Adjust the width and height as needed
            //           alignment: currentPage % 2 ? "left" : "right",
            //         },
            //         // Add other header elements if needed
            //       ],
            //       margin: [10, 20, 10, 10], // Adjust the margins (top, right, bottom, left) for the entire stack
            //     },
            //   ];
            // },
            content: [
              {
                text: "Alur Tujuan Pembelajaran",
                fontSize: 12,
                bold: true,
                alignment: "center",
                margin: [0, 20, 0, 20],
              },
              {
                table: {
                  headerRows: 1,
                  widths: [
                    "auto",
                    "auto",
                    "auto",
                    "auto",
                    "auto",
                    "auto",
                    "auto",
                  ],
                  body: [
                    [
                      { text: "Element", fontSize: 10, bold: true },
                      { text: "Lingkup Materi", fontSize: 10, bold: true },
                      { text: "Tujuan Pembelajaran", fontSize: 10, bold: true },
                      { text: "Kode TP", fontSize: 10, bold: true },
                      { text: "Alokasi Waktu", fontSize: 10, bold: true },
                      { text: "Sumber Belajar", fontSize: 10, bold: true },
                      { text: "Projek Pancasila", fontSize: 10, bold: true },
                    ],
                    [
                      data.Elemen,
                      data.LingkupMateri,
                      data.TujuanPembelajaran,
                      data.KodeTP,
                      data.AlokasiWaktu,
                      data.SumberBelajar,
                      data.ProjekPPancasila,
                    ],
                  ],
                },
                margin: [0, 0, 0, 20], // Adjust the margin for the table
              },
            ],

            pageSize: "A4",
            pageMargins: [40, 60, 40, 60],
          };

          const pdf = pdfMake.createPdf(docDefinition);

          // Open the PDF for printing
          pdf.open();

          // Alternatively, save the PDF as a file
          // pdf.download(`alur_${selectedItemId}.pdf`);
        } else {
          console.error("No data received from the server");
        }
      } catch (error) {
        console.error("Error fetching data for printing:", error);
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
    <va-input v-model="input.value" placeholder="Search"></va-input>
    <va-button-group
      icon-color="#000000"
      preset="secondary"
      border-color="#000000"
    >
      <va-button @click="toggleAddModal" preset="secondary" icon="add"
        >Add Alur Tujuan Pembelajaran</va-button
      >
    </va-button-group>
  </div>
  <div>
    <va-data-table :items="filteredItems" :columns="columns" striped>
      <template #cell(actions)="{ rowIndex }">
        <div class="action-buttons">
          <va-button preset="plain" icon="print" @click="printRow(rowIndex)" />
          <va-button
            preset="plain"
            icon="remove_red_eye"
            @click="openDetailModal(rowIndex)"
          />
          <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          />
          <va-button
            preset="plain"
            icon="delete"
            @click="deleteItemById(filteredItems[rowIndex].id)"
          />
        </div>
      </template>
    </va-data-table>

    <!-- Modal Content -->
    <va-modal
      blur
      class="modal-crud"
      stripe
      title="Add Alur Tujuan Pembelajaran"
      size="large"
      :model-value="showModal"
      @ok="addNewItem"
      @cancel="resetCreatedItem"
    >
      <va-textarea
        v-for="key in inputFields"
        :key="key"
        :label="displayNames[key]"
        v-model="createdItem[key]"
        class="my-6"
      />
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      :model-value="!!editedItem"
      title="Edit Alur Tujuan Pembelajaran"
      size="large"
      @ok="editItem"
      @cancel="resetEditedItem"
    >
      <va-textarea
        v-for="key in filteredInputFields"
        :key="key"
        :label="displayNames[key]"
        v-model="editedItem[key]"
        class="my-6"
      />
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      stripe
      title="Detail Alur Tujuan Pembelajaran"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <va-textarea
        v-for="key in filteredDetailFields"
        :key="key"
        :label="filteredDisplayNames[key]"
        v-model="detailItem[key]"
        class="my-6"
        readonly
      />
    </va-modal>
  </div>
</template>

<style>
.action-buttons {
  display: flex;
  gap: 8px; /* Adjust the gap to your preference */
}

.va-input {
  display: block;
  margin-bottom: 10px;
}

.modal-crud {
  .va-textarea {
    width: 100%;
    display: flex;
    box-sizing: border-box;
    margin-bottom: 20px;
  }
}
</style>
