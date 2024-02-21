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
  Pertanyaan: "",
};

const displayNames = {
  User: "Nama Penyusun",
  Mapel: "Mata Pelajaran",
  Kelas: "Kelas",
  TahunAjar: "Tahun Ajar",
  questionCount: "Total Soal",
  Pertanyaan: "Pertanyaan",
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
      dynamicFields: {},
      textAreaFields: ["Pertanyaan"],
      input: ref({ value: "" }),
      items: [],
      usersOptions: [],
      mapelsOptions: [],
      kelasOptions: [],
      tahunAjarOptions: [],
      bankSoalOptions: [],
      // textAreaFields: ["Pertanyaan"],
      showModal: false,
      viewModalVisible: false,
      detailItem: null,
      detailModalVisible: false,
      displayNames,
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
      return (
        this.createdItem.UserID &&
        this.createdItem.MapelID &&
        this.createdItem.KelasID &&
        this.createdItem.TahunAjarID &&
        this.createdItem.Pertanyaan
      );
    },

    inputFields() {
      return Object.keys(this.createdItem);
    },
  },

  methods: {
    /**
     * Fetches data from the API and populates component state.
     *
     * Makes requests to multiple API endpoints to get users, classes,
     * subjects, academic years, and formative assessment data.
     *
     * Merges question count data into the formative assessment items.
     *
     * Handles loading state and errors.
     */
    async fetchData() {
      this.loading = true;

      try {
        // Simulate a delay
        await new Promise((resolve) => setTimeout(resolve, 2000));

        const response = await axios.get("http://localhost:3000/api/formatif");
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

        const questionCountResponse = await axios.get(
          "http://localhost:3000/api/total/formatif",
        );

        console.log(
          "Question count response data:",
          questionCountResponse.data,
        );

        console.log("Server response:", response);

        // Populate usersOptions, mapelsOptions, kelasOptions, tahunAjarOptions
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
            Pertanyaan: item?.Pertanyaan || "",
            questionCount: questionCount, // Set the question count
          };
        });
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    /**
     * Adds a new item to the list of items.
     *
     * Collects the static and dynamic fields from the form data, constructs a newItem object,
     * logs it, makes a POST request to the API to add it, pushes the new item into the local list,
     * refetches the data to refresh the table, and resets the form.
     *
     * Handles errors from the API request.
     */
    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      // Collect static fields
      const dynamicFields = {
        Pertanyaan: this.createdItem.Pertanyaan,
      };

      // Collect dynamic fields
      this.textAreaFields.forEach((fieldKey) => {
        dynamicFields[fieldKey] = this.createdItem[fieldKey];
      });

      // Construct the newItem object without the questionCount field
      const newItem = {
        ...this.createdItem,
        DynamicFields: dynamicFields,
      };

      // Explicitly remove the questionCount field from the newItem object
      delete newItem.questionCount;

      // Remove the individual dynamic fields from the main object
      this.textAreaFields.forEach((fieldKey) => {
        delete newItem[fieldKey];
      });

      // Log the newItem before sending the request
      console.log("New Item:", newItem);

      try {
        await axios.post("http://localhost:3000/api/formatif", newItem);
        this.items.push({
          ...newItem,
        });

        // Re-fetch the data to refresh the table
        await this.fetchData();

        setTimeout(() => {
          this.resetCreatedItem();
        }, 500);
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },

    /**
     * Edits an existing item.
     *
     * Collects the dynamic fields from the form.
     * Creates an item to edit with the dynamic fields included.
     * Sends a PUT request to update the item on the server.
     * Refreshes the data table after updating.
     */
    async editItem() {
      // Collect dynamic fields from the form
      const dynamicFields = this.collectDynamicFields();

      // Create the item to edit with dynamic fields included
      const itemToEdit = {
        ...this.editedItem,
        DynamicFields: dynamicFields, // Include dynamic fields in the object
      };

      try {
        // Create a deep copy of the edited item
        const editedData = JSON.parse(JSON.stringify(itemToEdit));

        // Change 'ID' to 'id'
        editedData.id = editedData.ID;
        delete editedData.ID;
        delete editedData.User;
        delete editedData.Mapel;
        delete editedData.Kelas;
        delete editedData.TahunAjar;

        // Send the PUT request with the edited data
        await axios.put(
          `http://localhost:3000/api/formatif/${this.editedItem.id}`,
          editedData,
        );

        this.resetEditedItem();
        // Re-fetch the data to refresh the table
        await this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    /**
     * Deletes an item by ID.
     *
     * Prompts the user to confirm deletion.
     * Sends a DELETE request to the API to delete the item.
     * Removes the deleted item from the local items array.
     * Refetches the data to refresh the table.
     */
    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          const response = await axios.delete(
            `http://localhost:3000/api/formatif/${id}`,
          );

          // Check if deletedFormatif is not null
          if (response.data && response.data.data) {
            const deletedFormatif = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedFormatif.id,
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

    /**
     * Opens the detail modal to show more info about the item at the given row index.
     * Fetches the full item data from the API and populates the modal with it.
     */
    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/formatif/${selectedItemId}`,
        );
        const data = response.data.data;
        if (data) {
          console.log("Retrieved data:", data);

          // Use only the DynamicFields object
          this.detailItem = {
            DynamicFields: data.DynamicFields || {},
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

    /**
     * Prints a row from the table by fetching the full data for that row from the API.
     * Extracts relevant metadata and question data, and generates a PDF document containing the metadata and questions.
     *
     * @param {number} rowIndex - The index of the row in the filtered items array to print.
     */
    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log(`Selected item ID: ${selectedItemId}`);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/formatif/${selectedItemId}`,
        );
        const data = response.data.data;

        // Log the entire data object to inspect its structure
        console.log("Full data object:", data);

        if (data) {
          const metadata = {
            Mapel: data.Mapel.Mapel, // Access the 'Mapel' property within the nested object
            Kelas: data.Kelas.Kelas, // Access the 'Kelas' property within the nested object
          };
          console.log("Extracted metadata:", metadata);

          // Extract the questions from DynamicFields
          const questions = Object.values(data.DynamicFields);
          console.log("Extracted questions:", questions);

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
                text: "",
                fontSize: 14,
                bold: true,
                margin: [0, 20, 0, 20],
              },
              { text: `Mapel: ${metadata.Mapel}`, fontSize: 10 },
              { text: `Kelas: ${metadata.Kelas}`, fontSize: 10 },
              {
                canvas: [
                  {
                    type: "line",
                    x1: 0,
                    y1: 0,
                    x2: 510, // Adjust this value to match the width of your page
                    y2: 0,
                    lineWidth: 2,
                    color: "#000000", // Change the color as needed
                  },
                ],
                margin: [0, 10, 0, 10], // Adjust the margin as needed
              },
              ...questions.map((question, index) => ({
                text: `${index + 1}. ${question}`,
                fontSize: 10,
                margin: [0, 5, 0, 5],
              })),
            ],
            pageSize: "A4",
            pageMargins: [20, 20, 20, 20],
            pageOrientation: "portrait",
          };

          console.log("Doc definition:", docDefinition);

          const pdf = pdfMake.createPdf(docDefinition);
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
      }));
    },

    collectDynamicFields() {
      // Collect dynamic fields from textAreaFields and createdItem
      const dynamicFields = {};
      this.textAreaFields.forEach((fieldKey) => {
        dynamicFields[fieldKey] = this.createdItem[fieldKey];
      });
      return dynamicFields;
    },

    addField() {
      // Determine the next number suffix based on the current length of textAreaFields
      const nextNumber = this.textAreaFields.length + 1;
      const newKey = `pertanyaan ${nextNumber}`; // Generate a unique key with the number suffix

      // Add the new key to the textAreaFields array
      this.textAreaFields.push(newKey);

      // Use Vue's $set to reactively add the new field to createdItem with an empty string as the value
      this.$set(this.createdItem, newKey, "");
    },

    resetEditedItem() {
      this.editedItem = null;
      this.editedItemId = null;
    },

    resetCreatedItem() {
      // Reset the createdItem to the default values
      this.createdItem = { ...defaultItem };

      // Reset the dynamicFields to an empty object
      this.dynamicFields = {};

      // Preserve the Pertanyaan field in textAreaFields
      this.textAreaFields = ["Pertanyaan"];

      // Close the modal
      this.showModal = false;

      // Optionally, reset any form validation states
      // For example, if using Vuelidate:
      // this.$v.$reset();
    },

    resetAddState() {
      // Reset the createdItem to the default values
      this.createdItem = { ...defaultItem };

      // Reset the dynamicFields to an empty object
      this.dynamicFields = {};

      // Reset the textAreaFields array to its initial state
      this.textAreaFields = ["Pertanyaan"];

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

        // Log the editedItem to check the Pertanyaan value
        console.log("Edited item:", this.editedItem);

        // Populate textAreaFields with the keys of the dynamic fields
        this.textAreaFields = Object.keys(this.editedItem.DynamicFields || {});

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
        >Create Asesmen Formatif</va-button
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
            @click="openModalToEditItemById(filteredItems[rowIndex].ID)"
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
      title="Form Input Asesmen Formatf"
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

      <va-textarea
        v-for="(fieldKey, index) in textAreaFields"
        :key="index"
        :label="displayNames[fieldKey] || fieldKey"
        v-model="createdItem[fieldKey]"
        class="my-6"
      />

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

      <!-- Static 'pertanyaan' field -->
      <va-textarea
        :label="displayNames.Pertanyaanertanyaan || 'Pertanyaan'"
        v-model="editedItem.Pertanyaan"
        class="my-6"
      />

      <!-- Dynamic fields area -->
      <div v-for="(fieldKey, index) in textAreaFields" :key="index">
        <va-textarea
          :label="displayNames[fieldKey] || fieldKey"
          v-model="editedItem.DynamicFields[fieldKey]"
          class="my-6"
        />
      </div>

      <!-- Button to add more dynamic fields -->
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
      >
        Add Fields
      </va-button>
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      stripe
      title="Detail Asesmen Formatif"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <!-- Display the Pertanyaan field
      <va-textarea
        :label="'Pertanyaan'"
        v-model="detailItem.Pertanyaan"
        class="my-6"
        readonly
      /> -->

      <!-- Display dynamic fields -->
      <va-textarea
        v-for="(value, key) in detailItem.DynamicFields"
        :key="key"
        :label="key"
        v-model="detailItem.DynamicFields[key]"
        class="my-6"
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
