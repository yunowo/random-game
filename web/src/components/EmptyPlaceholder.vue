<template>
  <div class="empty">
    <div class="icon-circle md-primary">
      <md-ink-ripple />
      <img src="/static/img/ic_nothing.png"></img>
      <div class="md-headline">无名单</div>
      <div class="md-subheading">创建或导入一个</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
};
</script>

<style lang="scss">
.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  .md-headline {
    font-size: 26px;
    margin: 10px;
  }
  .icon-circle {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 400px;
    height: 400px;
    margin: 40px;
    background-color: #e1f5fe;
    border-radius: 50%;
    box-shadow: inset 0 0 2px rgba(0, 0, 0, .12), inset 0 2px 4px rgba(0, 0, 0, .24);
  }
}
</style>
