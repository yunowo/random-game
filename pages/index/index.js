//index.js
var app = getApp()
Page({
  data: {
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
    finished: false
  },
  listChange: function (e) {
    var selected = e.detail.value
    var nameLists = this.data.nameLists
    nameLists.forEach(function (e, i) {
      e.checked = false
    })
    nameLists[selected].checked = true
    var selectedId = nameLists[selected].id
    var names = nameLists[selected].names
    this.setData({
      selectedId: selectedId,
      selected: selected,
      nameLists: nameLists,
      arrayFull: names,
      arraySelected: names,
      arraySelectedString: names.join(', '),
      arrayFlags: Array(names.length).fill(true)
    })
  },
  namesChange: function (e) {
    var flags = Array(this.data.arrayFull.length).fill(false)
    e.detail.value.forEach(function (e, i) {
      flags[e] = true
    })
    this.setData({
      arrayFlags: flags
    })
  },
  selectAllTap: function () {
    var flags = Array(this.data.arrayFull.length).fill(true)
    this.setData({
      arrayFlags: flags
    })
  },
  clearAllTap: function () {
    var flags = Array(this.data.arrayFull.length).fill(false)
    this.setData({
      arrayFlags: flags
    })
  },
  selectListTap: function () {
    this.setData({ state: 1 })
  },
  selectNamesTap: function () {
    this.setData({ state: 2 })
  },
  pickerChange: function (e) {
    this.setData({ mode: Number(e.detail.value) })
    this.restartTap()
  },
  OKTap: function () {
    if (this.data.state == 2) {
      var arraySelected = []
      var that = this
      this.data.arrayFlags.forEach(function (e, i) {
        if (e) { arraySelected.push(that.data.arrayFull[i]) }
      })
      if (arraySelected.length == 0) { return }
      this.setData({
        state: 0,
        arraySelected: arraySelected,
        arraySelectedString: arraySelected.join(', '),
      })
    } else { this.setData({ state: 0 }) }
    this.restartTap()
    wx.setStorageSync('selectedId', this.data.selectedId)
  },
  randomTap: function () {
    var r = this.data.randomized
    if (this.data.mode == 0) {
      if (r.length == this.data.arraySelected.length) {
        return
      }
      r.push(Math.floor(Math.random() * 100))
      this.setData({ randomized: r, max: Math.max(...r), finished: r.length == this.data.arraySelected.length })
    }
    else if (this.data.mode == 1) {
      r = Array(this.data.arraySelected.length).fill(-1)
      r[Math.floor(Math.random() * this.data.arraySelected.length)] = 1
      this.setData({ randomized: r, max: Math.max(...r), finished: true })
    }
  },
  restartTap: function () {
    this.setData({ randomized: [], max: 0, finished: false })
  },
  emptyTap: function () {
    wx.switchTab({
      url: '../lists/lists'
    })
  },
  onShow: function (options) {
    var selectedId = wx.getStorageSync('selectedId')
    var nameLists = wx.getStorageSync('nameLists')
    if (nameLists.length == 0) {
      this.setData({ isEmpty: true })
      return
    }
    var selected = 0
    nameLists.forEach(function (e, i) {
      if (e.id == selectedId) {
        selected = i
        e.checked = true
      }
    })
    var names = nameLists[selected].names
    this.setData({
      nameLists: nameLists,
      isEmpty: false,
      selectedId: selectedId,
      selected: selected,
      arrayFull: names,
      arraySelected: names,
      arraySelectedString: names.join(', '),
      arrayFlags: Array(names.length).fill(true)
    })
  },
})
