<template>
  <page-content page-title="随机">
    <div id="main-content">
      <div class="page__bd" v-if="!isEmpty">
  
        <div v-if="state === 1">
          <div class="weui-cells__title">名单</div>
          <div class="weui-cells weui-cells_after-title">
            <radio-group bindchange="listChange">
              <label class="weui-cell weui-check__label" v-for="(item, index) in nameLists">
              <radio class="weui-check" :value="index" :checked="item.checked" />
              <div class="weui-cell__bd">{{item.title}}</div>
              <div class="weui-cell__ft weui-cell__ft_in-radio" v-if="item.checked">
                <icon class="weui-icon-radio" type="success_no_circle" size="16"></icon>
              </div>
            </label>
            </radio-group>
          </div>
          <md-button class="md-raised md-primary" @click="OKTap">确定</md-button>
        </div>
  
        <div v-if="state === 2">
          <div class="weui-cells__title">参与者</div>
          <div class="weui-cells weui-cells_after-title">
            <checkbox-group bindchange="namesChange">
              <label class="weui-cell weui-check__label" v-for="(item, index) in arrayFull" :key="value">
              <checkbox class="weui-check" :value="index" :checked="arrayFlags[index]" />
    
              <div class="weui-cell__hd weui-check__hd_in-checkbox">
                <icon class="weui-icon-checkbox_circle" type="circle" size="23" v-if="!arrayFlags[index]"></icon>
                <icon class="weui-icon-checkbox_success" type="success" size="23" v-if="arrayFlags[index]"></icon>
              </div>
              <div class="weui-cell__bd">{{item}}</div>
            </label>
            </checkbox-group>
            <div class="weui-cell weui-cell_link">
              <div @click="selectAllTap" class="weui-cell__bd">全选</div>
            </div>
            <div class="weui-cell weui-cell_link">
              <div @click="clearAllTap" class="weui-cell__bd">清空</div>
            </div>
          </div>
          <div class="weui-btn-area">
            <md-button class="md-raised md-primary" @click="OKTap">确定</md-button>
          </div>
        </div>
  
        <div v-if="state === 0">
          <div class="weui-cells__title">参与者</div>
  
          <div class="weui-cells weui-cells_after-title">
            <div class="weui-cell weui-cell_access" hover-class="weui-cell_active" @click="selectListTap">
              <div class="weui-cell__bd longtext_bd">名单</div>
              <div class="weui-cell__ft weui-cell__ft_in-access longtext_ft">
                <text>{{nameLists[selected].title}}</text>
              </div>
            </div>
            <div class="weui-cell weui-cell_access" hover-class="weui-cell_active" @click="selectNamesTap">
              <div class="weui-cell__bd longtext_bd">参与者</div>
              <div class="weui-cell__ft weui-cell__ft_in-access longtext_ft">
                <text>{{arraySelectedString}}</text>
              </div>
            </div>
          </div>
  
          <div class="weui-cells__title">选项</div>
          <div class="weui-cells weui-cells_after-title">
            <div class="weui-cell weui-cell_select">
              <div class="weui-cell__hd weui-cell__hd_in-select-after">
                <div class="weui-label">模式</div>
              </div>
              <div class="weui-cell__bd">
                <picker bindchange="pickerChange" :value="index" :range="modes">
                  <div class="weui-select">{{modes[mode]}}</div>
                </picker>
              </div>
            </div>
          </div>
  
          <div class="weui-btn-area">
            <md-button class="md-raised md-primary" v-if="!finished" type="primary" size="default" @click="randomTap">随机</md-button>
          </div>
          <div class="generated-container">
            <div :class="[generated, {'choosen': max === item}]" v-for="(item, index) in randomized" v-if="item != -1" :key="key">
              <text class="item-left">{{arraySelected[index]}}</text>
              <text class="item-center" v-if="mode != 1"> {{item}}</text>
              <icon class="item-right" type="success" v-if="max === item" />
            </div>
          </div>
          <div class="weui-btn-area">
            <md-button class="md-raised md-warn" v-if="finished" type="warn" size="default" @click="restartTap">重新开始</md-button>
          </div>
        </div>
      </div>
  
      <div class="page" v-if="isEmpty">
        <div class="weui-msg">
          <div class="weui-msg__icon-area">
            <icon type="info" size="93"></icon>
          </div>
          <div class="weui-msg__text-area">
            <div class="weui-msg__title">无名单</div>
            <div class="weui-msg__desc">请添加或导入名单</div>
          </div>
          <div class="weui-msg__opr-area">
            <div class="weui-btn-area">
              <md-button class="md-raised md-primary" @click="emptyTap" type="primary">→_→</md-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </page-content>
</template>

<style>
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
        state: 0,
        mode: 0,
        modes: ['多次随机', '一次随机'],
        nameLists: [],
        isEmpty: false,
        selectedId: '',
        selected: 0,
        arrayFull: [],
        arraySelected: [],
        arraySelectedString: '',
        arrayFlags: [],
        randomized: [],
        max: 0,
        finished: false,
      };
    },
    methods: {
      listChange: function listChange(e) {
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
        this.arraySelectedString = names.join(', ');
        this.arrayFlags = Array(names.length).fill(true);
  
      },
      namesChange: function namesChange(e) {
        var flags = Array(this.arrayFull.length).fill(false);
        e.detail.value.forEach(function(e, i) {
          flags[e] = true;
        });
        this.arrayFlags = flags;
      },
      selectAllTap: function selectAllTap() {
        var flags = Array(this.arrayFull.length).fill(true);
        this.arrayFlags = flags;
      },
      clearAllTap: function clearAllTap() {
        var flags = Array(this.arrayFull.length).fill(false);
        this.arrayFlags = flags;
      },
      selectListTap: function selectListTap() {
        this.state = 1;
      },
      selectNamesTap: function selectNamesTap() {
        this.state = 2;
      },
      pickerChange: function pickerChange(e) {
        this.mode = Number(e.detail.value);
        this.restartTap();
      },
      OKTap: function OKTap() {
        if (this.state === 2) {
          var arraySelected = [];
          var that = this;
          this.arrayFlags.forEach(function(e, i) {
            if (e) {
              arraySelected.push(that.arrayFull[i]);
            }
          });
          if (arraySelected.length === 0) {
            return;
          }
  
          this.state = 0;
          this.arraySelected = arraySelected;
          this.arraySelectedString = arraySelected.join(', ');
  
        } else {
          this.state = 0;
        }
        this.restartTap();
        localStorage.setItem('selectedId', this.selectedId);
      },
      randomTap: function randomTap() {
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
      restartTap: function restartTap() {
        this.randomized = [];
        this.max = 0;
        this.finished = false;
      },
      emptyTap: function emptyTap() {
        wx.switchTab({
          url: '../lists/lists',
        });
      },
    },
    mounted: function mounted() {
      const selectedId = localStorage.getItem('selectedId');
      const nameLists = JSON.parse(localStorage.getItem('nameLists'));
      if (nameLists.length === 0) {
        this.isEmpty = true;
        return;
      }
      let selected = 0;
      console.log(nameLists)
      nameLists.forEach(function(e, i) {
        if (e.id === selectedId) {
          selected = i;
          e.checked = true;
        }
      });
      const names = nameLists[selected].names;
  
      this.nameLists = nameLists;
      this.isEmpty = false;
      this.selectedId = selectedId;
      this.selected = selected;
      this.arrayFull = names;
      this.arraySelected = names;
      this.arraySelectedString = names.join(', ');
      this.arrayFlags = Array(names.length).fill(true);
    },
  };
</script>
