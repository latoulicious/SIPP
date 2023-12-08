import { createVuestic } from "vuestic-ui";
import "vuestic-ui/styles/essential.css";
import "vuestic-ui/styles/typography.css";
import "./assets/main.css";
import "vuestic-ui/css";
import "material-design-icons-iconfont/dist/material-design-icons.min.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
/* import axios from "@/router/axios"; // Import Axios */

import instance from "@/router/axios";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

// Use Vuestic, Pinia, and Axios
app.use(createVuestic());
app.use(createPinia());
app.use(router);

// Configure Axios and make it available globally
app.config.globalProperties.$axios = instance;

app.mount("#app");
