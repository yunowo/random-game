import Vue from 'vue';
import axios from 'axios';
import VueQr from 'vue-qr';

import App from './App';
import router from './router';
import './config';

import Hello from './components/Hello';
import PageContent from './components/PageContent';

Vue.component('hello', Hello);
Vue.component('page-content', PageContent);
Vue.component('vue-qr', VueQr);

/* eslint-disable no-new */
new Vue({
  router,
  template: '<App/>',
  components: {
    App,
  },
});

let Docs = Vue.component('app', App);
const handleSectionTheme = (currentRoute) => {
  let theme = 'blue';
  const name = currentRoute.name;

  if (name) {
    if (name === 'game') {
      theme = 'blue';
    } else if (name === 'lists') {
      theme = 'orange';
    } else if (name === 'my') {
      theme = 'green';
    } else if (name === 'error') {
      theme = 'red';
    }
  }

  Vue.material.setCurrentTheme(theme);
};

Docs = new Docs({
  el: '#app',
  router,
});

handleSectionTheme(router.currentRoute);

router.beforeEach((to, from, next) => {
  Vue.nextTick(() => {
    const mainContent = document.querySelector('.main-content');

    if (mainContent) {
      mainContent.scrollTop = 0;
    }

    Docs.closeSidenav();

    next();
  });
});

router.afterEach((to) => {
  handleSectionTheme(to);
});

localStorage.setItem('selectedId', 1);

const API_HOST = 'http://localhost:7000/random';
// axios.defaults.baseURL = 'https://api.liuyun.me/random';
axios.defaults.baseURL = 'http://localhost:7000/random';
axios.defaults.withCredentials = true;
