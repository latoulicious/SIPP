<script>
import { defineComponent } from "vue";

const defaultItem = {
  name: "",
  username: "",
  email: "",
};

export default defineComponent({
  data() {
    const items = [
      {
        name: "Leanne Graham",
        username: "Bret",
        email: "Sincere@april.biz",
      },
      {
        name: "Ervin Howell",
        username: "Antonette",
        email: "Shanna@melissa.tv",
      },
      {
        name: "Clementine Bauch",
        username: "Samantha",
        email: "Nathan@yesenia.net",
      },
      {
        name: "Patricia Lebsack",
        username: "Karianne",
        email: "Julianne.OConner@kory.org",
      },
    ];

    const columns = [
      { key: "name", sortable: true },
      { key: "username", sortable: true },
      { key: "email", sortable: true },
      { key: "actions", width: 80 },
    ];

    return {
      items,
      columns,
      editedItemId: null,
      editedItem: null,
      createdItem: { ...defaultItem },
      showModal: false, // Added showModal property
    };
  },

  computed: {
    isNewData() {
      return Object.keys(this.createdItem).every(
        (key) => !!this.createdItem[key],
      );
    },
  },

  methods: {
    resetEditedItem() {
      this.editedItem = null;
      this.editedItemId = null;
    },
    resetCreatedItem() {
      this.createdItem = { ...defaultItem };
      this.showModal = false; // Close the modal when resetting
    },
    deleteItemById(id) {
      this.items = [...this.items.slice(0, id), ...this.items.slice(id + 1)];
    },
    addNewItem() {
      this.items = [...this.items, { ...this.createdItem }];
      this.resetCreatedItem();
    },
    editItem() {
      this.items = [
        ...this.items.slice(0, this.editedItemId),
        { ...this.editedItem },
        ...this.items.slice(this.editedItemId + 1),
      ];
      this.resetEditedItem();
    },
    openModalToEditItemById(id) {
      this.editedItemId = id;
      this.editedItem = { ...this.items[id] };
    },
    toggleAddModal() {
      this.showModal = !this.showModal;
      if (!this.showModal) {
        // Reset the created item when closing the modal
        this.resetCreatedItem();
      }
    },
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
    <va-data-table class="table-crud" :items="items" :columns="columns" striped>
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
      class="modal-crud"
      :model-value="!!editedItem"
      title="Edit Item"
      size="small"
      @ok="editItem"
      @cancel="resetEditedItem"
    >
      <va-input
        v-for="key in Object.keys(editedItem)"
        :key="key"
        v-model="editedItem[key]"
        class="my-6"
        :label="key"
      />
    </va-modal>
    <va-modal
      class="modal-crud"
      :model-value="showModal"
      title="Add Item"
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
</style>
