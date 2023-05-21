import ErrorView from "@/views/ErrorView.vue";
import GenLink from "@/views/GeneratedLink.vue";
import Home from "@/views/Home.vue";
import ViewLink from "@/views/ViewLink.vue";

export const routes = [
  {
    name: "home",
    path: "/",
    component: Home,
  },
  {
    name: "generate-link",
    path: "/generate-link/:id",
    component: GenLink,
  },
  {
    name: "link",
    path: "/:id",
    component: ViewLink,
  },
  {
    name: "error",
    path: "/error",
    component: ErrorView,
  },
];
