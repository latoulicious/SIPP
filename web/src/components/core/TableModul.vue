<script>
import { defineComponent } from "vue";
import { ref } from "vue";
import pdfMake from "pdfmake/build/pdfmake";
import pdfFonts from "pdfmake/build/vfs_fonts";
import axios from "axios";

pdfMake.vfs = pdfFonts.pdfMake.vfs;

const defaultItem = {
  User: {}, // Initialize as an empty object
  Kelas: {}, // Initialize as an empty object
  TahunAjar: {}, // Initialize as an empty object
  Sekolah: "",
  AlokasiWaktu: "",
  KompetensiAwal: "",
  ProjekPPancasila: "",
  SaranaPrasarana: "",
  TargetPesertaDidik: "",
  ModelPembelajaran: "",
  TujuanPembelajaran: "",
  PemahamanBermakna: "",
  PertanyaanPemantik: "",
  KegiatanPembelajaran: "",
  Refleksi: "",
  Glosarium: "",
  DaftarPustaka: "",
};

const displayNames = {
  User: "Nama Penyusun",
  TahunAjar: "Tahun Ajar",
  Kelas: "Kelas",
  Sekolah: "Sekolah",
  AlokasiWaktu: "Alokasi Waktu",
  KompetensiAwal: "Kompetensi Awal",
  ProjekPPancasila: "Projek P Pancasila",
  SaranaPrasarana: "Sarana Prasarana",
  TargetPesertaDidik: "Target Peserta Didik",
  ModelPembelajaran: "Model Pembelajaran",
  TujuanPembelajaran: "Tujuan Pembelajaran",
  PemahamanBermakna: "Pemahaman Bermakna",
  PertanyaanPemantik: "Pertanyaan Pemantik",
  KegiatanPembelajaran: "Kegiatan Pembelajaran",
  Refleksi: "Refleksi",
  Glosarium: "Glosarium",
  DaftarPustaka: "Daftar Pustaka",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "User", label: "Nama Penyusun", sortable: false },
      { key: "Kelas", label: "Kelas", sortable: false },
      { key: "TahunAjar", label: "Tahun Ajar", sortable: false },
      { key: "AlokasiWaktu", label: "Alokasi Waktu", sortable: false },
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
      kelasOptions: [],
      tahunAjarOptions: [],
      textAreaFields: [
        "Sekolah",
        "AlokasiWaktu",
        "KompetensiAwal",
        "ProjekPPancasila",
        "SaranaPrasarana",
        "TargetPesertaDidik",
        "ModelPembelajaran",
        "TujuanPembelajaran",
        "PemahamanBermakna",
        "PertanyaanPemantik",
        "KegiatanPembelajaran",
        "Refleksi",
        "Glosarium",
        "DaftarPustaka",
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

        const response = await axios.get("http://localhost:3000/api/modul");
        const userResponse = await axios.get(
          "http://localhost:3000/api/public/user",
        );
        const kelasResponse = await axios.get(
          "http://localhost:3000/api/public/kelas",
        );
        const tahunResponse = await axios.get(
          "http://localhost:3000/api/public/tahun",
        );

        // Process the data and update the UI
        console.log("Response from server (Modul):", response.data);
        // console.log("Response from server (Users):", userResponse.data);
        // console.log("Response from server (Kelas):", kelasResponse.data);
        // console.log("Response from server (Tahun):", tahunResponse.data);

        // Populate usersOptions, mapelsOptions, kelasOptions, tahunAjarOptions
        this.usersOptions = this.extractOptions(userResponse.data.data, "Name");
        // console.log("Users options:", this.usersOptions);

        this.kelasOptions = this.extractOptions(
          kelasResponse.data.data,
          "Kelas",
        );
        // console.log("Kelas options:", this.kelasOptions);

        this.tahunAjarOptions = this.extractOptions(
          tahunResponse.data.data,
          "Tahun",
        );
        // console.log("Tahun Ajar options:", this.tahunAjarOptions);

        // Update the items array with Modul data
        this.items = response.data.data.map((item) => ({
          ...item,
          ID: item.ID || "", // Use 'ID' instead of 'id'
          User: item.User.Name || "",
          Kelas: item.Kelas.Kelas || "",
          TahunAjar: item.TahunAjar.Tahun || "",
          Sekolah: item.sekolah || "",
          AlokasiWaktu: item.alokasiWaktu || "",
          KompetensiAwal: item.kompetensiAwal || "",
          ProjekPPancasila: item.projekPPancasila || "",
          SaranaPrasarana: item.saranaPrasarana || "",
          TargetPesertaDidik: item.targetPesertaDidik || "",
          ModelPembelajaran: item.modelPembelajaran || "",
          TujuanPembelajaran: item.tujuanPembelajaran || "",
          PemahamanBermakna: item.pemahamanBermakna || "",
          PertanyaanPemantik: item.pertanyaanPemantik || "",
          KegiatanPembelajaran: item.kegiatanPembelajaran || "",
          Refleksi: item.refleksi || "",
          Glosarium: item.glosarium || "",
          DaftarPustaka: item.daftarPustaka || "",
        }));

        console.log("modul Items:", this.items);
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

        const response = await axios.post("http://localhost:3000/api/modul", {
          UserID: this.createdItem.UserID.toString(),
          KelasID: this.createdItem.KelasID.toString(),
          TahunAjarID: this.createdItem.TahunAjarID.toString(),
          Sekolah: this.createdItem.Sekolah,
          AlokasiWaktu: this.createdItem.AlokasiWaktu,
          KompetensiAwal: this.createdItem.KompetensiAwal,
          ProjekPPancasila: this.createdItem.ProjekPPancasila,
          SaranaPrasarana: this.createdItem.SaranaPrasarana,
          TargetPesertaDidik: this.createdItem.TargetPesertaDidik,
          ModelPembelajaran: this.createdItem.ModelPembelajaran,
          TujuanPembelajaran: this.createdItem.TujuanPembelajaran,
          PemahamanBermakna: this.createdItem.PemahamanBermakna,
          PertanyaanPemantik: this.createdItem.PertanyaanPemantik,
          KegiatanPembelajaran: this.createdItem.KegiatanPembelajaran,
          Refleksi: this.createdItem.Refleksi,
          Glosarium: this.createdItem.Glosarium,
          DaftarPustaka: this.createdItem.DaftarPustaka,
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

    async editItem() {
      try {
        // Create a deep copy of the edited item
        const editedData = JSON.parse(JSON.stringify(this.editedItem));

        /// Change 'ID' to 'id'
        editedData.id = editedData.ID;
        delete editedData.ID;
        delete editedData.User;
        delete editedData.Kelas;
        delete editedData.TahunAjar;

        // console.log("Edited item:", this.editedItem);
        // console.log("Edited item ID:", this.editedItem.id);

        const response = await axios.put(
          `http://localhost:3000/api/modul/${editedData.id}`,
          editedData,
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
            `http://localhost:3000/api/modul/${id}`,
          );

          // Check if deletedModul is not null
          if (response.data && response.data.data) {
            const deletedModul = response.data.data;

            // Remove the item from the items array
            this.items = this.items.filter(
              (item) => item.id !== deletedModul.id,
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

    openDetailModal(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;
      console.log("Opening detail modal with ID:", selectedItemId);

      axios
        .get(`http://localhost:3000/api/modul/${selectedItemId}`)
        .then((response) => {
          const data = response.data.data;
          if (data) {
            this.detailItem = {
              Sekolah: data.sekolah || "",
              AlokasiWaktu: data.alokasiWaktu || "",
              KompetensiAwal: data.kompetensiAwal || "",
              ProjekPPancasila: data.projekPPancasila || "",
              SaranaPrasarana: data.saranaPrasarana || "",
              TargetPesertaDidik: data.targetPesertaDidik || "",
              ModelPembelajaran: data.modelPembelajaran || "",
              TujuanPembelajaran: data.tujuanPembelajaran || "",
              PemahamanBermakna: data.pemahamanBermakna || "",
              PertanyaanPemantik: data.pertanyaanPemantik || "",
              KegiatanPembelajaran: data.kegiatanPembelajaran || "",
              Refleksi: data.refleksi || "",
              Glosarium: data.glosarium || "",
              DaftarPustaka: data.daftarPustaka || "",
              // id: selectedItemId,
              // User: data.User ? data.User.Name || "" : "",
              // Mapel: data.Mapel ? data.Mapel.Mapel || "" : "",
              // Kelas: data.Kelas ? data.Kelas.Kelas || "" : "",
              // TahunAjar: data.TahunAjar ? data.TahunAjar.Tahun || "" : "",
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

    async printRow(rowIndex) {
      const selectedItemId = this.filteredItems[rowIndex].ID;

      try {
        // Fetch the necessary data directly from the server
        const response = await axios.get(
          `http://localhost:3000/api/modul/${selectedItemId}`,
        );
        const data = response.data.data;

        // Log all properties of the data object
        console.log("Data properties:", Object.keys(data));

        if (data) {
          console.log("Printing row with ID:", selectedItemId);

          const tableBody = [
            [
              { text: "Sekolah", fontSize: 10, bold: true },
              { text: "Alokasi Waktu", fontSize: 10, bold: true },
              { text: "Kompetensi Awal", fontSize: 10, bold: true },
              { text: "Projek P Pancasila", fontSize: 10, bold: true },
              { text: "Sarana Prasarana", fontSize: 10, bold: true },
              {
                text: "Target Peserta Didik",
                fontSize: 10,
                bold: true,
              },
              { text: "Model Pembelajaran", fontSize: 10, bold: true },
              {
                text: "Tujuan Pembelajaran",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Pemahaman Bermakna",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Pertanyaan Pemantik",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Kegiatan Pembelajaran",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Refleksi",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Glosarium",
                fontSize: 10,
                bold: true,
              },
              {
                text: "Daftar Pustaka",
                fontSize: 10,
                bold: true,
              },
            ],
            [
              "sekolah" in data ? data.sekolah : "N/A",
              "alokasiWaktu" in data ? data.alokasiWaktu : "N/A",
              "kompetensiAwal" in data ? data.kompetensiAwal : "N/A",
              "projekPPancasila" in data ? data.projekPPancasila : "N/A",
              "saranaPrasarana" in data ? data.saranaPrasarana : "N/A",
              "targetPesertaDidik" in data ? data.targetPesertaDidik : "N/A",
              "modelPembelajaran" in data ? data.modelPembelajaran : "N/A",
              "tujuanPembelajaran" in data ? data.tujuanPembelajaran : "N/A",
              "pemahamanBermakna" in data ? data.pemahamanBermakna : "N/A",
              "pertanyaanPemantik" in data ? data.pertanyaanPemantik : "N/A",
              "kegiatanPembelajaran" in data
                ? data.kegiatanPembelajaran
                : "N/A",
              "refleksi" in data ? data.refleksi : "N/A",
              "glosarium" in data ? data.glosarium : "N/A",
              "daftarPustaka" in data ? data.daftarPustaka : "N/A",
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
        >Add Modul Ajar</va-button
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
      title="Add Modul Ajar"
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
      title="Edit Modul Ajar"
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
      title="Detail Modul Ajar"
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
