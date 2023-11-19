<script>
import { defineComponent } from "vue";

export default defineComponent({
  data() {
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

      return new Array(count).fill(null).map((_, idx) => {
        const user = { ...(users[idx % users.length]) };
        user.id = idx + 1;
        return user;
      });
    };

    const columns = [
      { key: "id", sortable: true },
      { key: "username", sortable: true },
      { key: "name", sortable: true },
      { key: "email", sortable: true },
      { key: "phone", sortable: true },
    ];

    const items = generateItems(100); // Adjust the count as needed

    const filtered = items.map(item => item.id);

    return {
      items,
      columns,
      perPage: 10,
      currentPage: 1,
      isTableStriped: true,
      animated: true,
      rowEventType: "",
      rowId: "",
      filtered,
      loading: true,
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
              <va-pagination
                v-model="currentPage"
                :pages="pages"
              />
            </div>
          </div>
        </td>
      </tr>
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

</style>
