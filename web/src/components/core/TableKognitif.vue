<script>
import { defineComponent } from "vue";
import { ref } from "vue";
import pdfMake from "pdfmake/build/pdfmake";
import pdfFonts from "pdfmake/build/vfs_fonts";
import axios from "axios";

pdfMake.vfs = pdfFonts.pdfMake.vfs;

const defaultItem = {
  User: {}, // Initialize as an empty object
  Mapel: {}, // Initialize as an empty object
  Kelas: {}, // Initialize as an empty object
  TahunAjar: {}, // Initialize as an empty object
  BankSoal: {
    Soal: "",
    OptionA: "",
    OptionB: "",
    OptionC: "",
    OptionD: "",
    OptionE: "",
  },
};

const displayNames = {
  User: "Nama Penyusun",
  Mapel: "Mata Pelajaran",
  Kelas: "Kelas",
  TahunAjar: "Tahun Ajar",
  BankSoal: "Soal",
  OptionA: "Pilihan A",
  OptionB: "Pilihan B",
  OptionC: "Pilihan C",
  OptionD: "Pilihan D",
  OptionE: "Pilihan E",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "User", label: "Nama Penyusun", sortable: false },
      { key: "Mapel", label: "Mata Pelajaran", sortable: false },
      { key: "Kelas", label: "Kelas", sortable: false },
      { key: "TahunAjar", label: "Tahun Ajar", sortable: false },
      { key: "Soal", label: "Soal", sortable: false },
      { key: "actions", label: "Actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: {},
      dynamicFieldsArray: [],
      input: ref({ value: "" }),
      items: [],
      usersOptions: [],
      mapelsOptions: [],
      kelasOptions: [],
      tahunAjarOptions: [],
      bankSoalOptions: [],
      showModal: false,
      viewModalVisible: false,
      displayNames,
      detailItem: {},
      detailModalVisible: false,
      loading: false,
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
      // Function to check if a value is truthy, allowing empty strings
      const isTruthyOrEmptyString = (value) => {
        if (typeof value === "object" && value !== null) {
          return Object.values(value).every(isTruthyOrEmptyString);
        }
        return value === "" || !!value;
      };

      // Check if BankSoal exists and has truthy or empty string Soal and Options
      const bankSoalExistsAndHasValues =
        this.createdItem.BankSoal &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.Soal) &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.OptionA) &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.OptionB) &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.OptionC) &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.OptionD) &&
        isTruthyOrEmptyString(this.createdItem.BankSoal.OptionE);

      // Check if all other keys in createdItem have truthy or empty string values
      const allOtherKeysTruthyOrEmptyString = Object.keys(this.createdItem)
        .filter((key) => key !== "BankSoal") // Exclude BankSoal from this check
        .every((key) => isTruthyOrEmptyString(this.createdItem[key]));

      // Return true only if BankSoal has values and all other keys have truthy or empty string values
      return bankSoalExistsAndHasValues && allOtherKeysTruthyOrEmptyString;
    },

    inputFields() {
      return Object.keys(this.createdItem);
    },
  },

  methods: {
    async fetchData() {
      this.loading = true;

      try {
        // Simulate a delay
        await new Promise((resolve) => setTimeout(resolve, 2000));

        const response = await axios.get("http://localhost:3000/api/kognitif");
        const userResponse = await axios.get(
          "http://localhost:3000/api/public/user",
        );
        const kelasResponse = await axios.get(
          "http://localhost:3000/api/public/kelas",
        );
        const mapelResponse = await axios.get(
          "http://localhost:3000/api/public/mapel",
        );
        const tahunResponse = await axios.get(
          "http://localhost:3000/api/public/tahun",
        );
        const bankResponse = await axios.get("http://localhost:3000/api/bank");

        this.usersOptions = this.extractOptions(userResponse.data.data, "Name");

        this.kelasOptions = this.extractOptions(
          kelasResponse.data.data,
          "Kelas",
        );

        this.mapelsOptions = this.extractOptions(
          mapelResponse.data.data,
          "Mapel",
        );

        this.tahunAjarOptions = this.extractOptions(
          tahunResponse.data.data,
          "Tahun",
        );

        this.bankSoalOptions = this.extractOptions(
          bankResponse.data.data,
          "Soal",
        );

        console.log("bankSoal options:", this.bankSoalOptions);

        this.items = response.data.data.map((item) => ({
          ...item,
          ID: item?.ID || "", // Use 'ID' instead of 'id'
          User: item?.User.Name || "",
          Mapel: item?.Mapel.Mapel || "",
          Kelas: item?.Kelas.Kelas || "",
          TahunAjar: item?.TahunAjar.Tahun || "",
          Soal: item?.BankSoal.Soal || "",
          OptionA: item?.BankSoal.OptionA || "",
          OptionB: item?.BankSoal.OptionB || "",
          OptionC: item?.BankSoal.OptionC || "",
          OptionD: item?.BankSoal.OptionD || "",
          OptionE: item?.BankSoal.OptionE || "",
        }));

        console.log("Kognitif items:", this.items);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    async addNewItem() {
      // Log the initial state of isNewData
      console.log("Initial isNewData:", this.isNewData);

      // Log the contents of createdItem and dynamicFieldsArray
      console.log("Contents of createdItem:", this.createdItem);
      console.log("Contents of dynamicFieldsArray:", this.dynamicFieldsArray);

      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        console.log("Creating new item with data:", this.createdItem);

        // Include the dynamic fields in the payload
        const payload = {
          ...this.createdItem,
          DynamicFields: this.collectDynamicFields(),
        };

        // Log the payload before sending it to the server
        console.log("Payload to send:", payload);

        const response = await axios.post(
          "http://localhost:3000/api/kognitif",
          payload,
        );

        this.items.push({
          ...payload,
        });

        console.log("Server Response:", response.data);

        // Re-fetch the data to refresh the table
        await this.fetchData();

        setTimeout(() => {
          this.resetCreatedItem();
        }, 500);
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },

    async editItem() {
      try {
        // Create a deep copy of the edited item
        const editedData = JSON.parse(JSON.stringify(this.editedItem));

        // Change 'ID' to 'id'
        editedData.id = editedData.ID;
        delete editedData.ID;
        delete editedData.User;
        delete editedData.Mapel;
        delete editedData.Kelas;
        delete editedData.TahunAjar;
        delete editedData.BankSoal;

        await axios.put(
          `http://localhost:3000/api/kognitif/${editedData.id}`,
          ...editedData,
        );

        this.resetEditedItem();
        // Re-fetch the data to refresh the table
        await this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          const response = await axios.delete(
            `http://localhost:3000/api/kognitif/${id}`,
          );

          // Check if deletedKognitif is not null
          if (response.data && response.data.data) {
            const deletedKognitif = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedKognitif.id,
            );
          }

          // Re-fetch the data to refresh the table
          await this.fetchData();

          // Optionally, you can show a success message
          alert("Item deleted successfully");
        } catch (error) {
          console.error("Error deleting item:", error);
        }
      }
    },

    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/kognitif/${selectedItemId}`,
        );
        const data = response.data.data;
        if (data) {
          console.log("Data:", data);

          this.detailItem = {
            BankSoalID: data.BankSoalID || "",
            Soal: data.BankSoal.Soal || "",
            OptionA: data.BankSoal.OptionA || "",
            OptionB: data.BankSoal.OptionB || "",
            OptionC: data.BankSoal.OptionC || "",
            OptionD: data.BankSoal.OptionD || "",
            OptionE: data.BankSoal.OptionE || "",
          };

          console.log("Detail item:", this.detailItem);
          this.detailModalVisible = true;
        } else {
          console.error("No data received from the server");
        }
      } catch (error) {
        console.error("Error fetching data for the detail modal:", error);
      }
    },

    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;

      try {
        // Fetch the necessary data directly from the server
        const response = await axios.get(
          `http://localhost:3000/api/kognitif/${selectedItemId}`,
        );
        const data = response.data.data;

        // Log all properties of the data object
        console.log("Data properties:", Object.keys(data));

        if (data) {
          console.log("Printing row with ID:", selectedItemId);

          const tableBody = [
            [
              { text: "Judul Capaian", fontSize: 10, bold: true },
              { text: "Judul Elemen", fontSize: 10, bold: true },
              { text: "Keterangan Elemen", fontSize: 10, bold: true },
              { text: "Keterangan Proses Mengamati", fontSize: 10, bold: true },
              {
                text: "Keterangan Proses Mempertanyakan",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Keterangan Proses Merencanakan",
                fontSize: 10,
                bold: true,
              },
              { text: "Keterangan Proses Memproses", fontSize: 10, bold: true },
              {
                text: "Keterangan Proses Mengevaluasi",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Keterangan Proses Mengkomunikasikan",
                fontSize: 10,
                bold: true,
              },
            ],
            [
              "judulCapaian" in data ? data.judulCapaian : "N/A",
              "judulElemen" in data ? data.judulElemen : "N/A",
              "ketElemen" in data ? data.ketElemen : "N/A",
              "ketProsesMengamati" in data ? data.ketProsesMengamati : "N/A",
              "ketProsesMempertanyakan" in data
                ? data.ketProsesMempertanyakan
                : "N/A",
              "ketProsesMerencanakan" in data
                ? data.ketProsesMerencanakan
                : "N/A",
              "ketProsesMemproses" in data ? data.ketProsesMemproses : "N/A",
              "ketProsesMengevaluasi" in data
                ? data.ketProsesMengevaluasi
                : "N/A",
              "ketProsesMengkomunikasikan" in data
                ? data.ketProsesMengkomunikasikan
                : "N/A",
            ],
          ];

          console.log("Table Body:", tableBody);

          const docDefinition = {
            footer: function (currentPage, pageCount) {
              return [
                {
                  text: currentPage.toString() + " of " + pageCount,
                  alignment: "center",
                  fontSize: 8,
                  margin: [10, 10, 10, 0],
                },
              ];
            },
            content: [
              {
                text: "Capaian Pembelajaran",
                fontSize: 12,
                bold: true,
                alignment: "center",
                margin: [0, 20, 0, 20],
              },
              {
                table: {
                  headerRows: 1,
                  widths: Array(tableBody[0].length).fill("auto"),
                  body: tableBody,
                },
                margin: [0, 0, 0, 20],
              },
            ],
            pageSize: "A4",
            pageMargins: [20, 20, 20, 20],
            pageOrientation: "landscape",
          };

          const pdf = pdfMake.createPdf(docDefinition);

          // Open the PDF for printing
          pdf.open();
        } else {
          console.error("No data received from the server");
        }
      } catch (error) {
        console.error("Error fetching data for printing:", error);
      }
    },

    extractOptions(data, labelProperty) {
      if (!Array.isArray(data)) {
        console.error("Data is not an array:", data);
        return [];
      }

      return data.map((item) => ({
        label: item[labelProperty] ? item[labelProperty] : "", // Use '' if labelProperty is null or undefined
        value: item.ID, // Use 'ID' instead of 'id'
        options: {
          OptionA: item.OptionA,
          OptionB: item.OptionB,
          OptionC: item.OptionC,
          OptionD: item.OptionD,
          OptionE: item.OptionE,
        },
      }));
    },

    collectDynamicFields() {
      // Collect dynamic fields from textAreaFields, createdItem, and dynamicFieldsArray
      const dynamicFields = {};

      this.dynamicFieldsArray.forEach((fieldValue, index) => {
        dynamicFields[`field${index + 1}`] = fieldValue;
      });
      return dynamicFields;
    },

    addField() {
      // Add a new entry to the dynamicFieldsArray
      this.dynamicFieldsArray.push("");

      // Directly assign the new field to createdItem with an empty string as the value
      this.createdItem[newKey] = "";
    },

    resetEditedItem() {
      this.editedItem = null;
      this.editedItemId = null;
    },

    resetCreatedItem() {
      // Reset the createdItem to the default values
      this.createdItem = { ...defaultItem };

      // Reset the dynamicFieldsArray to an empty array
      this.dynamicFieldsArray = [];

      // Close the modal
      this.showModal = false;

      // Optionally, reset any form validation states
      // For example, if using Vuelidate:
      // this.$v.$reset();
    },

    resetAddState() {
      // Reset the createdItem to the default values
      this.createdItem = { ...defaultItem };

      // Reset the dynamicFieldsArray to an empty array
      this.dynamicFieldsArray = [];

      // Optionally, reset any form validation states
      // For example, if using Vuelidate:
      // this.$v.$reset();
    },

    openModalToEditItemById(id) {
      // Find the item by its ID
      const item = this.items.find((item) => item.ID === id);

      // Check if the item exists
      if (item) {
        // Set the editedItem to the found item
        this.editedItem = { ...item };

        // Update createdItem with the dynamic fields from the editedItem
        this.createdItem = { ...item.DynamicFields };

        // Log the editedItem to check the Pertanyaan value
        console.log("Edited item:", this.editedItem);

        // Open the modal
        this.showEditModal = true;
      } else {
        console.error(`Item with ID ${id} not found.`);
      }
    },

    toggleAddModal() {
      this.resetAddState();

      this.showModal = !this.showModal;
      if (!this.showModal) {
        this.resetCreatedItem();
      }
    },

    resetDetailItem() {
      // Reset the detailItem to its initial state
      this.detailItem = {};

      // Reset the filteredDetailFields array to its initial state
      this.filteredDetailFields = [];

      // Close the detail modal
      this.detailModalVisible = false;
    },

    handleSelect(selectedOption) {
      this.createdItem.BankSoalID = selectedOption.value;
    },
  },

  watch: {
    "editedItem.DynamicFields.pertanyaan2"(newValue, oldValue) {
      console.log("pertanyaan2 changed from", oldValue, "to", newValue);
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
    <va-input
      v-model="input.value"
      type="text"
      placeholder="Search..."
    ></va-input>
    <va-button-group
      icon-color="#000000"
      preset="secondary"
      border-color="#000000"
    >
      <va-button @click="toggleAddModal" preset="secondary" icon="add"
        >Add Capaian Pembelajaran</va-button
      >
    </va-button-group>
  </div>
  <div>
    <va-data-table
      :items="filteredItems"
      :columns="columns"
      striped
      :loading="loading"
    >
      <template #cell(actions)="{ rowIndex }">
        <div class="action-buttons">
          <va-button preset="plain" icon="print" @click="printRow(rowIndex)" />
          <va-button
            preset="plain"
            icon="remove_red_eye"
            @click="openDetailModal(rowIndex)"
          />
          <!-- <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          /> -->
          <va-button
            preset="plain"
            icon="delete"
            @click="deleteItemById(filteredItems[rowIndex].ID)"
          />
        </div>
      </template>
    </va-data-table>

    <!-- Modal Content -->
    <va-modal
      blur
      class="modal-crud"
      stripe
      title="Add Asesmen Kognitif"
      size="large"
      :model-value="showModal"
      @ok="addNewItem"
      @cancel="resetCreatedItem"
    >
      <!-- Using va-select for user, mapel, kelas, and tahun ajar -->
      <va-select
        v-model="createdItem.UserID"
        :label="displayNames.User"
        :options="usersOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="createdItem.MapelID"
        :label="displayNames.Mapel"
        :options="mapelsOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="createdItem.KelasID"
        :label="displayNames.Kelas"
        :options="kelasOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="createdItem.TahunAjarID"
        :label="displayNames.TahunAjar"
        :options="tahunAjarOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />

      <!-- this va-select below is intended for soal -->

      <va-select
        v-model="createdItem.BankSoalID"
        :label="displayNames.BankSoal"
        :options="bankSoalOptions"
        class="my-6"
        text-by="label"
        value-by="value"
        @change="handleSelect"
        autocomplete
      />

      <div v-for="(field, index) in dynamicFieldsArray" :key="index">
        <va-select
          v-model="dynamicFieldsArray[index]"
          :label="'Field ' + (index + 1)"
          :options="bankSoalOptions"
          class="my-6"
          text-by="label"
          value-by="value"
          @change="handleSelect"
          autocomplete
        />
      </div>

      <va-button
        class="my-6"
        color="primary"
        @click="addField"
        style="
          width: 100%;
          display: flex;
          box-sizing: border-box;
          margin-bottom: 10px;
        "
        >Add Fields
      </va-button>
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      :model-value="!!editedItem"
      title="Edit Asesmen Kognitif"
      size="large"
      @ok="editItem"
      @cancel="resetEditedItem"
    >
      <!-- Using va-select for user, mapel, kelas, and tahun ajar -->
      <va-select
        v-model="editedItem.UserID"
        :label="displayNames.User"
        :options="usersOptions"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="editedItem.MapelID"
        :label="displayNames.Mapel"
        :options="mapelsOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="editedItem.KelasID"
        :label="displayNames.Kelas"
        :options="kelasOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="editedItem.TahunAjarID"
        :label="displayNames.TahunAjar"
        :options="tahunAjarOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />

      <!-- this va-select below is intended for soal -->

      <va-select
        v-model="editedItem.BankSoalID"
        :label="displayNames.BankSoal"
        :options="bankSoalOptions"
        class="my-6"
        text-by="label"
        value-by="value"
        @change="handleSelect"
        autocomplete
      />
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      stripe
      title="Detail Asesmen Kognitif"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <!-- this va-select below is intended for soal -->
      <va-select
        v-model="detailItem.Soal"
        :label="displayNames.BankSoal"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />

      <!-- this va-select below is intended for options a until e -->
      <va-select
        v-model="detailItem.OptionA"
        :label="displayNames.OptionA"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />
      <va-select
        v-model="detailItem.OptionB"
        :label="displayNames.OptionB"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />
      <va-select
        v-model="detailItem.OptionC"
        :label="displayNames.OptionC"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />
      <va-select
        v-model="detailItem.OptionD"
        :label="displayNames.OptionD"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />
      <va-select
        v-model="detailItem.OptionE"
        :label="displayNames.OptionE"
        class="my-6"
        text-by="label"
        value-by="value"
        readonly
      />
    </va-modal>
  </div>
</template>

<style>
.action-buttons {
  display: flex;
  gap: 8px;
  /* Adjust the gap to your preference */
}

.va-input {
  display: block;
  margin-bottom: 10px;
}

.va-select .dropdown-menu {
  display: block;
}

.modal-crud {
  .va-select {
    display: block;
    margin-bottom: 10px;
  }
  .va-textarea {
    width: 100%;
    display: flex;
    box-sizing: border-box;
    margin-bottom: 20px;
  }
}
</style>
