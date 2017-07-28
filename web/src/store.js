import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

axios.defaults.baseURL = process.env.NODE_ENV === 'production' ? 'https://api.liuyun.me/random' : 'http://localhost:7000/random';
axios.defaults.withCredentials = true;

Vue.use(Vuex);

/* eslint no-param-reassign:
["error", { "props": true, "ignorePropertyModificationsFor": ["state"] }] */

export default new Vuex.Store({
  state: {
    user: JSON.parse(localStorage.getItem('user')),
    selectedId: parseInt(localStorage.getItem('selectedId'), 10),
    loading: false,
    auth: localStorage.getItem('user') !== null,
    message: {
      text: '',
      time: Date.now(),
    },
  },
  getters: {
    user: state => state.user,
    selectedId: state => state.selectedId,
    loading: state => state.loading,
    auth: state => state.auth,
    message: state => state.message,
  },
  mutations: {
    user(state, u) {
      state.user = u;
    },
    selectedId(state, s) {
      state.selectedId = s;
      localStorage.setItem('selectedId', s);
    },
    loading(state, l) {
      state.loading = l;
    },
    auth(state, a) {
      state.auth = a;
    },
    message(state, m) {
      state.message = {
        text: m,
        time: Date.now(),
      };
    },
  },
  actions: {
    sync(context) {
      axios.get('/user').then((response) => {
        const user = response.data.data;
        localStorage.setItem('user', JSON.stringify(user));
        context.commit('user', user);
        context.commit('loading', false);
        context.commit('auth', true);
        context.commit('message', '已同步');
      }).catch((error) => {
        context.commit('loading', false);
        if (error.response) {
          if (error.response.status === 401) {
            context.commit('auth', false);
          }
          context.commit('message', `Error ${error.response.status} ${error.response.data.error.status}`);
        } else if (error.request) {
          context.commit('message', 'Network error');
        }
      });
    },
  },
});
