<template>
  <page-content page-title="随机">
    <div id="main-content">
      <div class="main-game" v-if="!isEmpty">
  
        <md-card md-with-hover id="card-select">
          <md-toolbar class="md-transparent">
            <h1 class="md-title">参与者</h1>
          </md-toolbar>
  
          <md-card-content>
            <md-input-container>
              <label for="user.name_lists">名单</label>
              <md-select id="namelists" v-model="selectedId">
                <md-option v-for="(item, index) in user.name_lists" :key="item.id" :value="item.id">{{item.title}}</md-option>
              </md-select>
            </md-input-container>
            <md-input-container>
              <md-select id="names" multiple v-model="arraySelected">
                <div style="display: flex">
                  <md-button @click="selectAll">全选</md-button>
                  <md-button @click="clearAll">清空</md-button>
                </div>
                <md-option v-for="(item, index) in arrayFull" :key="index" :value="item">{{item}}</md-option>
              </md-select>
            </md-input-container>
  
            <h4 class="md-subheading">选项</h4>
            <div class="mode-select">
              <md-icon>toys</md-icon>
              <md-button-toggle md-single class="md-button-group">
                <md-button class="md-toggle" @click="modeChange(0)">多次随机</md-button>
                <md-button @click="modeChange(1)">一次随机</md-button>
              </md-button-toggle>
            </div>
          </md-card-content>
  
          <md-card-actions>
            <md-button class="md-warn" v-if="finished" @click="restartTap">重新开始</md-button>
            <md-button class="md-primary" v-else @click="randomTap">随机</md-button>
          </md-card-actions>
  
        </md-card>
  
        <md-table-card md-with-hover v-if="randomized.length !== 0">
          <md-card-header>
            <div class="md-title">结果</div>
          </md-card-header>
  
          <md-table>
            <md-table-header>
              <md-table-row>
                <md-table-head>名字</md-table-head>
                <md-table-head md-numeric>结果</md-table-head>
              </md-table-row>
            </md-table-header>
  
            <md-table-body>
              <md-table-row :class="['generated', {'choosen': max === item}]" v-for="(item, index) in randomized" :key="index" v-if="item !== -1">
                <md-table-cell>{{arraySelected[index]}}</md-table-cell>
                <md-table-cell md-numeric>
                  <span v-if="mode === 0">{{item}}</span>
                  <md-icon v-if="mode === 1">done</md-icon>
                </md-table-cell>
              </md-table-row>
            </md-table-body>
          </md-table>
        </md-table-card>
      </div>
  
      <div class="page" v-else>
        <div class="weui-msg">
          <md-icon class="md-accent md-size-4x"></md-icon>
          <div class="weui-msg__text-area">
            <div class="weui-msg__title">无名单</div>
            <div class="weui-msg__desc">请添加或导入名单</div>
          </div>
          <router-link to="lists">→_→</router-link>
        </div>
      </div>
    </div>
  </page-content>
</template>

<style lang="scss">
.md-card {
  width: 400px;
  margin: 20px;
}

#card-select {
  max-height: 368px;
}

.md-with-hover {
  cursor: default !important;
}

.main-game {
  display: flex;
  flex-flow: row wrap;
}

.mode-select {
  margin-top: 8px;
  display: flex;
  align-items: flex-start;
  justify-content: space-around;

  .md-icon {
    margin: 8px;
    color: rgba(#000, .54) !important;
    flex: none;
    width: 24px;
  }

  .md-button-group {
    flex: 1;
  }

  .md-button {
    border-radius: 2px !important;
    margin-left: 8px;
  }
}

.generated-container {
  display: flex;
  flex-direction: column;
  border: 1rpx solid #e0e0e0;
}

.generated {
  padding: 15rpx;
  font-size: 32rpx;
  align-items: center;
}

.choosen {
  background: #B9F6CA;
}
</style>

<script>
export default {
  data() {
    return {
      userInfo: {},
      mode: 0,
      user: { name_lists: [], },
      selectedId: 0,
      selected: 0,
      arrayFull: [],
      arraySelected: [],
      randomized: [],
      finished: false,
    };
  },
  computed: {
    isEmpty() {
      return this.user.name_lists.length === 0;
    },
    arraySelectedString() {
      return this.arraySelected.join(', ');
    },
    max() {
      return Math.max(...this.randomized);
    },
  },
  methods: {
    openDialog(ref) {
      this.$refs[ref].open();
    },
    closeDialog(ref) {
      this.$refs[ref].close();
    },
    onOpen() {
      console.log('Opened');
    },
    onClose(type) {
      console.log('Closed', type);
    },
    listChange(e) {//todo
      var selected = e.detail.value;
      var nameLists = this.user.name_lists;
      nameLists.forEach((e, i) => {
        e.checked = false;
      });
      nameLists[selected].checked = true;
      var selectedId = nameLists[selected].id;
      var names = nameLists[selected].names;

      this.selectedId = selectedId;
      this.selected = selected;
      this.nameLists = nameLists;
      this.arrayFull = names;
      this.arraySelected = names;
    },
    namesChange(e) {
      var flags = Array(this.arrayFull.length).fill(false);
      e.detail.value.forEach((e, i) => {
        flags[e] = true;
      });
      this.arrayFlags = flags;
    },
    selectAll() {
      this.arraySelected = this.arrayFull;
    },
    clearAll() {
      this.arraySelected = [];
    },
    modeChange(mode) {
      this.mode = mode;
      this.restartTap();
    },
    OKTap() {
      this.restartTap();
      localStorage.setItem('selectedId', this.selectedId);
    },
    randomTap() {
      var r = this.randomized;
      if (this.mode === 0) {
        if (r.length === this.arraySelected.length) {
          return;
        }
        r.push(Math.floor(Math.random() * 100));

        this.randomized = r;
        this.finished = r.length === this.arraySelected.length;
      } else if (this.mode === 1) {
        r = Array(this.arraySelected.length).fill(-1);
        r[Math.floor(Math.random() * this.arraySelected.length)] = 1;
        this.randomized = r;
        this.finished = true;
      }
    },
    restartTap() {
      this.randomized = [];
      this.finished = false;
    },
  },
  mounted() {
    const selectedId = localStorage.getItem('selectedId');
    const user = JSON.parse(localStorage.getItem('user'));

    let selected = 0;
    user.name_lists.forEach((e, i) => {
      if (e.id === selectedId) {
        selected = i;
        e.checked = true;
      }
    });
    const names = user.name_lists[selected].names;

    this.user = user;
    this.selectedId = selectedId;
    this.selected = selected;
    this.arrayFull = names;
    this.arraySelected = names;
  },
};
</script>
