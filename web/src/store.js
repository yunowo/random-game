import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

/* eslint no-param-reassign:
["error", { "props": true, "ignorePropertyModificationsFor": ["state"] }] */

export default new Vuex.Store({
  state: {
    user: JSON.parse(localStorage.getItem('user')),
  },
  mutations: {
    sync(state) {
      axios.get('/user').then((response) => {
        const user = response.data.data;
        state.user = user;
        localStorage.setItem('user', JSON.stringify(user));
      }).catch((error) => {
        console.log(error);
      });
    },
  },
});
