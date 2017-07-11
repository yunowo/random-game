import Vue from 'vue';
import VueMaterial from 'vue-material';
import 'vue-material/dist/vue-material.css';

Vue.use(VueMaterial);

Vue.material.registerTheme({
  blue: {
    primary: 'light-blue',
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
  orange: {
    primary: 'orange',
    accent: 'purple',
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
  white: {
    primary: 'white',
    accent: 'blue',
  },
});
