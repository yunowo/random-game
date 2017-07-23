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
            <md-icon>ac_unit</md-icon>
          </md-avatar>
          <div class="md-list-text-container">
            <span>{{item.title}}</span>
            <span>{{item.names.join(', ')}}</span>
          </div>
          <md-button class="md-icon-button" @click="openDialog('dialog-create', item, true)">
            <md-icon>mode_edit</md-icon>
            <md-tooltip md-direction="top">编辑</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog-share', item)">
            <md-icon>share</md-icon>
            <md-tooltip md-direction="top">分享</md-tooltip>
          </md-button>
          <md-button class="md-icon-button" @click="openDialog('dialog-delete', item)">
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
        <md-dialog-content class="dense">
          <h3 class="md-title">导入名单</h3>
          <md-tabs md-right class="md-transparent title-tabs" ref="import-tabs">
            <md-tab id="tab-id" md-icon="insert_link" md-label="ID" md-active>
              <form>
                <md-input-container md-clearable>
                  <label>ID</label>
                  <md-input maxlength="20" required v-model="nameList.id"></md-input>
                </md-input-container>
                <div class="fork-select">
                  <md-icon v-if="fork === false">cloud</md-icon>
                  <md-icon v-else>cloud_queue</md-icon>
                  <md-button-toggle md-single class="md-button-group">
                    <md-button class="md-toggle" @click="fork = false">Clone</md-button>
                    <md-button @click="fork = true">Fork</md-button>
                  </md-button-toggle>
                </div>
              </form>
            </md-tab>
            <md-tab id="tab-base64" md-icon="content_paste" md-label="Base64">
              <form>
                <md-input-container md-clearable>
                  <label>Base64</label>
                  <md-textarea maxlength="10000" required v-model="b64"></md-textarea>
                </md-input-container>
              </form>
            </md-tab>
          </md-tabs>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog-import', false)">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-import', true)">导入</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog-share">
        <md-dialog-content class="dense">
          <h3 class="md-title">分享名单</h3>
          <md-tabs md-right class="md-transparent title-tabs">
            <md-tab id="tab-id" md-icon="insert_link" md-label="ID" md-active>
              <form>
                <md-input-container>
                  <label>ID</label>
                  <md-input maxlength="20" required v-model="nameList.id"></md-input>
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
              <vue-qr id="qrcode" bgSrc="/static/img/qr_background.jpg" text="Hello world!" height="200" width="200" v-if="mode === 'share'"></vue-qr>
            </md-tab>
          </md-tabs>
        </md-dialog-content>
        <md-dialog-actions>
          <md-button class="md-primary" @click="closeDialog('dialog-share', false)">关闭</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-share', true)">复制到剪贴板</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog md-open-from="#fab" md-close-to="#fab" ref="dialog-create">
        <md-dialog-title class="dense">{{mode === 'create' ? '创建新名单' : '编辑名单'}}</md-dialog-title>
        <md-dialog-content class="dense">
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
          <md-button class="md-primary" @click="closeDialog('dialog-create', false)">取消</md-button>
          <md-button class="md-primary" @click="closeDialog('dialog-create', true)">创建</md-button>
        </md-dialog-actions>
      </md-dialog>
  
      <md-dialog-confirm md-title="确认删除?" md-content-html="真的要删除这个名单吗?" md-ok-text="确定" md-cancel-text="取消" @open="remove" ref="dialog-delete">
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

.md-dialog {
  max-width: 360px;
}

.md-dialog-title.dense {
  margin-bottom: 0px !important;
}

.md-dialog-content.dense {
  padding-bottom: 0px !important;
}

#qrcode {
  display: flex;
  justify-content: center;
}
</style>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      nameLists: [],
      b64: '',
      fork: false,
      nameList: {},
      newNameList: {
        id: null,
        title: '',
        visibility: 1,
        names: []
      },
      message: "",
      mode: 'create',
    };
  },
  computed: {

  },
  methods: {
    openDialog(ref, nameList, edit) {
      this.$refs[ref].open();
      if (ref === 'dialog-share') {
        this.mode = 'share';
        this.nameList = nameList;
        this.b64 = this.tob64();
      }
      else if (ref === 'dialog-create' && edit) {
        this.mode = 'edit';
        this.nameList = nameList;
      } else if (ref === 'dialog-create' && !edit) {
        this.mode = 'create';
        this.nameList = this.newNameList;
        this.b64 = "";
      } else if (ref === 'dialog-import') {
        this.mode = 'import';
        this.nameList = this.newNameList;
        this.b64 = "";
      }
    },
    closeDialog(ref, ok) {
      this.$refs[ref].close();
      if (!ok) return;
      switch (this.mode) {
        case 'create': {
          this.add();
          break;
        }
        case 'edit': {
          this.add();
          break;
        }
        case 'share': {

          break;
        }
        case 'import': {
          this.import();
          break;
        }
      }
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
      }).then(response => {
        console.log(response);
        this.refresh();
      }).catch(error => {
        console.log(error);
      });
    },
    import() {
      let tab = this.$refs['import-tabs'].activeTabNumber;
      if (tab === 1) {
        let json = JSON.parse(this.b64tostr(this.data.b64));
        this.nameList = json;
      }

      let dup = false;
      this.nameLists.forEach((e, i) => {
        if (e.id == this.nameList.id) {
          this.message = '已有该名单';
          this.$refs.snackbar.open();
          dup = true;
        }
      });
      if (dup) return;

      if (tab === 0) {
        axios.post('/user/import', {
          params: {
            id: this.nameList.id,
            fork: this.fork,
          }
        }).then(response => {
          console.log(response);
          this.refresh();
        }).catch(error => {
          console.log(error);
        });
      } else {
        this.nameList.id = null;
        this.add();
      }
    },
    remove() {
      axios.post('/user/remove', {
        params: {
          id: this.nameList.id,
        }
      }).then(response => {
        console.log(response);
        this.refresh();
      }).catch(error => {
        console.log(error);
      });
    },
    refresh() {
      axios.get('/user').then(response => {
        console.log(response);
        localStorage.setItem('nameLists', response.data);
      }).catch(error => {
        console.log(error);
      });
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
    b64tostr(b64) {
      let str = atob(b64);
      console.log(str);
      return str;
    },
    tob64() {
      let b64 = btoa(JSON.stringify(this.nameList));
      console.log(b64);
      return b64;
    },
  },
  mounted() {
    const nameLists = JSON.parse(localStorage.getItem('nameLists'));
    this.nameLists = nameLists;
    this.nameList = this.newNameList;
  },
};
</script>
