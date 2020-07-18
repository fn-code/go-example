import Vue from 'vue'
import Vuex from 'vuex'
import { HTTP } from './plugins/axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    APPNAME: "funcode application",
    Dialog: false,
    users: [],
    userDetail: {},
    links: [
      { name: "Google Inc", url: "www.google.com", icon: 'assignment', iconClass: 'green white--text'},
      { name: "Facebook Inc", url: "www.facebook.com", icon: 'assignment', iconClass: 'blue white--text'},
      { name: "Microsoft Inc", url: "www.microsoft.com", icon: 'assignment', iconClass: 'red white--text'},
      { name: "Apple Inc", url: "www.microsoft.com", icon: 'assignment', iconClass: 'yellow white--text'},
    ]
  },
  getters: {
    counterLinks: function(state) {
      return state.links.length;
    },
    dialogStatus: function(state) {
      return state.Dialog;
    },
    fetchUsers: function (state) {
      HTTP.get('/users')
        .then(respons => {
          state.users = respons.data
        });

      return state.users;
    },

  },
  mutations: {
    fetchUsersByID: function (state, id) {
      HTTP.get("/users/" + id)
        .then(respons => {
          state.userDetail = respons.data
        });
    },

    ADD_PERUSAHAAN: function(state, link) {
      state.links.push(link);
    },
    CHANGE_DIALOG_STATE: function(state) {
      state.Dialog = !state.Dialog;
    },
    REMOVE_LINK: function(state, link) {
      state.links.splice(link, 1);
    },
    REMOVE_ALL: function(state) {
      state.links = [];
    }
  },
  actions: {
    removeLink: function(context, link) {
      context.commit("REMOVE_LINK", link);
    },
    removeAll({commit}) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("REMOVE_ALL");
          resolve();
        }, 5000);
      })
    }
  }
})
