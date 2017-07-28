<template>
  <page-content page-title="我的">
    <div id="main-content" class="main-my">
      <md-card md-with-hover id="card-my" v-if="auth">
        <md-card-header>
          <div class="md-title">我的</div>
        </md-card-header>
        <div class="userinfo">
          <md-avatar class="md-large">
            <img :src="user.avatar" :alt="user.name">
          </md-avatar>
          <span class="md-subheading">{{user.name}}</span>
  
          <div class="button-area">
            <md-button class="md-raised md-primary" md-theme="blue" @click="$store.dispatch('sync')">
              <div class="icon-button">
                <md-icon>cloud_download</md-icon>
                <div>同步</div>
              </div>
            </md-button>
            <md-button class="md-raised md-primary" md-theme="green" @click="logout">
              <div class="icon-button">
                <md-icon md-src="static/img/icon-github.svg"></md-icon>
                <div>退出</div>
              </div>
            </md-button>
          </div>
        </div>
      </md-card>
      <md-card md-with-hover id="card-about">
        <md-card-header>
          <div class="md-title">Random Game</div>
          <div class="md-subheading">Early Access 0.2.0</div>
        </md-card-header>
        <md-card-expand>
          <md-card-actions>
            <h2>Sources</h2>
            <md-button class="md-icon-button" md-expand-trigger>
              <md-icon>keyboard_arrow_down</md-icon>
            </md-button>
          </md-card-actions>
  
          <md-card-content>
            <about></about>
          </md-card-content>
        </md-card-expand>
  
      </md-card>
    </div>
  </page-content>
</template>

<style lang="scss">
.md-card {
  width: 400px;
  margin: 20px;
}

#card-my {
  max-height: 176px;
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
  display: flex;
  align-content: left;
  align-items: center;
  .md-avatar {
    margin: 20px;
  }
  .md-subheading {
    text-align: center;
    margin-right: 30px;
  }
  .button-area {
    display: flex;
    align-items: center;
    height: 64px;
  }

  .icon-button {
    display: flex;
    flex: 1;
    .md-icon {
      margin: auto 4px;
      color: white;
    }
    div {
      margin: auto 2px;
      color: white;
    }
  }
}
</style>

<script>
import { mapGetters } from 'vuex';

export default {
  data() {
    return {
      loggedin: true,
    };
  },
  computed: {
    ...mapGetters(['user', 'auth']),
  },
  methods: {
    logout() {
      document.cookie = 'rnd_session=;path=/random;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      this.$store.dispatch('sync');
    },
  },
  mounted() { },
};
</script>
