import { createVuestic } from "vuestic-ui";
import "vuestic-ui/styles/essential.css";
import "vuestic-ui/styles/typography.css";
import "./assets/main.css";
import "vuestic-ui/css";
import "material-design-icons-iconfont/dist/material-design-icons.min.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(createVuestic());
app.use(createPinia());
app.use(router);

app.mount("#app");
