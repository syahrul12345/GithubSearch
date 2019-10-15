import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Repos from './views/Repos.vue'
import ErrorPage from './views/Error.vue'
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/repos/:username/',
      name: 'repos',
      component: Repos,
    },
    {
      path:'/error',
      name: 'error',
      component: ErrorPage
    }

  ]
})