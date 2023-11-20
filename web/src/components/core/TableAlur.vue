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
          name: "Leanne Graham",
          username: "Bret",
          email: "Sincere@april.biz",
          phone: "1-770-736-8031 x56442",
          website: "hildegard.org",
        },
        {
          id: 2,
          name: "Ervin Howell",
          username: "Antonette",
          email: "Shanna@melissa.tv",
          phone: "010-692-6593 x09125",
          website: "anastasia.net",
        },
        {
          id: 3,
          name: "Clementine Bauch",
          username: "Samantha",
          email: "Nathan@yesenia.net",
          phone: "1-463-123-4447",
          website: "ramiro.info",
        },
        {
          id: 4,
          name: "Patricia Lebsack",
          username: "Karianne",
          email: "Julianne.OConner@kory.org",
          phone: "493-170-9623 x156",
          website: "kale.biz",
        },
        {
          id: 5,
          name: "Chelsey Dietrich",
          username: "Kamren",
          email: "Lucio_Hettinger@annie.ca",
          phone: "(254)954-1289",
          website: "demarco.info",
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
      { key: "id", sortable: true },
      { key: "username", sortable: true },
      { key: "name", sortable: true },
      { key: "email", sortable: true },
      { key: "phone", sortable: true },
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
    <va-input v-model="input" placeholder="Filter..."></va-input>
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
      <va-card-title>Input Data Alur Tujuan Pembelajaran</va-card-title>
      <va-card-content>
        <div>
          <div class="modal-container">
            <div>
              <va-input
                v-model="value"
                placeholder="Nama Penyusun ATP"
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
                label="Judul Elemen"
                placeholder="Judul mengenai elemen pemahaman suatu mata pelajaran"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Keterangan Elemen"
                placeholder="Menjelaskan inti dari judul elemen tersebut"
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
                label="Keterangan Proses Mengamati"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Keterangan Proses Mempertanyakan"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
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
                label="Keterangan Proses Merencanakan"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Keterangan Proses Memproses"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
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
                label="Keterangan Proses Mengevaluasi"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajarann"
                preset="bordered"
              />
            </div>
            <div
              class="flex flex-col md6"
              style="margin-left: 10px; width: 100%"
            >
              <va-textarea
                v-model="value"
                label="Keterangan Proses Mengkomunikasikan"
                placeholder="Menjelaskan output yang diharapkan dari tujuan pembelajaran untuk mencapai capaian pembelajaran"
                preset="bordered"
              />
            </div>
          </div>
          <va-card :bordered="false" stripe disabled>
            <va-card-title>Upload Data Capaian</va-card-title>
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
    <va-data-table
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
    </va-data-table>
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
