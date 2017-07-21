<template>
  <page-content page-title="我的">
    <div id="main-content">
      <div class="main-my">
        <md-card md-with-hover>
          <div class="userinfo">
            <text></text>
            <image class=".userinfo-avatar" src="/static/img/me.png"></image>
          </div>
          <div>
            <md-button @click="login">Login with GitHub</md-button>
          </div>
        </md-card>
        <md-card md-with-hover>
          <div class="page__title">Random Game</div>
          <div class="page__desc">Early Access 0.1.1</div>
          <hello></hello>
        </md-card>
      </div>
    </div>
  </page-content>
</template>

<style>
.md-card {
  width: 400px;
  margin: 20px;
}

.main-my {
  display: flex;
  flex-flow: row wrap;
}
</style>

<script>
export default {
  data() {
    return {

    };
  }, methods: {
    login() {

    },
  }
};
</script>
