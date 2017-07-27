import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

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
        context.commit('user', user);
        localStorage.setItem('user', JSON.stringify(user));
        context.commit('loading', false);
        context.commit('auth', true);
        context.commit('message', '已同步');
      }).catch((error) => {
        context.commit('loading', false);
        if (error.response) {
          console.log(error.response.data);
          console.log(error.response.status);
          context.commit('message', 'Fail');
          context.commit('auth', false);
          // 403 auth=false;dialog.open();
        } else if (error.request) {
          context.commit('message', '网络错误');
        }
      });
    },
  },
});
