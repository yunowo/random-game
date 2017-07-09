//lists.js
var app = getApp()
var sha1 = require('../../utils/sha1.min.js')
var Base64 = require('../../utils/base64.modified')
Page({
  data: {
    nameLists: [],
    state: 0,
    b64: ''
  },
  itemTap: function (e) {
    wx.navigateTo({
      url: '../listedit/listedit?isNew=false&id=' + e.currentTarget.id,
    })
  },
  addTap: function () {
    wx.navigateTo({
      url: '../listedit/listedit?isNew=true&id=' + sha1(Date.now().toString())
    })
  },
  importTap: function () {
    this.setData({ state: 1 })
  },
  input: function (e) {
    this.setData({ b64: e.detail.value })
  },
  okTap: function () {
    var nameList = {}
    try {
      var str = this.toStr(this.data.b64)
      nameList = JSON.parse(str)
    } catch (e) {
      this.setData({ state: 0 })
      console.log(e)
      return
    }
    var dup = false
    var nameLists = this.data.nameLists
    nameLists.forEach(function (e, i) {
      if (e.id == nameList.id) {
        wx.showToast({
          title: '已有该名单',
          icon: 'success',
          duration: 2000
        })
        dup = true
      }
    })
    if (!dup) {
      nameLists.push(nameList)
      this.setData({
        state: 0,
        nameLists: nameLists
      })
      wx.setStorageSync('nameLists', nameLists)
      wx.showToast({
        title: '已保存 ' + nameList.title,
        icon: 'success',
        duration: 2000
      })
      this.onShow()
    } else {
      this.setData({ state: 0 })
    }
  },
  onShow: function () {
    var nameLists = wx.getStorageSync('nameLists')
    nameLists.forEach(function (e, i) {
      e.namesString = e.names.join(', ')
    })
    this.setData({
      nameLists: nameLists
    })
  },
  onLoad: function (options) {
    console.log(options)
    if (typeof options.b64 !== 'undefined') {
      try {
        this.setData({ b64: options.b64 })
        this.onShow()
        this.okTap()
      } catch (e) {
        console.log(e)
      }
    }
  },
  toStr: function (b64) {
    var str = Base64.decode(b64)
    console.log(str)
    return str
  }
})