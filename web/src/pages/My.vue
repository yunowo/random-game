<template>
  <page-content page-title="我的">
    <div id="main-content">
      <div class="main-my">
        <md-card md-with-hover>
          <md-toolbar>
            <h1 class="md-title">我的</h1>
            <md-button class="md-icon-button">
              <md-icon>filter_list</md-icon>
            </md-button>
  
            <md-button class="md-icon-button">
              <md-icon>search</md-icon>
            </md-button>
          </md-toolbar>
          <div class="userinfo">
            <text></text>
            <image class=".userinfo-avatar" src="/static/img/me.png"></image>
          </div>
          <div>
            <md-button @click="login">
              <md-icon md-src="static/img/icon-github.svg"></md-icon>
              <div>Login with GitHub</div>
            </md-button>
            <md-button @click="sync">
              <md-icon>cloud_download</md-icon>
              <div>Sync with server</div>
            </md-button>
          </div>
        </md-card>
        <md-card md-with-hover>
          <md-card-header>
            <div class="md-title">Random Game</div>
            <div class="md-subhead">Early Access 0.2.0</div>
          </md-card-header>
          <md-card-expand>
            <md-card-actions>
              <md-button>Action</md-button>
              <md-button>Action</md-button>
              <span style="flex: 1"></span>
              <md-button class="md-icon-button" md-expand-trigger>
                <md-icon>keyboard_arrow_down</md-icon>
              </md-button>
            </md-card-actions>
  
            <md-card-content>
              <hello></hello>
            </md-card-content>
          </md-card-expand>
  
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
    sync() { },
  }
};
</script>
