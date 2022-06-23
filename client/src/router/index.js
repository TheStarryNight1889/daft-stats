import * as VueRouter from 'vue-router';
// 1. Define route components.
// These can be imported from other files
const HomeView = () => import('../views/HomeView.vue');

// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [{ path: '/', component: HomeView }];

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes, // short for `routes: routes`
});

export default router;
