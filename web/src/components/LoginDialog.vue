<template>
  <md-dialog ref="dialog" :mdClickOutsideToClose="false" :mdEscToClose="false">
    <md-icon class="md-size-5x md-primary">casino</md-icon>
    <span class="header">随机游戏</span>
    <md-button class="md-raised md-primary" @click="login">
      <div class="icon-button">
        <md-icon md-src="static/img/icon-github.svg"></md-icon>
        <div>Login with GitHub</div>
      </div>
    </md-button>
  
  </md-dialog>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {

    };
  },
  methods: {
    login() {
      const api = axios.defaults.baseURL;
      window.location.replace(`${api}/login`);
    },
  },
};
</script>

<style lang="scss">
.header {
  display: flex;
  font-size: 22px;
  justify-content: center;
  margin-bottom: 18px;
}

.icon-button {
  display: flex;
  justify-content: center;
  .md-icon {
    margin: auto 8px;
  }
  div {
    color: white;
  }
}
</style>
