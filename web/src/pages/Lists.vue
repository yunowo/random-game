<template>
  <page-content page-title="名单">
    <div id="main-content">
      <md-subheader class="info">
        <md-icon>info</md-icon>
        <div>离线状态下不可修改</div>
      </md-subheader>
      <md-list v-model="nameLists">
        <md-list-item v-for="(item, index) in nameLists" :key="item.id" @click="itemTap">
          <md-avatar class="md-avatar-icon md-primary">
            <md-icon>folder</md-icon>
          </md-avatar>
          <div class="md-list-text-container">
            <span>{{item.title}}</span>
            <span>{{item.names.join(', ')}}</span>
          </div>
          <md-button class="md-icon-button" @click="openDialog('dialog-share')">
            <md-icon>share</md-icon>
            <md-tooltip md-direction="top">分享</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog-create')">
            <md-icon>mode_edit</md-icon>
            <md-tooltip md-direction="top">编辑</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog-delete')">
            <md-icon>delete</md-icon>
            <md-tooltip md-direction="top">删除</md-tooltip>
          </md-button>
          <md-divider class="md-inset"></md-divider>
        </md-list-item>
      </md-list>
  
      <md-speed-dial md-open="hover" md-mode="scale" class="md-fab-bottom-right" md-theme="light-blue">
        <md-button id="fab" class="md-fab" md-fab-trigger @click="openDialog('dialog-create')">
          <md-icon md-icon-morph>mode_edit</md-icon>
          <md-icon>add</md-icon>
          <md-tooltip md-direction="left">添加</md-tooltip>
        </md-button>
  
        <md-button class="md-fab md-primary md-mini md-clean" @click="openDialog('dialog-import')">
          <md-icon>content_copy</md-icon>
          <md-tooltip md-direction="left">导入</md-tooltip>
        </md-button>
      </md-speed-dial>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog-import">
        <md-dialog-content>
          <h3 class="md-title">导入名单</h3>
          <md-tabs md-right class="md-transparent title-tabs">
            <md-tab id="tab-id" md-icon="insert_link" md-label="ID" md-active>
              <form>
                <md-input-container>
                  <label>ID</label>
                  <md-textarea maxlength="20" required v-model="nid"></md-textarea>
                </md-input-container>
                <div class="fork-select">
                  <md-icon>ac_unit</md-icon>
                  <md-button-toggle md-single class="md-button-group">
                    <md-button class="md-toggle" @click="fork = false">Clone</md-button>
                    <md-button @click="fork = true">Fork</md-button>
                  </md-button-toggle>
                </div>
              </form>
            </md-tab>
            <md-tab id="tab-base64" md-icon="content_paste" md-label="Base64">
              <form>
                <md-input-container>
                  <label>Base64</label>
                  <md-textarea maxlength="10000" required v-model="b64"></md-textarea>
                </md-input-container>
              </form>
            </md-tab>
          </md-tabs>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog-import')">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-import')">导入</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog-share">
        <md-dialog-content>
          <h3 class="md-title">分享名单</h3>
          <md-tabs md-right class="md-transparent title-tabs">
            <md-tab id="tab-id" md-icon="insert_link" md-label="ID" md-active>
              <form>
                <md-input-container>
                  <label>ID</label>
                  <md-textarea maxlength="20" required v-model="nid"></md-textarea>
                </md-input-container>
              </form>
            </md-tab>
            <md-tab id="tab-base64" md-icon="content_paste" md-label="Base64">
              <form>
                <md-input-container>
                  <label>Base64</label>
                  <md-textarea maxlength="10000" required v-model="b64"></md-textarea>
                </md-input-container>
              </form>
            </md-tab>
            <md-tab id="tab-qr" md-icon="select_all" md-label="QR">
  
            </md-tab>
          </md-tabs>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog-share')">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-share')">导入</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog-create">
        <md-dialog-title>Create new note</md-dialog-title>
        <md-dialog-content>
          <form>
            <md-input-container>
              <label>标题</label>
              <md-input required v-model="nameList.title"></md-input>
            </md-input-container>
            <div class="fork-select">
              <md-icon>security</md-icon>
              <md-button-toggle md-single class="md-button-group">
                <md-button @click="nameList.visibility = 0">公开</md-button>
                <md-button class="md-toggle" @click="nameList.visibility = 1">私有</md-button>
              </md-button-toggle>
            </div>
            <md-chips v-model="nameList.names" md-input-placeholder="Add a name"></md-chips>
          </form>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog-create')">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-create')">创建</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog-confirm md-title="确认删除?" md-content-html="真的要删除这个名单吗?" md-ok-text="确定" md-cancel-text="取消" @open="onOpen" @close="onClose" ref="dialog-delete">
      </md-dialog-confirm>
  
      <md-snackbar md-position="bottom center" ref="snackbar" :md-duration="4000">
        <span>{{message}}</span>
        <md-button class="md-accent" md-theme="light-blue" @click="$refs.snackbar.close()">Retry</md-button>
      </md-snackbar>
    </div>
  </page-content>
</template>

<style lang="scss">
.info {
  display: flex;
  align-items: center;
  .md-icon {
    margin: 8px;
    color: rgba(#000, .54) !important;
    flex: initial;
    width: 24px;
  }
}

.md-title {
  position: relative;
  z-index: 3;
  width: 300px;
  margin-top: 16px;
}

.title-tabs {
  margin-top: -72px;
}

.fork-select {
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
</style>

<script>
import Base64 from "js-base64";
import sha1 from "js-sha1";
export default {
  data() {
    return {
      nameLists: [],
      nid: 0,
      b64: '',
      fork: false,
      nameList: {
        title: '',
        visibility: 1,
        names: []
      },
      message: "",
    };
  },
  computed: {

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
    add() {
      axios.post('/namelists', {
        data: {
          data: this.nameList
        }
      })
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.log(error);
        });
    },
    import() {
      axios.post('/user/import', {
        params: {
          id: 0,
          fork: false,
        }
      })
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.log(error);
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
      nameLists.forEach((e, i) => {
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
    this.nameLists = nameLists;
  },
};
</script>
