<template>
  <div class="left-menu">
    <div>
      <span>
        <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25">
          <path
            d="M426.666667 170.666667H170.666667c-47.36 0-85.333333 37.973333-85.333334 85.333333v512a85.333333 85.333333 0 0 0 85.333334 85.333333h682.666666a85.333333 85.333333 0 0 0 85.333334-85.333333V341.333333a85.333333 85.333333 0 0 0-85.333334-85.333333h-341.333333l-85.333333-85.333333z"
            fill
            p-id="20523"
          />
        </svg>
      </span>
      <span>我的文件</span>
    </div>
    <div>
      <span>
        <svg viewBox="0 0 1024 1024" version="1.1" width="25" height="25">
          <path
            d="M426.666667 170.666667l85.333333 85.333333h341.333333a85.333333 85.333333 0 0 1 85.333334 85.333333v426.666667a85.333333 85.333333 0 0 1-85.333334 85.333333H170.666667a85.333333 85.333333 0 0 1-85.333334-85.333333V256c0-47.36 37.973333-85.333333 85.333334-85.333333h256m213.333333 213.333333v128h-128v85.333333h128v128h85.333333v-128h128v-85.333333h-128V384h-85.333333z"
            fill
            p-id="20656"
          />
        </svg>
      </span>
      <span>新建文件夹</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "left-menu"
};
</script>

<style scoped>
.left-menu {
  width: 200px;
  height: 100%;
  min-height: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  box-shadow: 5px 5px 5px -5px rgba(0,0,0,0.25);
}
.left-menu > div {
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  font-weight: 600;
  cursor: pointer;
}
.left-menu > div:hover {
    background: #eee;
}
.left-menu > div > span:nth-child(2) {
    margin-left: 5px;
}
.left-menu > div > span:first-child {
    text-indent: 40px;
}
</style>