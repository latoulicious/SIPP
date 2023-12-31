<script>
import { defineComponent } from "vue";
import { useModal } from "vuestic-ui";

export default defineComponent({
  data() {
    const { confirm } = useModal();

    const generateItems = (count) => {
      const users = [
        {
          id: 1,
          nama_penyusun: "Leanne Graham",
          mata_pelajaran: "Kimia",
          kelas: "10",
          tahun_ajaran: "Ganjil",
        },
        {
          id: 2,
          nama_penyusun: "Ervin Howell",
          mata_pelajaran: "Kimia",
          kelas: "11",
          tahun_ajaran: "Genap",
        },
        {
          id: 3,
          nama_penyusun: "Clementine Bauch",
          mata_pelajaran: "Kimia",
          kelas: "12",
          tahun_ajaran: "Ganjil & Genap",
        },
        // Add more users if needed
      ];

      // Generate user Array
      return new Array(count).fill(null).map((_, idx) => {
        const user = { ...users[idx % users.length] };
        user.id = idx + 1;
        return user;
      });
    };

    const onButtonClick = () => {
      confirm({
        blur: true,
        title: "Confirm",
        message: "Are you sure you want to delete this item?",
      });
    };

    const columns = [
      { key: "id", sortable: false },
      { key: "nama_penyusun", sortable: false },
      { key: "mata_pelajaran", sortable: false },
      { key: "kelas", sortable: false },
      { key: "tahun_ajaran", sortable: false },
    ];

    const items = generateItems(50); // Adjust the count as needed
    const filtered = items.map((item) => item.id);

    return {
      items,
      columns,
      perPage: 10,
      currentPage: 1,
      isTableStriped: true,
      animated: true,
      selectedRows: [],
      rowEventType: "",
      rowId: "",
      filtered,
      loading: true,
      showModal: false,
      input: "",
      onButtonClick,
    };
  },

  computed: {
    pages() {
      return this.perPage && this.perPage !== 0
        ? Math.ceil(this.filtered.length / this.perPage)
        : this.filtered.length;
    },
  },

  methods: {
    fetchData() {
      // Simulate a delay and then set loading to false
      setTimeout(() => {
        this.loading = false;
      }, 600); // Adjust the delay value (in milliseconds) as needed
    },
  },

  watch: {
    currentPage(newPage, oldPage) {
      // Watch for changes in currentPage and update loading state
      if (newPage !== oldPage) {
        this.loading = true;

        // Simulate fetching data and set loading to false after a delay
        this.fetchData();
      }
    },
  },

  mounted() {
    // Initial data fetching
    this.loading = true;
    this.fetchData();
  },
});
</script>

<template>
  <div
    class="header-container"
    style="display: flex; justify-content: space-between; align-items: center"
  >
    <va-input v-model="input" placeholder="Search"></va-input>
    <va-button-group
      icon-color="#000000"
      preset="secondary"
      border-color="bordered"
    >
      <va-button @click="showModal = !showModal" icon="add">Add</va-button>
      <va-button icon="edit">Edit</va-button>
      <va-button @click="onButtonClick" type="delete" icon="delete"
        >Delete</va-button
      >
    </va-button-group>
  </div>
  <va-modal v-model="showModal" blur size="large" fixed-layout>
    <va-card :bordered="false" stripe>
      <va-card-title>Input Data Penilaian Ranah Pengetahuan</va-card-title>
      <va-card-content>
        <div>
          <div class="modal-container">
            <div>
              <va-input
                v-model="value"
                placeholder="Nama Penyusun"
                label="Nama Penyusun"
                preset="bordered"
                style="width: 100%"
              />
            </div>
            <div style="margin-top: 10px">
              <va-select
                v-model="value"
                :options="options"
                label="Mata Pelajaran"
                placeholder="Pilih Mata Pelajaran"
                preset="bordered"
                style="width: 100%"
              />
            </div>
          </div>
          <div class="txt flex justify-between">
            <div
              class="flex flex-col md6"
              style="margin-right: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Tujuan Pembelajaran"
                placeholder="Deskripsi pencapaian tiga aspek kompetensi (pengetahuan, keterampilan, sikap) murid "
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Pemahaman Bermakna"
                placeholder="Informasi tentang manfaat yang akan peserta didik peroleh setelah mengikuti proses pembelajaran"
                preset="bordered"
              />
            </div>
          </div>
          <div class="txt flex justify-between">
            <div
              class="flex flex-col md6"
              style="margin-right: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Pertanyaan Pemantik"
                placeholder="Pertanyaan yang bertujuan untuk memandu siswa mendapatkan pemahaman sesuai dengan tujuan pembelajaran"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Kegiatan Pembelajaran"
                placeholder="Suatu proses yang mengandung serangkaian kegiatan guru dan siswa atas dasar hubungan timbal balik yang berlangsung dalam situasi edukatif untuk mencapai tujuan tertentu."
                preset="bordered"
              />
            </div>
          </div>
          <div class="txt flex justify-between">
            <div
              class="flex flex-col md6"
              style="margin-right: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Refleksi"
                placeholder="Kegiatan yang dilakukan dalam proses belajar mengajar dalam bentuk penilaian tertulis dan lisan oleh guru untuk siswa dan oleh siswa untuk guru untuk mengekspresikan kesan konstruksif, pesan, harapan, dan kritik terhadap proses pembelajaran"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Glosarium"
                placeholder="Kumpulan istilah-istilah dalam suatu bidang secara alfabetikal yang dilengkapi dengan definisi dan artinya"
                preset="bordered"
              />
            </div>
          </div>
          <div>
            <div
              class="flex flex-col md12"
              style="margin-right: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Daftar Pustaka"
                placeholder="Sumber-sumber referensi yang digunakan dalam pengembangan modul ajar"
                preset="bordered"
              />
            </div>
          </div>
          <va-card :bordered="false" stripe disabled>
            <va-card-title
              >Upload Data Penilaian Ranah Pengetahuan</va-card-title
            >
            <va-card-content>
              <va-file-upload
                v-model="basic"
                dropzone
                undo
                undo-button-text="Restore"
                file-types="doc,docs,rtf,xls,xlsx,ppt,pptx,pdf,txt"
              />
            </va-card-content>
          </va-card>
        </div>
      </va-card-content>
    </va-card>
  </va-modal>
  <div class="table-container">
    <vaDataTable
      :items="items"
      :columns="columns"
      :striped="isTableStriped"
      :current-page="currentPage"
      :per-page="perPage"
      selectable
      :animated="animated"
      :delay="500"
      :loading="loading"
    >
      <template #bodyAppend>
        <tr>
          <td colspan="6">
            <div class="flex justify-center mt-4">
              <div class="pagination-container">
                <va-pagination v-model="currentPage" :pages="pages" />
              </div>
            </div>
          </td>
        </tr>
      </template>
      <template #bodyCellCheckbox="{ value }">
        <input type="checkbox" v-model="selectedRows" :value="value" />
      </template>
    </vaDataTable>
  </div>
</template>

<style>
.table-container {
  border: solid black;
}

.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 10px;
}

.header-container {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>

<style scoped>
.modal-container {
  margin-bottom: 10px;
}

.txt {
  display: flex;
  justify-content: space-between;
}

.va-textarea {
  width: 100%;
  box-sizing: border-box;
  margin-bottom: 20px;
}
</style>
