import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import VueRouter from 'vue-router'
import router from "./router/index"
import axios from 'axios'
import api from './util/index.js'

Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(VueRouter)

Vue.prototype.$api = api
Vue.prototype.$axios = axios

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
