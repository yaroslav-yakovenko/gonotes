import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueHighlightJS from 'vue-highlight.js'
import 'highlight.js/styles/vs2015.css'
// import 'highlight.js/styles/an-old-hope.css'
// import 'highlight.js/styles/railscasts.css'
// import 'highlight.js/styles/atom-one-dark-reasonable.css'

Vue.use(VueHighlightJS)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
