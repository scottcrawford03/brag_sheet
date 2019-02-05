// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  data () {
    return {
      brags: null
    }
  },
  mounted () {
    axios
      .get('http://brag-sheet-api:8081/')
      .then(response => (this.brags = response))
  },
  template: '<App/>'
})
