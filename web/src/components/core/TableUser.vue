<script>
import { defineComponent } from "vue";
import axios from "axios";

const defaultItem = {
  username: "",
  password: "",
  name: "",
  nip: "",
  golongan: "",
  role: "",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "username", sortable: false },
      { key: "name", sortable: false },
      { key: "nip", sortable: false },
      { key: "golongan", sortable: false },
      { key: "role", sortable: false },
      { key: "actions", width: 80 },
    ];

    return {
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: { ...defaultItem },
      items: [],
      showModal: false,
    };
  },

  computed: {
    filteredInputFields() {
      // Exclude 'password' field from the inputFields
      return this.inputFields.filter((key) => key !== "password");
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
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          // Handle the case where the token is not available (e.g., redirect to login)
          console.error("JWT token not available");
          return;
        }

        console.log("Token:", jwtToken); // Log the token for debugging

        const response = await this.$axios.get(
          "http://localhost:3000/api/user",
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        this.items = response.data.data.map((user) => ({
          username: user.Username,
          password: user.Password,
          name: user.Name,
          nip: user.NIP,
          golongan: user.Golongan,
          role: user.Role,
          id: user.ID,
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },

    async addNewItem() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          console.error("JWT token not available");
          // Handle the case where the token is not available (e.g., redirect to login)
          return;
        }

        const response = await axios.post(
          "http://localhost:3000/api/user",
          this.createdItem,
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        // Assuming the API response is a single user object
        const newUser = {
          username: response.data.data.Username,
          password: response.data.data.Password,
          name: response.data.data.Name,
          nip: response.data.data.NIP,
          golongan: response.data.data.Golongan,
          role: response.data.data.Role,
          id: response.data.data.ID,
        };

        this.items.push(newUser);
        this.resetCreatedItem();
        this.fetchData();
      } catch (error) {
        console.error("Error adding new item:", error);
      }
    },

    async editItem() {
      try {
        const jwtToken = localStorage.getItem("jwtToken");

        if (!jwtToken) {
          console.error("JWT token not available");
          // Handle the case where the token is not available (e.g., redirect to login)
          return;
        }

        const editedData = {
          username: this.editedItem.username,
          password: this.editedItem.password,
          name: this.editedItem.name,
          nip: this.editedItem.nip,
          golongan: this.editedItem.golongan,
          role: this.editedItem.role,
        };

        const response = await axios.put(
          `http://localhost:3000/api/user/${this.editedItem.id}`,
          editedData,
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          },
        );

        // Update the item directly without using $set
        this.items[this.editedItemId] = {
          username: response.data.data.Username,
          password: response.data.data.Password,
          name: response.data.data.Name,
          nip: response.data.data.NIP,
          golongan: response.data.data.Golongan,
          role: response.data.data.Role,
          id: response.data.data.ID,
        };

        this.resetEditedItem();
        this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    async deleteItemById(id) {
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
          `http://localhost:3000/api/user/${this.items[id].id}`,
          { headers },
        );

        // Remove the item from the array
        this.items.splice(id, 1);
      } catch (error) {
        console.error("Error deleting item:", error);
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
        >Add User</va-button
      >
    </va-button-group>
  </div>
  <div>
    <va-data-table :items="items" :columns="columns" striped>
      <template #cell(actions)="{ rowIndex }">
        <div class="action-buttons">
          <!-- <va-button
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
      title="Add user"
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
        :type="key === 'password' ? 'password' : 'text'"
        v-model="createdItem[key]"
        class="my-6"
      />
    </va-modal>
    <va-modal
      blur
      class="modal-crud"
      :model-value="!!editedItem"
      title="Edit User"
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
        :type="key === 'password' ? 'password' : 'text'"
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
