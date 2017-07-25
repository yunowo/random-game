import Vue from 'vue';
import VueMaterial from 'vue-material';
import 'vue-material/dist/vue-material.css';
import VueClipboards from 'vue-clipboards';

Vue.use(VueMaterial);
Vue.use(VueClipboards);

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
  red: {
    primary: 'red',
    accent: 'pink',
  },
  white: {
    primary: 'white',
    accent: 'blue',
  },
});
