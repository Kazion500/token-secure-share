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
    component: ViewLink,
  },
  {
    name: "link",
    path: "/link/:id",
    component: ViewLink,
  },
];
