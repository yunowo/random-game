import Vue from 'vue';
import Router from 'vue-router';
import Game from '@/pages/Game';
import Lists from '@/pages/Lists';
import My from '@/pages/My';
import Error404 from '@/pages/Error404';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [{
    path: '/random/',
    name: 'game',
    component: Game,
  }, {
    path: '/random/lists',
    name: 'lists',
    component: Lists,
  }, {
    path: '/random/my',
    name: 'my',
    component: My,
  }, {
    path: '*',
    name: 'error',
    component: Error404,
  }],
});
