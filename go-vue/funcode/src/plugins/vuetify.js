import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import 'vuetify/src/stylus/app.styl'

Vue.use(Vuetify, {
  iconfont: 'md',
  theme: {
    primary: '#333333',
    secondary: '#F03',
    accent: '#F03',
    error: '#b71c1c'
  }
})
