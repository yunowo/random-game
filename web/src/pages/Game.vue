<template>
  <page-content page-title="随机">
    <div id="main-content">
      <div class="main-game" v-if="!isEmpty">
  
        <md-card md-with-hover>
          <md-toolbar class="md-transparent">
            <h1 class="md-title">参与者</h1>
          </md-toolbar>

          <md-card-content>
            <md-input-container>
              <label for="namelists">名单</label>
              <md-select id="namelists" v-model="selectedId">
                <md-option v-for="(item, index) in nameLists" :key="item.id" :value="item.id">{{item.title}}</md-option>
              </md-select>
            </md-input-container>
            <md-select id="names" multiple v-model="arraySelected">
              <md-button @click="selectAllTap">全选</md-button>
              <md-button @click="clearAllTap">清空</md-button>
              <md-option v-for="(item, index) in arrayFull" :key="index" :value="item">{{item}}</md-option>
            </md-select>
          </md-card-content>
  
          <div class="card-reservation">
            <md-subheader>选项</md-subheader>
            <md-button-toggle md-single class="md-button-group">
              <md-button class="md-toggle" @click="mode = 0">多次随机</md-button>
              <md-button @click="mode = 1">一次随机</md-button>
            </md-button-toggle>
          </div>
  
          <md-card-actions>
            <md-button class="md-raised md-warn" v-if="finished" @click="restartTap">重新开始</md-button>
            <md-button class="md-raised md-primary" v-if="!finished" @click="randomTap">随机</md-button>
          </md-card-actions>
  
        </md-card>
  
        <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog2">
          <md-dialog-title>Create new note</md-dialog-title>
          <md-dialog-content>
            <form>
              <md-checkbox id="my-test2" name="my-test2" v-model="finished" class="md-primary">Primary Color</md-checkbox>
            </form>
          </md-dialog-content>
          <md-dialog-actions>
            <md-button class="md-primary" @click="closeDialog('dialog2')">Cancel</md-button>
            <md-button class="md-primary" @click="closeDialog('dialog2')">Create</md-button>
          </md-dialog-actions>
        </md-dialog>
        <md-button class="md-fab md-fab-bottom-right" id="fab" @click="openDialog('dialog2')">
          <md-icon>add</md-icon>
        </md-button>
  
        <md-table-card md-with-hover>
          <md-toolbar>
            <h1 class="md-title">结果</h1>
            <md-button class="md-icon-button">
              <md-icon>filter_list</md-icon>
            </md-button>
  
            <md-button class="md-icon-button">
              <md-icon>search</md-icon>
            </md-button>
          </md-toolbar>
  
          <md-table v-once>
            <md-table-header>
              <md-table-row>
                <md-table-head>名字</md-table-head>
                <md-table-head md-numeric>结果</md-table-head>
              </md-table-row>
            </md-table-header>
  
            <md-table-body>
              <md-table-row :class="['generated', {'choosen': max === item}]" v-for="(item, index) in randomized" :key="index" v-if="item != -1">
                <md-table-cell>{{arraySelected[index]}}</md-table-cell>
                <md-table-cell md-numeric v-if="mode != 1">
                  <md-icon>message</md-icon>
                  <span>{{item}}</span>
                </md-table-cell>
              </md-table-row>
            </md-table-body>
          </md-table>
        </md-table-card>
      </div>
  
      <div class="page" v-if="isEmpty">
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

<style>
  .md-card {
    width: 400px;
    margin: 20px;
  }
  
  .main-game {
    display: flex;
    flex-flow: row wrap;
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
        nameLists: [],
        selectedId: '',
        selected: 0,
        arrayFull: [],
        arraySelected: [],
        randomized: [],
        max: 0,
        finished: false,
      };
    },
    computed: {
      isEmpty() {
        return this.nameLists.length === 0;
      },
      arraySelectedString() {
        return this.arraySelected.join(', ');
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
      listChange(e) {
        var selected = e.detail.value;
        var nameLists = this.nameLists;
        nameLists.forEach(function(e, i) {
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
        e.detail.value.forEach(function(e, i) {
          flags[e] = true;
        });
        this.arrayFlags = flags;
      },
      selectAllTap() {
        this.arraySelected = this.arrayFull;
      },
      clearAllTap() {
        this.arraySelected = [];
      },
      pickerChange(e) {
        this.mode = Number(e.detail.value);
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
          this.max = Math.max(...r);
          this.finished = r.length === this.arraySelected.length;
        } else if (this.mode === 1) {
          r = Array(this.arraySelected.length).fill(-1);
          r[Math.floor(Math.random() * this.arraySelected.length)] = 1;
          this.randomized = r;
          this.max = Math.max(...r);
          this.finished = true;
        }
      },
      restartTap() {
        this.randomized = [];
        this.max = 0;
        this.finished = false;
      },
    },
    mounted() {
      const selectedId = localStorage.getItem('selectedId');
      const nameLists = JSON.parse(localStorage.getItem('nameLists'));
  
      let selected = 0;
      nameLists.forEach(function(e, i) {
        if (e.id === selectedId) {
          selected = i;
          e.checked = true;
        }
      });
      const names = nameLists[selected].names;
  
      this.nameLists = nameLists;
      this.selectedId = selectedId;
      this.selected = selected;
      this.arrayFull = names;
      this.arraySelected = names;
    },
  };
</script>
