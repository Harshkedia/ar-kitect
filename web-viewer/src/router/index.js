import Vue from "vue";
import VueRouter from "vue-router";
import Viewer from "../views/Viewer.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "viewer",
    component: Viewer
  },
  {
    path: "/upload",
    name: "upload",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import("../views/Uploader.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
  meta: { title: "AR Viewer" }
});

export default router;
