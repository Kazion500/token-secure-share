import App from "@/App.vue";
import { routes } from "@/routes";
import { createPinia } from "pinia";
import { createApp } from "vue";
import "./style.css";

import * as VueRouter from "vue-router";

const router = VueRouter.createRouter({
  history: VueRouter.createWebHistory(),
  routes,
});

const pinia = createPinia();
const app = createApp(App);

app.use(router);
app.use(pinia);
app.mount("#app");
