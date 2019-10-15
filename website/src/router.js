import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Repos from './views/Repos.vue'
import axios from 'axios'
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
      beforeEnter: async(to, from, next) => {
        const params = to.params.username
        axios.post("http://localhost:9999/api/v1/getUser",{
          username:params
        }).then((response) => {
          //Valid response with all Repos
          const repos = response.data.repositories
        }).catch((error) => {
          //Catch error from golang backend
          console.log(error)
        })
        next()
      }
    }
  ]
})
