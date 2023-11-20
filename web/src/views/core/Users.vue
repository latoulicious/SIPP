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
      showModal: false,
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
  },
});
</script>

<template>
  <va-card class="container" stripe stripe-color="danger">
    <va-card-title>Manage Users</va-card-title>
    <va-card-content>
      <div
        class="header-container"
        style="
          display: flex;
          justify-content: space-between;
          align-items: center;
        "
      >
        <va-input v-model="input" placeholder="Filter..."></va-input>
        <va-button
          @click="showModal = !showModal"
          icon="add"
          preset="secondary"
          border-color="#FFFFFF"
          >Add</va-button
        >
      </div>
      <va-data-table
        class="table-crud"
        :items="items"
        :columns="columns"
        striped
      >
        <template #cell(actions)="{ rowIndex }">
          <va-button
            preset="plain"
            icon="edit"
            @click="openModalToEditItemById(rowIndex)"
          />
          <va-button
            preset="plain"
            icon="delete"
            class="ml-3"
            @click="deleteItemById(rowIndex)"
          />
        </template>
      </va-data-table>

      <va-modal v-model="showModal" blur size="large" fixed-layout>
        <va-card :bordered="false" stripe>
          <va-card-title>Add User</va-card-title>
          <va-card-content>
            <va-form
              class="w-[300px]"
              tag="form"
              @submit.prevent="handleSubmit"
            >
              <va-input v-model="username" label="Username" />

              <va-input
                v-model="password"
                class="mt-3"
                type="password"
                label="Password"
              />

              <va-select
                v-model="value"
                class="mt-3"
                label="Role"
                :options="options"
                clearable
              />
            </va-form>
          </va-card-content>
        </va-card>
      </va-modal>
    </va-card-content>
  </va-card>
</template>

<style lang="scss" scoped>
.table-crud {
  --va-form-element-default-width: 0;

  .va-input {
    display: block;
  }

  &__slot {
    th {
      vertical-align: middle;
    }
  }
}

.modal-crud {
  .va-input {
    display: block;
  }
}

.container {
  margin-left: 60px;
  margin-right: 60px;
  margin-bottom: 10px;
  margin-top: 10px;
  border: shadow black;
}
</style>
