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
  Soal: "",
  OptionA: "",
  OptionB: "",
  OptionC: "",
  OptionD: "",
  OptionE: "",
  KunciJawaban: "",
  Materi: "",
  Indikator: "",
  TingkatKesukaran: "",
};

const displayNames = {
  User: "Nama Penyusun",
  Mapel: "Mata Pelajaran",
  Kelas: "Kelas",
  Jurusan: "Jurusan",
  Soal: "Soal",
  OptionA: "Pilihan A",
  OptionB: "Pilihan B",
  OptionC: "Pilihan C",
  OptionD: "Pilihan D",
  OptionE: "Pilihan E",
  KunciJawaban: "Kunci Jawaban",
  Materi: "Materi",
  Indikator: "Indikator",
  TingkatKesukaran: "Tingkat Kesukaran",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "User", label: "Nama Penyusun", sortable: false },
      { key: "Mapel", label: "Mata Pelajaran", sortable: false },
      { key: "Kelas", label: "Kelas", sortable: false },
      { key: "Jurusan", label: "Jurusan", sortable: false },
      { key: "Materi", label: "Materi", sortable: false },
      { key: "actions", label: "Actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: { ...defaultItem },
      input: ref({ value: "" }),
      items: [],
      usersOptions: [],
      mapelsOptions: [],
      kelasOptions: [],
      jurusanOptions: [],
      textAreaFields: [
        "Soal",
        "OptionA",
        "OptionB",
        "OptionC",
        "OptionD",
        "OptionE",
        "KunciJawaban",
        "Materi",
        "Indikator",
        "TingkatKesukaran",
      ],
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
     * Fetches BankSoal data from the API.
     *
     * Makes requests to the API to get BankSoal and related data.
     * Processes the responses to extract useful options for dropdowns
     * and constructs the items array with formatted BankSoal data.
     * Handles loading state and errors.
     */
    async fetchData() {
      this.loading = true;

      try {
        // Simulate a delay
        await new Promise((resolve) => setTimeout(resolve, 2000));

        const response = await axios.get("http://localhost:3000/api/bank");
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

        // Process the data and update the UI
        console.log("Response from server (BankSoal):", response.data);
        // console.log("Response from server (Users):", userResponse.data);
        // console.log("Response from server (Kelas):", kelasResponse.data);
        console.log("Response from server (Jurusan):", jurusanResponse.data);
        // console.log("Response from server (Mapel):", mapelResponse.data);
        // console.log("Response from server (Tahun):", tahunResponse.data);

        // Populate usersOptions, mapelsOptions, kelasOptions, tahunAjarOptions
        this.usersOptions = this.extractOptions(userResponse.data.data, "Name");
        // console.log("Users options:", this.usersOptions);

        this.kelasOptions = this.extractOptions(
          kelasResponse.data.data,
          "Kelas",
        );
        // console.log("Kelas options:", this.kelasOptions);

        this.jurusanOptions = this.extractOptions(
          jurusanResponse.data.data,
          "Jurusan",
        );
        console.log("Jurusan options:", this.jurusanOptions);

        this.mapelsOptions = this.extractOptions(
          mapelResponse.data.data,
          "Mapel",
        );
        // console.log("Mapels options:", this.mapelsOptions);

        // Update the items array with BankSoal data
        this.items = response.data.data.map((item) => ({
          ...item,
          ID: item?.ID || "", // Use 'ID' instead of 'id'
          User: item?.User.Name || "",
          Mapel: item?.Mapel.Mapel || "",
          Kelas: item?.Kelas.Kelas || "",
          Jurusan: item?.Jurusan.Jurusan || "",
          Soal: item?.Soal || "",
          OptionA: item?.OptionA || "",
          OptionB: item?.OptionB || "",
          OptionC: item?.OptionC || "",
          OptionD: item?.OptionD || "",
          OptionE: item?.OptionE || "",
          KunciJawaban: item?.KunciJawaban || "",
          Materi: item?.Materi || "",
          Indikator: item?.Indikator || "",
          TingkatKesukaran: item?.TingkatKesukaran || "",
        }));

        console.log("BankSoal items:", this.items);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    /**
     * Adds a new bank soal item by making a POST request to the API.
     *
     * The new item data is taken from the component's createdItem property.
     * After successfully adding the item, the items array is updated and the table refreshed.
     */
    async addNewItem() {
      if (!this.isNewData) {
        alert("Please fill in all fields.");
        return;
      }

      try {
        console.log("Creating new item with data:", this.createdItem);

        const response = await axios.post("http://localhost:3000/api/bank", {
          UserID: this.createdItem.UserID.toString(),
          MapelID: this.createdItem.MapelID.toString(),
          KelasID: this.createdItem.KelasID.toString(),
          JurusanID: this.createdItem.JurusanID.toString(),
          Soal: this.createdItem.Soal,
          OptionA: this.createdItem.OptionA,
          OptionB: this.createdItem.OptionB,
          OptionC: this.createdItem.OptionC,
          OptionD: this.createdItem.OptionD,
          OptionE: this.createdItem.OptionE,
          KunciJawaban: this.createdItem.KunciJawaban,
          Materi: this.createdItem.Materi,
          Indikator: this.createdItem.Indikator,
          TingkatKesukaran: this.createdItem.TingkatKesukaran,
        });

        this.items.push({
          ...this.createdItem,
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
     * Edits an existing bank soal item by making a PUT request to the API.
     *
     * Takes the edited item data from the component's editedItem property.
     * Converts the 'ID' field to 'id' and removes other unneeded fields.
     * Makes the PUT request to update the item on the server.
     * Refreshes the data table after successfully updating.
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
        delete editedData.Jurusan;

        await axios.put(
          `http://localhost:3000/api/bank/${editedData.id}`,
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
     * Deletes a bank soal item by ID.
     *
     * Prompts user to confirm deletion.
     * Makes a DELETE request to the API to delete the item.
     * Removes the deleted item from the local items array.
     * Refreshes the data table after deleting.
     * Shows success message.
     * Handles errors from the API request.
     */
    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          const response = await axios.delete(
            `http://localhost:3000/api/bank/${id}`,
          );

          // Check if deletedBankSoal is not null
          if (response.data && response.data.data) {
            const deletedBankSoal = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedBankSoal.id,
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
     * Opens the detail modal popup to show details for the bank soal item
     * at the specified row index.
     *
     * Makes API call to get full data for the item.
     * Populates the detailItem object with the data.
     * Shows the modal popup.
     * Handles errors from API call.
     */
    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/bank/${selectedItemId}`,
        );
        const data = response.data.data; // Access the data object directly
        console.log("Full response data:", data); // Log the full response data

        if (data) {
          console.log("Data:", data);

          this.detailItem = {
            ...defaultItem,
            BankSoal: {
              Soal: data.Soal || "",
              OptionA: data.OptionA || "",
              OptionB: data.OptionB || "",
              OptionC: data.OptionC || "",
              OptionD: data.OptionD || "",
              OptionE: data.OptionE || "",
            },
            // If you need to include DynamicFields, you can add it here
            // DynamicFields: data.DynamicFields || [],
          };

          console.log("Soal:", this.detailItem.BankSoal.Soal);
          console.log(
            "Option Values:",
            ["OptionA", "OptionB", "OptionC", "OptionD", "OptionE"].map(
              (option) => this.detailItem.BankSoal[option],
            ),
          );

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
     *
     * It gets the ID of the selected row, makes a request to the API to get the full data,
     * extracts the relevant question and option data, and uses pdfMake to generate a PDF
     * with the question and options formatted properly.
     *
     * The PDF is generated with page numbers, and the metadata like subject and class is included.
     *
     * Handles any errors from the API request.
     */
    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log(`Selected item ID: ${selectedItemId}`);

      try {
        const response = await axios.get(
          `http://localhost:3000/api/bank/${selectedItemId}`,
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

          // Extract the question and options from the data object
          const question = data.Soal;
          const options = [
            data.OptionA,
            data.OptionB,
            data.OptionC,
            data.OptionD,
            data.OptionE,
          ];

          // Create an array with the question and its options
          const questionsAndOptions = [
            {
              question: question,
              options: options,
            },
          ];

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
      }));
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
      this.editedItem = { ...this.items[id] };
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

    incrementTableKey() {
      this.tableKey++;
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
        >Create Bank Soal</va-button
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
          <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          />
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
      title="Form Input Bank Soal"
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

      <!-- Using va-textarea for other fields -->
      <va-textarea
        v-for="key in textAreaFields"
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
      title="Form Edit Bank Soal"
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
        v-model="editedItem.JurusanID"
        :label="displayNames.Jurusan"
        :options="jurusanOptions"
        class="my-6"
        text-by="label"
        value-by="value"
      />

      <!-- Using va-textarea for other fields -->
      <va-textarea
        v-for="key in textAreaFields"
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
      title="Detail Bank Soal"
      size="large"
      :model-value="detailModalVisible"
      @ok="resetDetailItem"
      @cancel="resetDetailItem"
    >
      <!-- Display the Soal -->
      <div class="textarea-container">
        <label>Soal</label>
        <textarea
          class="detailarea"
          readonly
          :value="detailItem.BankSoal.Soal"
        ></textarea>
      </div>
      <!-- Display the options -->
      <div
        v-for="(option, index) in [
          'OptionA',
          'OptionB',
          'OptionC',
          'OptionD',
          'OptionE',
        ]"
        :key="index"
        class="textarea-container"
      >
        <label>{{ `Option ${String.fromCharCode(65 + index)}` }}</label>
        <textarea
          class="detailarea"
          readonly
          :value="detailItem.BankSoal[option]"
        ></textarea>
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
