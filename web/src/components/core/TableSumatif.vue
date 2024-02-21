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
    /**
 * Fetches sumatif data from the API and processes it:
 * - Makes requests to the API endpoints to get the sumatif data and related entities. 
 * - Extracts options for select inputs from the entity data.
 * - Gets question counts for each sumatif and merges into the data.
 * - Transforms the raw API data into a cleaned up array for the component.
 * - Catches any errors and sets the loading state.
*/
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

    /**
 * Adds a new item to the list of sumatif items. 
 * 
 * Constructs a payload object containing the new item data, including static fields like 
 * BankSoalID and dynamic fields from the dynamicFieldsArray. Sends the payload to the API 
 * to add the new item. Updates the local items array and resets the form.
 */
    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        // Start with the static field
        let fieldsPayload = [
          {
            value: this.createdItem.BankSoalID,
            label: this.bankSoalOptions.find(
              (opt) => opt.value === this.createdItem.BankSoalID,
            ).label,
          },
        ];

        // Then add the dynamic fields
        const dynamicFieldsArray = this.dynamicFieldsArray
          .map((value) => {
            const selectedOption = this.bankSoalOptions.find(
              (option) => option.value === value,
            );
            return selectedOption
              ? { value: selectedOption.value, label: selectedOption.label }
              : null;
          })
          .filter(Boolean); // Filters out null values

        // Concatenate the static and dynamic fields
        fieldsPayload = fieldsPayload.concat(dynamicFieldsArray);

        // Include the fields in the payload
        const payload = {
          ...this.createdItem,
          DynamicFields: fieldsPayload,
        };

        // Remove the questionCount field from the payload
        delete payload.questionCount;

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

    /**
 * Edits an existing item in the data table.
 * 
 * Makes a deep copy of the edited item object. 
 * Renames the 'ID' property to 'id'.
 * Removes unwanted properties from the copied object.
 * Sends a PUT request to update the item on the server.
 * Resets the editedItem data after update.
 * Refetches data to refresh the table after update.
 */
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

    /**
 * Deletes an item by ID.
 * 
 * Prompts user to confirm deletion.
 * Sends DELETE request to server.
 * Removes item from items array.
 * Refetches data to refresh table.
 * Shows success message.
 * Handles errors.
 */
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

    /**
 * Opens the detail modal for the row at the given index.
 * 
 * Gets the item ID from the filtered items array.
 * Makes API call to get full item data.
 * Maps response data to detailItem object.
 * Shows modal with detailItem data.
 * Handles errors.
 */
    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/sumatif/${selectedItemId}`,
        );
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
          console.error(
            "No data received from the server or BankSoal is undefined",
          );
        }
      } catch (error) {
        console.error("Error fetching data for the detail modal:", error);
      }
    },

    /**
 * Prints a row from the table by fetching the data for the given row index.
 * 
 * It makes API calls to fetch the full data for the row, extracts relevant metadata, 
 * fetches additional data from another API endpoint, extracts questions and options,
 * generates a PDF definition, and opens the PDF.
 * 
 * The PDF contains the metadata, questions and options extracted from the data.
 * 
 * @param {number} rowIndex - The index of the row to print
 */
    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log(`Selected item ID: ${selectedItemId}`);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/sumatif/${selectedItemId}`,
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

          // Fetch the banksoal data to get the actual options
          const bankSoalResponse = await axios.get(
            "http://localhost:3000/api/bank",
          );
          const bankSoalData = bankSoalResponse.data.data;
          console.log("BankSoal data:", bankSoalData);

          // Extract the questions and their corresponding options from DynamicFields
          const questionsAndOptions = [];
          data.DynamicFields.forEach((dynamicField) => {
            // Use the 'value' property as the optionId
            const optionId = dynamicField.value;
            // Find the corresponding question in the bankSoalData using the optionId
            const questionData = bankSoalData.find(
              (bankSoal) => bankSoal.ID === optionId,
            );
            if (questionData) {
              // Push the question and its options to the questionsAndOptions array
              questionsAndOptions.push({
                question: dynamicField.label,
                options: [
                  questionData.OptionA,
                  questionData.OptionB,
                  questionData.OptionC,
                  questionData.OptionD,
                  questionData.OptionE,
                ],
              });
            }
          });
          console.log("Extracted questions and options:", questionsAndOptions);

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
              // {
              //   text: "Questions and Options",
              //   fontSize: 14,
              //   bold: true,
              //   margin: [0, 20, 0, 20],
              // },
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
              ...questionsAndOptions.flatMap((item, index) => [
                {
                  text: `${index + 1}. ${item.question}`,
                  fontSize: 10,
                  margin: [0, 5, 0, 5],
                },
                ...item.options.map((option, optionIndex) => ({
                  text: `${String.fromCharCode(97 + optionIndex)}). ${option}`,
                  fontSize: 10,
                  margin: [0, 5, 0, 5],
                })),
              ]),
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
        // Update the label and value of the object at the given index
        this.$set(this.dynamicFieldsArray, index, {
          label: selectedOption.label,
          value: selectedOption.value,
        });
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
        >Create Asesmen Sumatif</va-button
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
      title="Form Input Asesmen Sumatif"
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
