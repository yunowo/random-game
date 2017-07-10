<template>
  <page-content page-title="名单">
    <div id="main-content">
      <div class="page__bd">
        <div v-if="state == 0">
          <div class="weui-cells weui-cells_after-title">
            <div v-for="(item, index) in nameLists" :id="item.id" class="weui-cell" @click="itemTap">
              <div class="weui-cell__bd longtext_bd">{{item.title}}</div>
              <div class="weui-cell__ft longtext_ft">
                <text>{{item.namesString}}</text>
              </div>
            </div>
          </div>
          <md-button class="md-raised md-primary" bindtap="addTap">添加</md-button>
          <md-button class="md-raised md-primary" bindtap="importTap">导入</md-button>
        </div>
  
        <div v-if="state == 1">
          <div class="weui-cells__title">导入</div>
          <div class="weui-cells weui-cells_after-title">
            <div class="weui-cell">
              <div class="weui-cell__bd">
                <!-- <textarea class="weui-textarea" bindinput="input" maxlength="-1" auto-height placeholder="Base64" /> -->
              </div>
            </div>
          </div>
          <div class="weui-btn-area">
            <md-button class="md-raised md-primary" bindtap="okTap">确定</md-button>
          </div>
        </div>
      </div>
    </div>
  </page-content>
</template>

<style>
  
</style>

<script>
  import Base64 from "js-base64";
  import sha1 from "js-sha1";
  export default {
    data() {
      return {
        nameLists: [],
        state: 0,
        b64: '',
      };
    },
    methods: {
      itemTap: function itemTap(e) {
        wx.navigateTo({
          url: '../listedit/listedit?isNew=false&id=' + e.currentTarget.id,
        });
      },
      addTap: function addTap() {
        wx.navigateTo({
          url: '../listedit/listedit?isNew=true&id=' + sha1(Date.now().toString()),
        });
      },
      importTap: function importTap() {
        this.setData({
          state: 1,
        });
      },
      input: function input(e) {
        this.setData({
          b64: e.detail.value,
        });
      },
      okTap: function okTap() {
        var nameList = {};
        try {
          var str = this.toStr(this.data.b64);
          nameList = JSON.parse(str);
        } catch (e) {
          this.setData({
            state: 0,
          });
          console.log(e);
          return;
        }
        var dup = false;
        var nameLists = this.data.nameLists;
        nameLists.forEach(function(e, i) {
          if (e.id == nameList.id) {
            wx.showToast({
              title: '已有该名单',
              icon: 'success',
              duration: 2000,
            });
            dup = true;
          }
        });
        if (!dup) {
          nameLists.push(nameList);
          this.setData({
            state: 0,
            nameLists: nameLists,
          });
          localStorage.setItem('nameLists', nameLists);
          wx.showToast({
            title: '已保存 ' + nameList.title,
            icon: 'success',
            duration: 2000,
          });
          this.onShow();
        } else {
          this.setData({
            state: 0,
          });
        }
      },
      onLoad: function onLoad(options) {
        console.log(options);
        if (typeof options.b64 !== 'undefined') {
          try {
            this.setData({
              b64: options.b64,
            });
            this.onShow();
            this.okTap();
          } catch (e) {
            console.log(e);
          }
        }
      },
      toStr: function toStr(b64) {
        var str = Base64.decode(b64);
        console.log(str);
        return str;
      },
    },
    mounted: function mounted() {
        const nameLists = JSON.parse(localStorage.getItem('nameLists'));
        nameLists.forEach(function(e, i) {
          e.namesString = e.names.join(', ');
        });
        this.setData({
          nameLists: nameLists,
        });
      },
  };
</script>
