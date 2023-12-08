<script>
import { defineComponent } from "vue";
import axios from "axios";

const defaultItem = {
  username: "",
  password: "",
  name: "",
  mapel: "",
  role: "",
};

export default defineComponent({
  data() {
    const columns = [
      { key: "username", sortable: false },
      { key: "password", sortable: false },
      { key: "name", sortable: false },
      { key: "mapel", sortable: false },
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
    async deleteItemById(id) {
      try {
        // Use the correct ID when making the DELETE request
        await axios.delete(
          `http://localhost:3000/api/users/${this.items[id].id}`,
        );

        // Remove the item from the array
        this.items.splice(id, 1);
      } catch (error) {
        console.error("Error deleting item:", error);
      }
    },

    async addNewItem() {
      try {
        const response = await axios.post(
          "http://localhost:3000/api/users",
          this.createdItem,
        );

        // Assuming the API response is a single user object
        const newUser = {
          username: response.data.data.Username,
          password: response.data.data.Password,
          name: response.data.data.Name,
          mapel: response.data.data.Mapel,
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
        const editedData = {
          username: this.editedItem.username,
          password: this.editedItem.password,
          name: this.editedItem.name,
          mapel: this.editedItem.mapel,
          role: this.editedItem.role,
        };

        const response = await axios.put(
          `http://localhost:3000/api/users/${this.editedItem.id}`,
          editedData,
        );

        // Update the item directly without using $set
        this.items[this.editedItemId] = {
          username: response.data.data.Username,
          password: response.data.data.Password,
          name: response.data.data.Name,
          mapel: response.data.data.Mapel,
          role: response.data.data.Role,
          id: response.data.data.ID,
        };

        this.resetEditedItem();
        this.fetchData();
      } catch (error) {
        console.error("Error editing item:", error);
      }
    },

    async fetchData() {
      try {
        const response = await axios.get("http://localhost:3000/api/users");

        this.items = response.data.data.map((user) => ({
          username: user.Username,
          password: user.Password,
          name: user.Name,
          mapel: user.Mapel,
          role: user.Role,
          id: user.ID, // Add this line to include the id in your items
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
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
      blur
      class="modal-crud"
      :model-value="!!editedItem"
      title="Edit User"
      size="small"
      @ok="editItem"
      @cancel="resetEditedItem"
    >
      <va-input
        v-for="key in inputFields"
        :key="key"
        v-model="editedItem[key]"
        class="my-6"
        :label="key"
      />
    </va-modal>

    <va-modal
      blur
      class="modal-crud"
      :model-value="showModal"
      title="Add user"
      size="small"
      @ok="addNewItem"
      @cancel="resetCreatedItem"
    >
      <va-input
        v-for="key in Object.keys(createdItem)"
        :key="key"
        v-model="createdItem[key]"
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
  }
}
</style>
