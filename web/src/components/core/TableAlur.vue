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
    /**
     * Fetches alur data from the API and maps it to the component's items.
     * Handles errors from the API request.
     */
    async fetchData() {
      try {
        const response = await axios.get("http://localhost:3000/api/alur");

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

    /**
     * Adds a new item to the list by making a POST request to the API.
     * Validates input data, handles errors from the API request.
     * Updates component state by pushing new item and resetting form.
     * Refetches latest data after successful creation.
     */
    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        const response = await axios.post(
          "http://localhost:3000/api/alur",
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
        await this.fetchData();
      } catch (error) {
        console.error("Error Creating new item:", error);
      }
    },

    /**
     * Edits an existing item by making a PUT request to the API.
     * Validates input data, handles errors from the API request.
     * Updates component state by mapping over items and replacing the edited item.
     * Refetches latest data after successful edit.
     */
    async editItem() {
      try {
        // Create a deep copy of the edited item
        const editedData = JSON.parse(JSON.stringify(this.editedItem));

        const response = await axios.put(
          `http://localhost:3000/api/alur/${this.editedItem.id}`,
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
        await this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    /**
     * Deletes an item by ID by making a DELETE request to the API.
     * Prompts user to confirm deletion.
     * Handles errors from the API request.
     * Updates component state by filtering out the deleted item.
     * Refetches latest data after successful deletion.
     */
    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          console.log("Before axios.delete");
          await axios.delete(`http://localhost:3000/api/alur/${id}`);
          console.log("After axios.delete");

          this.items = this.items.filter((item) => item.id !== id);

          await this.fetchData();

          // Optionally, you can show a success message
          alert("Item deleted successfully");
        } catch (error) {
          console.error("Error deleting item:", error);
        }
      }
    },

    /**
     * Opens the detail modal for the row at the given index.
     * Gets the item ID for that row.
     * Makes API call to get full item data for that ID.
     * Sets detailItem data with relevant fields from response.
     * Shows modal with detailItem data.
     */
    openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      axios
        .get(`http://localhost:3000/api/alur/${selectedItemId}`)
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

    /**
     * Prints a row from the table to a PDF.
     *
     * Fetches the full data for the row from the API using the row ID.
     * Converts the data into a PDF table using pdfMake.
     * Opens the PDF for printing or saving.
     */
    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;

      try {
        // Fetch the necessary data directly from the server
        const response = await axios.get(
          `http://localhost:3000/api/alur/${selectedItemId}`,
        );
        const data = response.data.data;

        if (data) {
          console.log("Printing row with ID:", selectedItemId);

          const tableBody = [
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
              "Elemen" in data ? data.Elemen : "N/A",
              "LingkupMateri" in data ? data.LingkupMateri : "N/A",
              "TujuanPembelajaran" in data ? data.TujuanPembelajaran : "N/A",
              "KodeTP" in data ? data.KodeTP : "N/A",
              "AlokasiWaktu" in data ? data.AlokasiWaktu : "N/A",
              "SumberBelajar" in data ? data.SumberBelajar : "N/A",
              "ProjekPPancasila" in data ? data.ProjekPPancasila : "N/A",
            ],
          ];

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
                  widths: Array(tableBody[0].length).fill("auto"),
                  body: tableBody.map((row) =>
                    row.map((cell) => {
                      // Check if the cell has a 'text' property and adjust the fontSize
                      if (typeof cell === "object" && cell.text) {
                        return { ...cell, fontSize: 10 };
                      }
                      // If the cell is a string, wrap it in an object with the desired fontSize
                      return { text: cell, fontSize: 10 };
                    }),
                  ),
                },
                margin: [0, 0, 0, 20], // Adjust the margin for the table
              },
            ],

            pageSize: "A4",
            pageMargins: [20, 20, 20, 20],
            pageOrientation: "landscape",
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

    resetDetailItem() {
      this.detailItem = null;
      this.detailModalVisible = false;
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
        >Create Alur Tujuan Pembelajaran</va-button
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
      title="Form Input Alur Tujuan Pembelajaran"
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
      title="Form Edit Alur Tujuan Pembelajaran"
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
