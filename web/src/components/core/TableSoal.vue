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
  Jurusan: {}, // Initialize as an empty object
  questionCount: "",
  BankSoal: {
    Soal: "",
    OptionA: "",
    OptionB: "",
    OptionC: "",
    OptionD: "",
    OptionE: "",
  },
  TipeSoal: "",
  Hari: "",
  Tanggal: "",
  Waktu: "",
};

const displayNames = {
  User: "Nama Penyusun",
  Mapel: "Mata Pelajaran",
  Kelas: "Kelas",
  Jurusan: "Jurusan",
  questionCount: "Total Soal",
  BankSoal: "Soal",
  OptionA: "Pilihan A",
  OptionB: "Pilihan B",
  OptionC: "Pilihan C",
  OptionD: "Pilihan D",
  OptionE: "Pilihan E",
  TipeSoal: "Tipe Soal",
  Hari: "Hari",
  Tanggal: "Tanggal",
  Waktu: "Waktu",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "User", label: "Nama Penyusun", sortable: false },
      { key: "Mapel", label: "Mata Pelajaran", sortable: false },
      { key: "Kelas", label: "Kelas", sortable: false },
      { key: "TipeSoal", label: "Tipe Soal", sortable: false },
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
      jurusanOptions: [],
      bankSoalOptions: [],
      tipeSoalOptions: [
        { label: "Ulangan Harian", value: "Ulangan Harian" },
        { label: "Ulangan Tengah Semester", value: "Ulangan Tengah Semester" },
        { label: "Ulangan Akhir Semester", value: "Ulangan Akhir Semester" },
        // ... add more options as needed
      ],
      textAreaFields: ["Hari", "Tanggal", "Waktu"],
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

    dynamicFields() {
      return this.detailItem.DynamicFields;
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
     * Fetches data from API endpoints and processes response to populate table.
     *
     * Makes requests to API endpoints to get soal, user, kelas, jurusan, mapel, bank soal,
     * and question count data.
     *
     * Maps question count and dynamic fields to soal items after fetching all data.
     *
     * Returns processed list of items to display in table.
     */
    async fetchData() {
      this.loading = true;

      try {
        // Simulate a delay
        await new Promise((resolve) => setTimeout(resolve, 2000));

        const response = await axios.get("http://localhost:3000/api/soal");
        const userResponse = await axios.get(
          "http://localhost:3000/api/public/user",
        );
        const kelasResponse = await axios.get(
          "http://localhost:3000/api/public/kelas",
        );
        const jurusanResponse = await axios.get(
          "http://localhost:3000/api/public/jurusan",
        );
        const mapelResponse = await axios.get(
          "http://localhost:3000/api/public/mapel",
        );
        const bankResponse = await axios.get("http://localhost:3000/api/bank");
        const questionCountResponse = await axios.get(
          "http://localhost:3000/api/total/item",
        );

        console.log("Soal response data:", response.data);
        console.log(
          "ItemSoal question counts response data:",
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
        this.jurusanOptions = this.extractOptions(
          jurusanResponse.data.data,
          "Jurusan",
        );
        this.bankSoalOptions = this.extractOptions(
          bankResponse.data.data,
          "Soal",
        );

        // Populate questionCountLookup using SoalID
        const questionCountLookup = {};
        questionCountResponse.data.data.forEach((item) => {
          questionCountLookup[item.SoalID] = item.questionCount;
        });

        // Assuming that the first ItemSoal in the Items array contains the DynamicFields
        const dynamicFieldsLookup = {};
        questionCountResponse.data.data.forEach((item) => {
          dynamicFieldsLookup[item.SoalID] = item.DynamicFields;
        });

        // Log the populated dynamicFieldsLookup
        console.log("Populated dynamicFieldsLookup:", dynamicFieldsLookup);

        // Log the populated questionCountLookup
        console.log("Populated questionCountLookup:", questionCountLookup);

        // Map the question counts to the corresponding items
        this.items = response.data.data.map((item) => {
          const soalId = item.ID; // Use the ID field as the SoalID
          const questionCount = questionCountLookup[soalId] || 0;
          const dynamicFields = dynamicFieldsLookup[soalId] || {};

          console.log("Dynamic Fields Map:", dynamicFields);

          console.log(
            "After mapping for Soal ID:",
            soalId,
            "questionCount:",
            questionCount,
          );

          console.log(
            "After mapping for Soal ID:",
            soalId,
            "Dynamic Fields:",
            dynamicFields,
          );

          return {
            ...item,
            ID: item.ID || "",
            User: item.User ? item.User.Name : "",
            Mapel: item.Mapel ? item.Mapel.Mapel : "",
            Kelas: item.Kelas ? item.Kelas.Kelas : "",
            Jurusan: item.Jurusan ? item.Jurusan.Jurusan : "",
            Soal: item.BankSoal ? item.BankSoal.Soal : "",
            OptionA: item.BankSoal ? item.BankSoal.OptionA : "",
            OptionB: item.BankSoal ? item.BankSoal.OptionB : "",
            OptionC: item.BankSoal ? item.BankSoal.OptionC : "",
            OptionD: item.BankSoal ? item.BankSoal.OptionD : "",
            OptionE: item.BankSoal ? item.BankSoal.OptionE : "",
            DynamicFields: dynamicFields, // Correctly map DynamicFields from the first ItemSoal
            questionCount: questionCount,
            TipeSoal: item.TipeSoal || "",
            Hari: item.Hari || "",
            Tanggal: item.Tanggal || "",
            Waktu: item.Waktu || "",
          };
        });

        console.log("Final items list:", this.items);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    /**
     * Adds a new item to the list of items.
     *
     * Makes a POST request to the API to add the new item.
     * Constructs the payload from the static BankSoalID field and
     * dynamic fields selected by the user.
     *
     * Updates the local items array and table after successful
     * addition.
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
          TipeSoal: this.createdItem.TipeSoal, // Ensure TipeSoal is included in the payload
          DynamicFields: fieldsPayload,
        };

        // Remove the questionCount field from the payload
        delete payload.questionCount;

        const response = await axios.post(
          "http://localhost:3000/api/soal",
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
     * Edits an existing item in the list.
     *
     * Makes a deep copy of the edited item to avoid mutating the original.
     * Converts the 'ID' field to 'id' and removes unneeded fields.
     * Sends a PUT request to update the item on the server.
     * Refetches the data to refresh the table after updating.
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
          `http://localhost:3000/api/soal/${editedData.id}`,
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
     * Deletes a soal item by ID.
     *
     * Prompts user to confirm deletion.
     * Makes API call to delete soal item.
     * Removes deleted item from items array.
     * Refreshes data table after deletion.
     * Shows alert on success.
     * Logs error on failure.
     */
    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          const response = await axios.delete(
            `http://localhost:3000/api/soal/${id}`,
          );

          // Check if deletedSoal is not null
          if (response.data && response.data.data) {
            const deletedSoal = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedSoal.id,
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
     * Fetches the data for that row from the API and populates the modal with it.
     */
    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/soal/${selectedItemId}`,
        );
        console.log("Response status:", response.status);
        console.log("Response headers:", response.headers);
        console.log("Response data:", response.data);

        const data = response.data.data; // Access the data directly
        console.log("Data:", data);

        if (data) {
          console.log("Data:", data);

          this.detailItem = {
            ...defaultItem,
            DynamicFields:
              data.Items && data.Items[0] ? data.Items[0].DynamicFields : {}, // Access DynamicFields from the first ItemSoal
          };

          console.log("Detail item:", this.detailItem);
          this.detailModalVisible = true;
        } else {
          console.error(
            "No data received from the server or data is undefined",
          );
        }
      } catch (error) {
        console.error("Error fetching data for the detail modal:", error);
      }
    },

    /**
     * Prints a PDF report for the row at the given index.
     *
     * Fetches the data for the row from the API and extracts the metadata and questions/options.
     * Uses pdfMake to generate a PDF report with the metadata, questions, and options formatted.
     * Opens the generated PDF in the browser for previewing/printing.
     */
    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log(`Selected item ID: ${selectedItemId}`);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/soal/${selectedItemId}`,
        );
        const data = response.data.data;
        console.log("Data from server:", data);

        if (data) {
          const metadata = {
            TipeSoal: data.TipeSoal,
            Mapel: data.Mapel.Mapel, // Access the 'Mapel' property within the nested object
            Kelas: data.Kelas.Kelas, // Access the 'Kelas' property within the nested object
            Jurusan: data.Jurusan.Jurusan, // Access the 'Jurusan' property within the nested object
            Tanggal: data.Tanggal,
            Hari: data.Hari,
            Waktu: data.Waktu,
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
          data.Items.forEach((item) => {
            if (item.DynamicFields) {
              Object.entries(item.DynamicFields).forEach(
                ([question, optionId]) => {
                  // Find the corresponding question in the bankSoalData using the optionId
                  const questionData = bankSoalData.find(
                    (bankSoal) => bankSoal.ID === optionId,
                  );
                  if (questionData) {
                    // Push the question and its options to the questionsAndOptions array
                    questionsAndOptions.push({
                      question: question,
                      options: [
                        questionData.OptionA,
                        questionData.OptionB,
                        questionData.OptionC,
                        questionData.OptionD,
                        questionData.OptionE,
                      ],
                    });
                  }
                },
              );
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
                text: metadata.TipeSoal, // Use data.TipeSoal instead of the static text
                fontSize: 14,
                bold: true,
                margin: [0, 20, 0, 20],
              },
              { text: `Mapel: ${metadata.Mapel}`, fontSize: 10 },
              { text: `Kelas: ${metadata.Kelas}`, fontSize: 10 },
              { text: `Jurusan: ${metadata.Jurusan}`, fontSize: 10 },
              { text: `Tanggal: ${metadata.Tanggal}`, fontSize: 10 },
              { text: `Hari: ${metadata.Hari}`, fontSize: 10 },
              { text: `Waktu: ${metadata.Waktu}`, fontSize: 10 },
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
        >Create Soal</va-button
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
      title="Form Input Soal"
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
        v-model="createdItem.JurusanID"
        :label="displayNames.Jurusan"
        :options="jurusanOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-select
        v-model="createdItem.TipeSoal"
        :label="displayNames.TipeSoal"
        :options="tipeSoalOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />
      <va-textarea
        v-for="key in textAreaFields"
        :key="key"
        :label="displayNames[key]"
        v-model="createdItem[key]"
        class="my-6"
      />

      <!-- this va-select below is intended for soal -->

      <va-select
        v-model="createdItem.BankSoalID"
        :label="displayNames.BankSoal"
        :options="bankSoalOptions"
        class="my-6"
        text-by="label"
        value-by="value"
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
      title="Edit Soal"
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
      title="Detail Soal"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <div
        v-for="(value, key, index) in detailItem.DynamicFields"
        :key="key"
        class="textarea-container"
      >
        <label :for="'textarea-' + index" class="textarea-label"
          >Soal {{ index + 1 }}</label
        >
        <textarea
          :id="'textarea-' + index"
          class="detailarea"
          readonly
          :value="key"
        ></textarea>
      </div>
      <!-- If there are no DynamicFields, display a placeholder message -->
      <div
        v-if="
          !detailItem.DynamicFields ||
          Object.keys(detailItem.DynamicFields).length === 0
        "
      >
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

  .textarea {
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
  flex-direction: column;
  /* Stack top-row and bottom-row vertically */
  gap: 10px;
  /* Match the margin-bottom of the non-scoped va-select */
}

.top-row,
.bottom-row {
  display: flex;
  justify-content: space-between;
  /* Distribute space evenly between children */
  gap: 10px;
  /* Match the margin-bottom of the non-scoped va-select */
}

.option-select {
  flex: 1;
  /* Allow each select to grow and shrink equally */
  max-width: calc(20% - 10px);
  /* Limit the width to 20% minus the gap size */
}
</style>

<style scoped>
.textarea-container {
  margin-bottom: 1rem;
}

.textarea-label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}

.detailarea {
  width: 100%;
  min-height: 100px;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  resize: vertical;
}
</style>
