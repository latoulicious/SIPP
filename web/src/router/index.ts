import { createRouter, createWebHistory } from "vue-router";

import requireAuth from "@/router/auth.js";
import AppLayout from "@/layouts/AppLayout.vue";
import AuthLayout from "@/layouts/AuthLayout.vue";

const router = createRouter({
  history: createWebHistory("/"),
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
          beforeEnter: requireAuth,
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
          component: () => import("@/views/core/Penilaian.vue"),
        },
        {
          path: "/asesmen",
          name: "asesmen",
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
        // Protected Routes
        {
          path: "/kelas",
          name: "kelas",
          component: () => import("@/views/others/PageNotFound.vue"),
          beforeEnter: requireAuth,
        },
        {
          path: "/mapel",
          name: "mapel",
          component: () => import("@/views/others/PageNotFound.vue"),
          beforeEnter: requireAuth,
        },
        {
          path: "/tahun",
          name: "tahun",
          component: () => import("@/views/others/PageNotFound.vue"),
          beforeEnter: requireAuth,
        },
        {
          path: "/user",
          name: "user",
          component: () => import("@/views/core/User.vue"),
          beforeEnter: requireAuth,
        },
        {
          path: "/change-password",
          name: "change-password",
          component: () => import("@/views/auth/ChangePassword.vue"),
          beforeEnter: requireAuth,
        },
        // Misc Routes
        {
          path: "/settings",
          name: "settings",
          component: () => import("@/views/others/PageNotFound.vue"),
        },
        {
          path: "/notfound",
          name: "notfound",
          component: () => import("@/views/others/PageNotFound.vue"),
        },
        {
          path: "/unauthorized",
          name: "unauthorized",
          component: () => import("@/views/others/Unauthorized.vue"),
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
