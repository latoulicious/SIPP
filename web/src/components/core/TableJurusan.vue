<script>
import { defineComponent } from "vue";
import axios from "axios";

const defaultItem = {
  jurusan: "",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "jurusan", sortable: false },
      { key: "actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: { ...defaultItem },
      items: [],
      showModal: false,
      loading: false,
    };
  },

  computed: {
    filteredInputFields() {
      return this.inputFields.filter((key) => key);
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
     * Fetches jurusan data from the API and updates component state.
     * Requires valid JWT token with Authorization header.
     * Handles loading state and errors.
     */
    async fetchData() {
      this.loading = true;

      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          // Handle the case where the token is not available (e.g., redirect to login)
          console.error("JWT token not available");
          return;
        }

        console.log("Token:", jwtToken); // Log the token for debugging

        const response = await this.$axios.get(
          "http://localhost:3000/api/jurusan",
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        this.items = response.data.data.map((jurusan) => ({
          jurusan: jurusan.Jurusan,
          id: jurusan.ID,
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        this.loading = false;
      }
    },

    /**
     * Adds a new jurusan item by making a POST request to the API.
     * Requires valid JWT token.
     * Handles API response and updates component state on success.
     * Logs errors on failure.
     */
    async addNewItem() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          console.error("JWT token not available");
          // Handle the case where the token is not available (e.g., redirect to login)
          return;
        }

        const response = await axios.post(
          "http://localhost:3000/api/jurusan",
          this.createdItem,
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        // Assuming the API response is a single user object
        const newJurusan = {
          jurusan: response.data.data.Jurusan,
          id: response.data.data.ID,
        };

        this.items.push(newJurusan);
        this.resetCreatedItem();
        await this.fetchData();
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },

    /**
     * Edits an existing jurusan item by making a PUT request to the API.
     * Requires valid JWT token.
     * Handles API response and updates component state on success.
     * Logs errors on failure.
     */
    async editItem() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          console.error("JWT token not available");
          // Handle the case where the token is not available (e.g., redirect to login)
          return;
        }

        const editedData = {
          jurusan: this.editedItem.jurusan,
        };

        const response = await axios.put(
          `http://localhost:3000/api/jurusan/${this.editedItem.id}`,
          editedData,
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        // Update the item directly without using $set
        this.items[this.editedItemId] = {
          jurusan: response.data.data.Jurusan,
          id: response.data.data.ID,
        };

        this.resetEditedItem();
        await this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    /**
     * Deletes a jurusan item by ID by making a DELETE request to the API.
     * Requires valid JWT token.
     * Removes the deleted item from component state on success.
     * Logs errors on failure.
     */
    async deleteItemById(id) {
      if (window.confirm("Are you sure you want to delete this item?")) {
        try {
          const jwtToken = localStorage.getItem("jwtToken");

          if (!jwtToken) {
            console.error("JWT token not available");
            // Handle the case where the token is not available (e.g., redirect to login)
            return;
          }

          // Add headers with the authentication token
          const headers = {
            Authorization: `Bearer ${jwtToken}`,
          };

          // Make the DELETE request with headers
          await axios.delete(
            `http://localhost:3000/api/jurusan/${this.items[id].id}`,
            { headers },
          );

          // Remove the item from the array
          this.items.splice(id, 1);
          await this.fetchData();
          // Optionally, you can show a success message
          alert("Item deleted successfully");
        } catch (error) {
          console.error("Error deleting item:", error);
        }
      }
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
      this.editedItem = { ...this.items[id], id: this.items[id].id }; // Ensure the id is set
    },

    toggleAddModal() {
      this.showModal = !this.showModal;
      if (!this.showModal) {
        this.resetCreatedItem();
      }
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
    <va-input v-model="input" placeholder="Search"></va-input>
    <va-button-group
      icon-color="#000000"
      preset="secondary"
      border-color="bordered"
    >
      <va-button @click="toggleAddModal" preset="secondary" icon="add"
        >Add Jurusan</va-button
      >
    </va-button-group>
  </div>
  <div>
    <va-data-table :items="items" :columns="columns" striped>
      <template #cell(actions)="{ rowIndex }">
        <div class="action-buttons">
          <!--<va-button
            preset="plain"
            icon="remove_red_eye"
            @click="deleteItemById(rowIndex)"
          /> -->
          <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          />
          <va-button
            preset="plain"
            icon="delete"
            @click="deleteItemById(rowIndex)"
          />
        </div>
      </template>
    </va-data-table>
    <!-- Modal Content -->
    <va-modal
      class="modal-crud"
      stripe
      title="Form Input Jurusan"
      size="small"
      blur
      :model-value="showModal"
      @ok="addNewItem"
      @cancel="resetCreatedItem"
    >
      <va-input
        v-for="key in Object.keys(createdItem)"
        :key="key"
        :label="key"
        v-model="createdItem[key]"
        class="my-6"
      />
    </va-modal>
    <va-modal
      blur
      class="modal-crud"
      :model-value="!!editedItem"
      title="Form Edit Jurusan"
      size="small"
      @ok="editItem"
      @cancel="resetEditedItem"
    >
      <va-input
        v-for="key in filteredInputFields"
        :key="key"
        v-model="editedItem[key]"
        class="my-6"
        :label="key"
      />
    </va-modal>
  </div>
</template>

<style>
.action-buttons {
  display: flex;
  gap: 8px; /* Adjust the gap to your preference */
}

.modal-crud {
  .va-input {
    display: block;
    margin-bottom: 10px;
  }
}
</style>
