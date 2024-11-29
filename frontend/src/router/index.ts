import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
 
// 2. 配置路由
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    // 路由跳转
    redirect: "/Index"
  },
  {
    path: "/Index",
    name: "Index",
    component: () => import("../views/Index.vue"),
    meta: {
        title: "首页"
    }
  },
  {
    path: "/Tool1",
    name: "Tool1",
    component: () => import("../views/Tool1.vue"),
    meta: {
        title: "工具1"
    }
  },
  {
    path: "/Tool2",
    name: "Tool2",
    component: () => import("../views/Tool2.vue"),
    meta: {
        title: "工具2"
    }
  },
];
// 1.返回一个 router 实列，为函数，里面有配置项（对象） history
const router = createRouter({
  history: createWebHistory(),
  routes,
});
 
// 3导出路由   然后去 main.ts 注册 router.ts
export default router