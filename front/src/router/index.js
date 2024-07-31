import Vue from 'vue'
import VueRouter from 'vue-router'
import MainView from '@/view/MainView.vue'
import ProblemView from '@/view/ProblemView.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', name: "index", component: MainView },
  { path: '/problem', name: "problem", component: ProblemView },
]

const router = new VueRouter({
  routes
})

export default router