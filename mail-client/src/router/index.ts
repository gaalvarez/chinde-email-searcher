import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      redirect(to) {
        return { name: "mail-search" };
      },
    },
    {
      path: "/mail-search",
      name: "mail-search",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/MailSearchView.vue"),
    },
  ],
});

export default router;
