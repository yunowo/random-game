<template>
  <div class="container">
    <md-progress md-indeterminate class="progress md-accent" v-if="loading"></md-progress>
    <md-sidenav class="main-sidebar md-left md-fixed" md-swipeable ref="main-sidebar">
      <md-toolbar>
        <router-link exact to="/">
          <md-icon class="md-size-5x md-primary">casino</md-icon>
          <span class="black">随机游戏</span>
        </router-link>
      </md-toolbar>
  
      <div class="main-sidebar-links">
        <md-list class="md-dense">
          <md-list-item>
            <router-link exact to="/random/">随机</router-link>
          </md-list-item>
  
          <md-list-item>
            <router-link exact to="/random/lists">名单</router-link>
          </md-list-item>
  
          <md-list-item>
            <router-link exact to="/random/my">我的</router-link>
          </md-list-item>
        </md-list>
      </div>
  
    </md-sidenav>
  
    <transition name="md-router" appear>
      <router-view></router-view>
    </transition>
  
    <login-dialog v-if="!auth" ref="login"></login-dialog>
  
    <md-snackbar md-position="bottom center" ref="snackbar" :md-duration="4000">
      <span>{{message.text}}</span>
      <md-button class="md-accent" md-theme="blue" @click="$refs.snackbar.close()">确定</md-button>
    </md-snackbar>
  </div>
</template>

<style lang="scss">
@import "style/variables.scss";
$sizebar-size: 280px;
[v-cloak] {
  display: none;
}

html,
body {
  height: 100%;
  overflow: hidden;
}

body {
  display: flex;
}

.container {
  min-height: 100%;
  display: flex;
  flex-flow: column nowrap;
  flex: 1;
  transition: $swift-ease-out;
  @media (min-width: 1281px) {
    padding-left: $sizebar-size;
  }
}

.main-sidebar.md-sidenav {
  .md-sidenav-content {
    width: $sizebar-size;
    display: flex;
    flex-flow: column;
    overflow: hidden;
    @media (min-width: 1281px) {
      top: 0;
      pointer-events: auto;
      transform: translate3d(0, 0, 0) !important;
      box-shadow: $material-shadow-2dp;
    }
  }
  .md-backdrop {
    @media (min-width: 1281px) {
      opacity: 0;
      pointer-events: none;
    }
  }
  .md-toolbar {
    min-height: 172px;
    border-bottom: 1px solid rgba(#000, .12);
    background-color: #fff;
    font-size: 24px;
    a {
      width: 100%;
      display: flex;
      flex-flow: column;
      justify-content: center;
      align-items: center;
      color: inherit;
      text-decoration: none;
      user-select: none;
      &:hover {
        color: inherit;
        text-decoration: none;
      }
    }
    img {
      width: 160px;
      margin-bottom: 16px;
    }
  }
  .main-sidebar-links {
    overflow: auto;
    flex: 1;
    .md-inset .md-list-item-container {
      padding-left: 36px;
    }
    .md-list-item-container {
      font-size: 14px;
      font-weight: 500;
    }
  }
}

.main-content {
  padding: 16px;
  flex: 1;
  overflow: auto;
  background-color: #fff;
  transform: translate3D(0, 0, 0);
  transition: $swift-ease-out;
  transition-delay: .2s;
}

.md-router-enter,
.md-router-leave {
  position: absolute;
  top: 0;
  right: 0;
  left: 0;
  @media (min-width: 1281px) {
    left: $sizebar-size;
  }
  .main-content {
    opacity: 0;
    overflow: hidden;
  }
}

.md-router-leave {
  z-index: 1;
  transition: $swift-ease-in;
  transition-duration: .25s;
}

.md-router-enter {
  z-index: 2;
  transition: $swift-ease-out;
  .main-content {
    transform: translate3D(0, 10%, 0);
  }
}

code {
  &:not(.hljs) {
    margin-left: 1px;
    margin-right: 1px;
    padding: 0 4px;
    display: inline-block;
    border-radius: 2px;
    font-family: "Operator Mono", "Fira Code", Menlo, Hack, "Roboto Mono", "Liberation Mono", Monaco, monospace;
    pre {
      margin: 8px 0;
    }
  }
}

.phone-viewport {
  width: 360px;
  height: 540px;
  margin-right: 16px;
  display: inline-block;
  position: relative;
  overflow: hidden;
  background-color: #fff;
  border: 1px solid rgba(#000, .12);
}

.api-table tr>td:first-child {
  white-space: nowrap;
}

.progress {
  position: fixed;
  top: 0px;
  left: 0px;
  z-index: 1000;
  box-shadow: 0 0 10px rgba(128, 128, 128, 0.7);
}

.black {
  color: black;
}
</style>

<script>
import { mapGetters } from 'vuex';

export default {
  data() {
    return {
      toolbar: true,
      pageTitle: '',
    };
  },
  computed: {
    ...mapGetters(['loading', 'auth', 'message']),
  },
  watch: {
    auth(a) {
      this.showLogin(a);
    },
    message() {
      this.$refs.snackbar.open();
    },
  },
  methods: {
    toggleSidenav() {
      this.$refs['main-sidebar'].toggle();
    },
    closeSidenav() {
      this.$refs['main-sidebar'].close();
    },
    showLogin(a) {
      if (!a) {
        setTimeout(() => {
          this.$refs.login.$refs.dialog.open();
        }, 100);
      }
    },
  },
  mounted() {
    this.$store.dispatch('sync');
    this.showLogin(this.auth);
  },
};
</script>
