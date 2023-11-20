import { createRouter, createWebHistory } from "vue-router";

import AppLayout from "@/layouts/AppLayout.vue";
import AuthLayout from "@/layouts/AuthLayout.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/dashboard",
      component: AppLayout,
      children: [
        {
          path: "/dashboard",
          name: "dashboard",
          component: () => import("@/views/Dashboard.vue"),
        },
        {
          path: "/capaian",
          name: "capaianpendidikan",
          component: () => import("@/views/core/Capaian.vue"),
        },
        {
          path: "/alurtujuan",
          name: "alurtujuanpendidikan",
          component: () => import("@/views/core/Alur.vue"),
        },
        {
          path: "/modulajar",
          name: "modulajar",
          component: () => import("@/views/core/Modul.vue"),
        },
        {
          path: "/penilaian",
          name: "penilaianranahpengetahuan",
          component: () => import("@/views/core/Asesmen.vue"),
        },
        {
          path: "/soal",
          name: "soal",
          component: () => import("@/views/core/Soal.vue"),
        },
        {
          path: "/banksoal",
          name: "banksoal",
          component: () => import("@/views/core/BankSoal.vue"),
        },
        {
          path: "/users",
          name: "users",
          component: () => import("@/views/core/Users.vue"),
        },
        {
          path: "/settings",
          name: "settings",
          component: () => import("@/views/others/PageNotFound.vue"),
        },
      ],
    },
    {
      path: "/auth",
      component: AuthLayout,
      children: [
        {
          path: "/login",
          name: "login",
          component: () => import("@/views/auth/Login.vue"),
        },
        {
          path: "",
          redirect: { name: "login" },
        },
      ],
    },
    /**  {
      path: '/:catchAll(.*)',
      redirect: { name: 'dashboard' },
    } */
  ],
});

export default router;
