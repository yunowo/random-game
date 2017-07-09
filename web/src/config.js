import Vue from 'vue';
import VueMaterial from 'vue-material';

Vue.use(VueMaterial);

Vue.material.registerTheme({
  app: {
    primary: 'blue',
    accent: 'pink',
  },
  indigo: {
    primary: 'indigo',
    accent: 'pink',
  },
  green: {
    primary: 'green',
    accent: 'pink',
  },
  'light-blue': {
    primary: 'light-blue',
    accent: 'yellow',
  },
  teal: {
    primary: 'teal',
    accent: 'orange',
  },
  'blue-grey': {
    primary: 'blue-grey',
    accent: 'blue',
  },
  cyan: {
    primary: 'cyan',
    accent: 'pink',
  },
  red: {
    primary: 'red',
    accent: 'pink',
  },
});
