<template>
  <page-content page-title="我的">
    <div id="main-content" class="main-my">
      <md-card md-with-hover id="card-my">
        <md-card-header>
          <div class="md-title">我的</div>
        </md-card-header>
        <div class="userinfo">
          <md-avatar class="md-large">
            <img :src="user.avatar" :alt="user.name">
          </md-avatar>
          <span class="md-subheading">{{user.name}}</span>
        </div>
        <div>
          <md-button class="md-raised md-primary" md-theme="green" @click="logout">
            <md-icon md-src="static/img/icon-github.svg"></md-icon>
            <div>Logout</div>
          </md-button>
          <md-button class="md-raised md-primary" md-theme="blue" @click="$store.commit('sync')">
            <md-icon>cloud_download</md-icon>
            <div>Sync with server</div>
          </md-button>
        </div>
      </md-card>
      <md-card md-with-hover id="card-about">
        <md-card-header>
          <div class="md-title">Random Game</div>
          <div class="md-subheading">Early Access 0.2.0</div>
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
            <about></about>
          </md-card-content>
        </md-card-expand>
  
      </md-card>
  
      <md-snackbar md-position="bottom center" ref="snackbar" :md-duration="4000">
        <span>{{message}}</span>
        <md-button class="md-accent" md-theme="blue" @click="$refs.snackbar.close()">Retry</md-button>
      </md-snackbar>
    </div>
  </page-content>
</template>

<style lang="scss">
.md-card {
  width: 400px;
  margin: 20px;
}

#card-my {
  max-height: 300px;
}

.md-with-hover {
  cursor: default !important;
}

.main-my {
  display: flex;
  justify-content: center;
  flex-flow: row wrap;
}

.userinfo {
  margin: 10px;
}
</style>

<script>
export default {
  data() {
    return {
      loggedin: true,
      message: '',
    };
  },
  computed: {
    user() {
      return this.$store.state.user;
    },
  },
  methods: {
    logout() {

    },

  },
  mounted() { },
};
</script>
