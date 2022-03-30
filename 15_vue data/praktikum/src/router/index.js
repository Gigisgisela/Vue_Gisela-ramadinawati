import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import deskripsi from '../views/DeskripsiViews.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/deskripsi/:id',
    name: 'deskripsi',
    
    component: deskripsi
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
