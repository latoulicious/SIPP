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
  questionCount: "",
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
  questionCount: "Total Soal",
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
      { key: "questionCount", label: "Total Soal", sortable: false },
      { key: "actions", label: "Actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: {},
      dynamicFieldsArray: [],
      selectedValues: {},
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
      detailItem: {
        DynamicFields: [], // Initialize DynamicFields as an empty array
        BankSoal: {
          Soal: "",
          OptionA: "",
          OptionB: "",
          OptionC: "",
          OptionD: "",
          OptionE: "",
        },
      },
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

    dynamicFieldsSelectItems() {
      // Check if DynamicFields is defined and is an array
      if (Array.isArray(this.detailItem.DynamicFields)) {
        return this.detailItem.DynamicFields.map((field) => ({
          text: field.label,
          value: field.value,
        }));
      }
      // Return an empty array if DynamicFields is not defined or not an array
      return [];
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

        const response = await axios.get("http://localhost:3000/api/sumatif");
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

        const questionCountResponse = await axios.get(
          "http://localhost:3000/api/total/sumatif",
        );

        console.log(
          "Question count response data:",
          questionCountResponse.data,
        );

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

        console.log("Sumatif options:", response.data);

        // Create a lookup object for quick access to question counts by ID
        const questionCountLookup = {};
        questionCountResponse.data.data.forEach((item) => {
          questionCountLookup[item.ID] = item.questionCount;
        });

        // Merge the question counts with the base data
        this.items = response.data.data.map((item) => {
          // Retrieve the question count using the item's ID
          const questionCount = questionCountLookup[item.ID] || 0;

          return {
            ...item,
            ID: item?.ID || "",
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
            DynamicFields: item?.BankSoal.DynamicFields || [],
            questionCount: questionCount, // Set the question count
          };
        });
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        console.log("Creating new item with data:", this.createdItem);

        // Before adding DynamicFields to the payload
        console.log("Before adding DynamicFields:", this.createdItem);

        // Transform the bankSoalOptions into an array of objects with 'value' and 'label'
        const dynamicFieldsArray = this.bankSoalOptions.map((option) => ({
          value: option.value,
          label: option.label,
        }));

        // Include the dynamic fields in the payload
        const payload = {
          ...this.createdItem,
          DynamicFields: dynamicFieldsArray,
        };

        // After adding DynamicFields to the payload
        console.log("After adding DynamicFields:", payload);

        // Log the payload before sending it to the server
        console.log("Payload to send:", payload);

        const response = await axios.post(
          "http://localhost:3000/api/sumatif",
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
          `http://localhost:3000/api/sumatif/${editedData.id}`,
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
            `http://localhost:3000/api/sumatif/${id}`,
          );

          // Check if deletedSumatif is not null
          if (response.data && response.data.data) {
            const deletedSumatif = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedSumatif.id,
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
    const response = await axios.get(`http://localhost:3000/api/sumatif/${selectedItemId}`);
    const data = response.data.data; // Access the data object directly
    console.log("Full response data:", data); // Log the full response data

    if (data && data.BankSoal) {
      console.log("Data:", data);

      this.detailItem = {
        ...defaultItem,
        BankSoal: {
          Soal: data.BankSoal.Soal || "",
          OptionA: data.BankSoal.OptionA || "",
          OptionB: data.BankSoal.OptionB || "",
          OptionC: data.BankSoal.OptionC || "",
          OptionD: data.BankSoal.OptionD || "",
          OptionE: data.BankSoal.OptionE || "",
        },
        DynamicFields: data.DynamicFields || [], // Access DynamicFields directly from data
      };

      console.log("Detail item:", this.detailItem);
      this.detailModalVisible = true;
    } else {
      console.error("No data received from the server or BankSoal is undefined");
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
          `http://localhost:3000/api/sumatif/${selectedItemId}`,
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
      // Gather the data from the BankSoal proxy object
      const dynamicFields = {
        Soal: this.createdItem.BankSoal.Soal,
        OptionA: this.createdItem.BankSoal.OptionA,
        OptionB: this.createdItem.BankSoal.OptionB,
        OptionC: this.createdItem.BankSoal.OptionC,
        OptionD: this.createdItem.BankSoal.OptionD,
        OptionE: this.createdItem.BankSoal.OptionE,
      };

      // Filter out any fields that are empty strings
      Object.keys(dynamicFields).forEach((key) => {
        if (dynamicFields[key] === "") {
          delete dynamicFields[key];
        }
      });

      // Return the collected dynamic fields
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

    handleSelect(selectedValue, index) {
      // Find the selected option by its value
      const selectedOption = this.bankSoalOptions.find(
        (option) => option.value === selectedValue,
      );
      if (selectedOption) {
        // Update the Soal property of the object at the given index
        this.$set(this.dynamicFieldsArray, index, {});
      }
    },
  },

  watch: {
    detailItem: {
      handler(newVal) {
        if (newVal && Array.isArray(newVal.DynamicFields)) {
          newVal.DynamicFields.forEach((field, index) => {
            console.log(`Field Label ${index}:`, field.label);
          });
        }
      },
      deep: true, // Watch nested properties inside detailItem
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
      title="Add Asesmen Sumatif"
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
          :label="'Soal ' + (index + 2)"
          :options="bankSoalOptions"
          class="my-6"
          text-by="label"
          value-by="value"
          @change="handleSelect($event, index)"
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
      title="Edit Asesmen Sumatif"
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
      title="Detail Asesmen Sumatif"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <!-- Dynamically generate va-textarea components for DynamicFields -->
      <div
        v-if="detailItem.DynamicFields && detailItem.DynamicFields.length > 0"
      >
        <div v-for="(field, index) in detailItem.DynamicFields" :key="index">
          <va-textarea
            :label="'Soal ' + (index + 1)"
            class="my-6"
            :modelValue="field.label"
            readonly
          />
        </div>
      </div>
      <!-- If there are no DynamicFields, display a placeholder message -->
      <div v-else>
        <p>No dynamic fields found.</p>
      </div>

      <div class="options-container">
        <div class="top-row">
          <!-- Option A -->
          <va-textarea
            :label="displayNames.OptionA"
            class="my-6"
            :value="detailItem.BankSoal.OptionA"
            readonly
          />
          <!-- Option B -->
          <va-textarea
            :label="displayNames.OptionB"
            class="my-6"
            :value="detailItem.BankSoal.OptionB"
            readonly
          />
          <!-- Option C -->
          <va-textarea
            :label="displayNames.OptionC"
            class="my-6"
            :value="detailItem.BankSoal.OptionC"
            readonly
          />
        </div>
        <div class="bottom-row">
          <!-- Option D -->
          <va-textarea
            :label="displayNames.OptionD"
            class="my-6"
            :value="detailItem.BankSoal.OptionD"
            readonly
          />
          <!-- Option E -->
          <va-textarea
            :label="displayNames.OptionE"
            class="my-6"
            :value="detailItem.BankSoal.OptionE"
            readonly
          />
        </div>
      </div>
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

<style scoped>
.options-container {
  display: flex;
  flex-direction: column; /* Stack top-row and bottom-row vertically */
  gap: 10px; /* Match the margin-bottom of the non-scoped va-select */
}

.top-row,
.bottom-row {
  display: flex;
  justify-content: space-between; /* Distribute space evenly between children */
  gap: 10px; /* Match the margin-bottom of the non-scoped va-select */
}

.option-select {
  flex: 1; /* Allow each select to grow and shrink equally */
  max-width: calc(20% - 10px); /* Limit the width to 20% minus the gap size */
}
</style>
