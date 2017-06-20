// listedit.js
var app = getApp()
var Base64 = require('../../utils/base64.modified')
Page({
  data: {
    visibilities: ['公开', '受保护', '私有'],
    nameLists: [],
    nameList: {}
  },
  inputTitle: function (e) {
    var nameList = this.data.nameList
    nameList.title = e.detail.value
    this.setData({
      nameList: nameList
    })
  },
  inputName: function (e) {
    var nameList = this.data.nameList
    nameList.names[Number(e.currentTarget.id)] = e.detail.value
    this.setData({
      nameList: nameList
    })
  },
  deleteTap: function (e) {
    var nameList = this.data.nameList
    nameList.names.splice(Number(e.currentTarget.id), 1)
    this.setData({
      nameList: nameList
    })
  },
  addTap: function () {
    var nameList = this.data.nameList
    if (nameList.names[nameList.names.length - 1] != '') {
      nameList.names.push('')
    }
    this.setData({
      nameList: nameList
    })
  },
  saveTap: function () {
    var nameList = this.data.nameList
    nameList.names.filter(function (value) {
      return value != ''
    })
    var nameLists = this.data.nameLists
    if (!this.data.isNew) {
      nameLists.forEach(function (e, i) {
        if (e.id == nameList.id) {
          nameLists[i] = nameList
        }
      })
    } else {
      nameLists.push(nameList)
    }
    wx.setStorageSync('nameLists', nameLists)
    wx.navigateBack()
  },
  deleteAllTap: function () {
    var nameLists = this.data.nameLists
    var nameList = this.data.nameList
    nameLists.forEach(function (e, i) {
      if (e.id == nameList.id) {
        nameLists.splice(i, 1)
      }
    })
    wx.setStorageSync('nameLists', nameLists)
    wx.navigateBack()
  },
  clipboardTap: function () {
    var b64 = this.toBase64()
    wx.setClipboardData({
      data: b64,
      success: function (res) {
        wx.showToast({
          title: '已复制到剪贴板',
          icon: 'success',
          duration: 2000
        })
      }
    })
  },
  onLoad: function (options) {
    var nameLists = wx.getStorageSync('nameLists')
    var nameList = {}
    this.setData({
      nameLists: nameLists
    })

    var isNew = (options.isNew == 'true')
    var id = options.id
    if (isNew) {
      nameList = {
        id: id,
        visibility: 0,
        title: '',
        names: ['']
      }
    } else {
      nameLists.forEach(function (e, i) {
        if (e.id == id) {
          nameList = e
        }
      })
    }
    this.setData({
      nameList: nameList,
      isNew: isNew
    })
  },
  onReady: function () {

  },
  onShareAppMessage: function (res) {
    return {
      title: this.data.nameList.title,
      path: '/pages/lists/lists?b64=' + this.toBase64(),
      success: function (res) {
        console.log(res)
      },
      fail: function (res) {
        console.log(res)
      }
    }
  }, toBase64: function () {
    var b64 = Base64.encode(JSON.stringify(this.data.nameList))
    console.log(b64)
    return b64
  }
})