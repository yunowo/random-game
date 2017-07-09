<template>
   <div id="main">
    <div class="page__bd">
      <div class="weui-cells__title">标题</div>
      <div class="weui-cells weui-cells_after-title">
        <div class="weui-cell weui-cell_input">
          <div class="weui-cell__bd">
            <input class="weui-input" bindinput="inputTitle" :value="nameList.title" placeholder="标题" />
          </div>
        </div>
      </div>

      <!--<picker mode="selector" value="{{nameList.visibility}}" range="{{visibilities}}">
  <div class="picker">{{visibilities[nameList.visibility]}}</div>
</picker>-->

      <div class="weui-cells__title">名单</div>
      <div class="weui-cells weui-cells_after-title">
        <div class="weui-cell weui-cell_input nameitem" wx:for="{{nameList.names}}">
          <div class="weui-cell__bd">
            <input :id="index" bindinput="inputName" class="weui-input" :value="item"></input>
          </div>
          <div class="weui-cell__ft nameitem">
            <button :id="index" class="weui-btn mini-btn" type="default" size="mini" bindtap="deleteTap">删除</button>
          </div>
        </div>
        <div class="weui-cell weui-cell_link">
          <div @click="addTap" class="weui-cell__bd">添加</div>
        </div>
      </div>
      <div class="weui-btn-area">
        <button @click="saveTap" type="primary">保存</button>
      </div>
      <div class="weui-btn-area">
        <button v-if="!isNew" @click="deleteAllTap" type="warn">删除</button>
      </div>
      <div class="weui-btn-area">
        <button v-if="!isNew" open-type="share">分享</button>
      </div>
      <div class="weui-btn-area">
        <button v-if="!isNew" @click="clipboardTap">复制到剪贴板</button>
      </div>
    </div>
  </div>
</template>
<style>

</style>
<script>
export default {
  data() {
      return {
        visibilities: ['公开', '受保护', '私有'],
        nameLists: [],
        nameList: {},
      };
  },
  methods: {
      inputTitle: function inputTitle(e) {
        var nameList = this.data.nameList;
        nameList.title = e.detail.value;
        this.setData({
          nameList: nameList,
        });
      },
      inputName: function inputName(e) {
        var nameList = this.data.nameList;
        nameList.names[Number(e.currentTarget.id)] = e.detail.value;
        this.setData({
          nameList: nameList,
        });
      },
      deleteTap: function deleteTap(e) {
        var nameList = this.data.nameList;
        nameList.names.splice(Number(e.currentTarget.id), 1);
        this.setData({
          nameList: nameList,
        });
      },
      addTap: function addTap() {
        var nameList = this.data.nameList;
        if (nameList.names[nameList.names.length - 1] != '') {
          nameList.names.push('');
        }
        this.setData({
          nameList: nameList,
        });
      },
      saveTap: function saveTap() {
        var nameList = this.data.nameList;
        nameList.names.filter(function (value) {
          return value != '';
        });
        var nameLists = this.data.nameLists;
        if (!this.data.isNew) {
          nameLists.forEach(function (e, i) {
            if (e.id == nameList.id) {
              nameLists[i] = nameList;
            }
          });
        } else {
          nameLists.push(nameList);
        }
        wx.setStorageSync('nameLists', nameLists);
        wx.navigateBack();
      },
      deleteAllTap: function deleteAllTap () {
        var nameLists = this.data.nameLists;
        var nameList = this.data.nameList;
        nameLists.forEach(function (e, i) {
          if (e.id == nameList.id) {
            nameLists.splice(i, 1);
          }
        });
        wx.setStorageSync('nameLists', nameLists);
        wx.navigateBack();
      },
      clipboardTap: function clipboardTap() {
        var b64 = this.toBase64();
        wx.setClipboardData({
          data: b64,
          success: function (res) {
            wx.showToast({
              title: '已复制到剪贴板',
              icon: 'success',
              duration: 2000,
            });
          },
        });
      },
      onLoad: function onLoad(options) {
        var nameLists = wx.getStorageSync('nameLists');
        var nameList = {};
        this.setData({
          nameLists: nameLists,
        });

        var isNew = (options.isNew === 'true');
        var id = options.id;
        if (isNew) {
          nameList = {
            id: id,
            visibility: 0,
            title: '',
            names: [''],
          };
        } else {
          nameLists.forEach(function (e, i) {
            if (e.id === id) {
              nameList = e;
            }
          });
        }
        this.setData({
          nameList: nameList,
          isNew: isNew,
        });
      },
      onReady: function onReady() {

      },
      onShareAppMessage: function onShareAppMessage(res) {
        return {
          title: this.data.nameList.title,
          path: '/pages/lists/lists?b64=' + this.toBase64(),
          success: function (res) {
            console.log(res);
          },
          fail: function (res) {
            console.log(res);
          },
        };
      },
      toBase64: function toBase64() {
        var b64 = Base64.encode(JSON.stringify(this.data.nameList));
        console.log(b64);
        return b64;
      },
    },
};
</script>
