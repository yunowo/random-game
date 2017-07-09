import Vue from 'vue';
import App from './App';
import router from './router';
// import 'babel-polyfill';
// require('es6-promise').polyfill();

import './config';

import Hello from './components/Hello';

Vue.component('page-content', Hello);

/* eslint-disable no-new */
new Vue({
  router,
  template: '<App/>',
  components: { App },
});

let Docs = Vue.component('app', App);
const handleSectionTheme = (currentRoute) => {
  let theme = 'game';
  const name = currentRoute.name;

  if (name) {
    if (name === 'game') {
      theme = 'indigo';
    } else if (name === 'lists') {
      theme = 'cyan';
    } else if (name === 'my') {
      theme = 'orange';
    } else if (name === 'about') {
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
