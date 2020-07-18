import Vue from 'vue'
import Router from 'vue-router'
import HomePage from './views/Home.vue'
import AboutPage from './views/About'
import PenggunaPage from './views/Pengguna'
import DetailPenggunaPage from './views/DetailPengguna'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/home',
      name: 'home',
      component: HomePage,
      alias: "/"
    },
    {
      path: '/pengguna',
      name: 'pengguna',
      component: PenggunaPage,
    },
    {
      path: "/pengguna/detail/:id",
      name: "detail-pengguna",
      component: DetailPenggunaPage
    },
    {
      path: '/about',
      name: 'about',
      component: AboutPage
    }
  ],
  mode: 'history'
})
