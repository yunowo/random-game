<template>
  <page-content page-title="名单">
    <div id="main-content">
      <md-subheader>名单</md-subheader>
      <md-list v-model="nameLists">
        <md-list-item v-for="(item, index) in nameLists" :key="item.id">
          <md-avatar class="md-avatar-icon md-primary">
            <md-icon>folder</md-icon>
          </md-avatar>
          <div class="md-list-text-container">
            <span>{{item.title}}</span>
            <span>{{item.namesString}}</span>
          </div>
          <md-button class="md-icon-button" @click="openDialog('dialog2')">
            <md-icon>share</md-icon>
            <md-tooltip md-direction="top">分享</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog2')">
            <md-icon>mode_edit</md-icon>
            <md-tooltip md-direction="top">编辑</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog2')">
            <md-icon>delete</md-icon>
            <md-tooltip md-direction="top">删除</md-tooltip>
          </md-button>
          <md-divider class="md-inset"></md-divider>
        </md-list-item>
      </md-list>
  
      <md-speed-dial md-open="hover" md-mode="scale" class="md-fab-bottom-right" md-theme="light-blue">
        <md-button id="fab" class="md-fab" md-fab-trigger @click="openDialog('dialog2')">
          <md-icon md-icon-morph>mode_edit</md-icon>
          <md-icon>add</md-icon>
          <md-tooltip md-direction="left">添加</md-tooltip>
        </md-button>
  
        <md-button class="md-fab md-primary md-mini md-clean" @click="openDialog('dialog1')">
          <md-icon>content_copy</md-icon>
          <md-tooltip md-direction="left">导入</md-tooltip>
        </md-button>
      </md-speed-dial>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog1">
        <md-dialog-title>导入名单</md-dialog-title>
        <md-dialog-content>
          <form>
            <md-input-container>
              <label>Base64</label>
              <md-textarea maxlength="10000" required v-model="b64"></md-textarea>
            </md-input-container>
          </form>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog1')">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog1')">导入</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog2">
        <md-dialog-title>Create new note</md-dialog-title>
        <md-dialog-content>
          <form>
            <md-input-container>
              <label>标题</label>
              <md-input required v-model="nameList.title"></md-input>
            </md-input-container>
            <md-input-container>
              <label for="visibility">可见性</label>
              <md-select id="visibility" v-model="nameList.visibility">
                <md-option :value="0">公开</md-option>
                <md-option :value="1">受保护</md-option>
                <md-option :value="2">私有</md-option>
              </md-select>
            </md-input-container>
            <md-chips v-model="nameList.names" md-input-placeholder="Add a name"></md-chips>
          </form>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog2')">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog2')">创建</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-snackbar md-position="bottom center" ref="snackbar" :md-duration="4000">
        <span>{{message}}</span>
        <md-button class="md-accent" md-theme="light-blue" @click="$refs.snackbar.close()">Retry</md-button>
      </md-snackbar>
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
      b64: '',
      nameList: {
        title: '',
        names: []
      },
      message: "",
    };
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
    itemTap(e) {
      wx.navigateTo({
        url: '../listedit/listedit?isNew=false&id=' + e.currentTarget.id,
      });
    },
    addTap() {
      wx.navigateTo({
        url: '../listedit/listedit?isNew=true&id=' + sha1(Date.now().toString()),
      });
    },
    okTap() {
      var nameList = {};
      try {
        var str = this.toStr(this.data.b64);
        nameList = JSON.parse(str);
      } catch (e) {
        console.log(e);
        return;
      }
      var dup = false;
      var nameLists = this.data.nameLists;
      nameLists.forEach(function (e, i) {
        if (e.id == nameList.id) {
          this.message = '已有该名单';
          this.$refs.snackbar.open();
          dup = true;
        }
      });
      if (!dup) {
        nameLists.push(nameList);
        this.nameLists = nameLists;
        localStorage.setItem('nameLists', JSON.stringify(nameLists));
        this.message = '已保存 ' + nameList.title;
        this.$refs.snackbar.open();
        this.onShow();
      }
    },
    onLoad(options) {
      console.log(options);
      if (typeof options.b64 !== 'undefined') {
        try {
          this.b64 = options.b64;
          this.onShow();
          this.okTap();
        } catch (e) {
          console.log(e);
        }
      }
    },
    toStr(b64) {
      var str = Base64.decode(b64);
      console.log(str);
      return str;
    },
  },
  mounted() {
    const nameLists = JSON.parse(localStorage.getItem('nameLists'));
    nameLists.forEach(function (e, i) {
      e.namesString = e.names.join(', ');
    });
    this.nameLists = nameLists;
  },
};
</script>
