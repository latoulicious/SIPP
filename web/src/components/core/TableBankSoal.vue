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
  TahunAjar: {}, // Initialize as an empty object
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
  TahunAjar: "Tahun Ajar",
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
      { key: "TahunAjar", label: "Tahun Ajar", sortable: false },
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
      tahunAjarOptions: [],
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
        const tahunResponse = await axios.get(
          "http://localhost:3000/api/public/tahun",
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

        this.tahunAjarOptions = this.extractOptions(
          tahunResponse.data.data,
          "Tahun",
        );
        // console.log("Tahun Ajar options:", this.tahunAjarOptions);

        // Update the items array with BankSoal data
        this.items = response.data.data.map((item) => ({
          ...item,
          ID: item?.ID || "", // Use 'ID' instead of 'id'
          User: item?.User.Name || "",
          Mapel: item?.Mapel.Mapel || "",
          Kelas: item?.Kelas.Kelas || "",
          Jurusan: item?.Jurusan.Jurusan || "",
          TahunAjar: item?.TahunAjar.Tahun || "",
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
          TahunAjarID: this.createdItem.TahunAjarID.toString(),
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

        this.resetCreatedItem();
        // Re-fetch the data to refresh the table
        await this.fetchData();
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
        delete editedData.Jurusan;
        delete editedData.TahunAjar;

        const response = await axios.put(
          `http://localhost:3000/api/bank/${this.editedItem.id}`,
          editedData,
        );

        // Handle the response from the server
        if (response.status === 200) {
          // Update the local item with the edited data
          const itemIndex = this.items.findIndex(
            (item) => item.id === this.editedItem.id,
          );
          if (itemIndex !== -1) {
            this.$set(this.items, itemIndex, {
              ...editedData,
              id: this.editedItem.id,
            });
          }

          console.log("Item updated successfully");
        } else {
          console.error("Failed to update item", response.data);
        }

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

    async openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      axios
        .get(`http://localhost:3000/api/kognitif/${selectedItemId}`)
        .then((response) => {
          const data = response.data.data;
          if (data) {
            console.log("Data:", data); // Log the data

            this.detailItem = reactive({
              BankSoalID: data.BankSoalID || "",
              Soal: data.BankSoal.Soal || "",
              OptionA: data.BankSoal.OptionA || "",
              OptionB: data.BankSoal.OptionB || "",
              OptionC: data.BankSoal.OptionC || "",
              OptionD: data.BankSoal.OptionD || "",
              OptionE: data.BankSoal.OptionE || "",
            });

            console.log("Detail item:", this.detailItem); // Log the detail item

            this.detailModalVisible = true;
          } else {
            console.error("No data received from the server");
          }
        })
        .catch((error) => {
          console.error("Error fetching data for the detail modal:", error);
        });
    },

    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;

      try {
        // Fetch the necessary data directly from the server
        const response = await axios.get(
          `http://localhost:3000/api/bank/${selectedItemId}`,
        );
        const data = response.data.data;

        // Log all properties of the data object
        console.log("Data properties:", Object.keys(data));

        if (data) {
          console.log("Printing row with ID:", selectedItemId);

          const tableBody = [
            [
              { text: "Judul BankSoal", fontSize: 10, bold: true },
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
              "judulBankSoal" in data ? data.judulBankSoal : "N/A",
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
                text: "BankSoal Pembelajaran",
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
      this.editedItem = { ...this.items[id], id: this.items[id].ID }; // Use 'id' instead of 'ID'
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
        >Add BankSoal Pembelajaran</va-button
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
      title="Add Bank Soal"
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
        v-model="createdItem.TahunAjarID"
        :label="displayNames.TahunAjar"
        :options="tahunAjarOptions"
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
      title="Edit Bank Soal"
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
      <va-select
        v-model="editedItem.TahunAjarID"
        :label="displayNames.TahunAjar"
        :options="tahunAjarOptions"
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
