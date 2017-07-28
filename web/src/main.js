import Vue from 'vue';
import VueMaterial from 'vue-material';
import 'vue-material/dist/vue-material.css';
import VueClipboards from 'vue-clipboards';
import VueQr from 'vue-qr';

import App from './App';
import router from './router';
import store from './store';

import About from './components/About';
import PageContent from './components/PageContent';
import LoginDialog from './components/LoginDialog';
import EmptyPlaceholder from './components/EmptyPlaceholder';

Vue.use(VueMaterial);
Vue.use(VueClipboards);

Vue.component('about', About);
Vue.component('page-content', PageContent);
Vue.component('vue-qr', VueQr);
Vue.component('login-dialog', LoginDialog);
Vue.component('empty-placeholder', EmptyPlaceholder);

Vue.material.registerTheme({
  blue: {
    primary: 'light-blue',
    accent: 'pink',
  },
  green: {
    primary: 'green',
    accent: 'pink',
  },
  teal: {
    primary: 'teal',
    accent: 'pink',
  },
  red: {
    primary: 'red',
    accent: 'pink',
  },
});

const handleSectionTheme = (currentRoute) => {
  let theme = 'blue';
  const name = currentRoute.name;

  if (name) {
    if (name === 'game') {
      theme = 'blue';
    } else if (name === 'lists') {
      theme = 'teal';
    } else if (name === 'my') {
      theme = 'green';
    } else if (name === 'error') {
      theme = 'red';
    }
  }

  Vue.material.setCurrentTheme(theme);
};

let AppComponent = Vue.component('app', App);
AppComponent = new AppComponent({
  el: '#app',
  router,
  store,
});

handleSectionTheme(router.currentRoute);

router.beforeEach((to, from, next) => {
  Vue.nextTick(() => {
    const mainContent = document.querySelector('.main-content');

    if (mainContent) {
      mainContent.scrollTop = 0;
    }

    AppComponent.closeSidenav();

    next();
  });
});

router.afterEach((to) => {
  handleSectionTheme(to);
});
