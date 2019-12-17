<template>
  <div id="app">
    <umenu></umenu>
    <div class="mainContent">
      <lmenu /> 
    </div>
  </div>
</template>

<script>
import menu from "./components/Menu.vue"
import lm from "./components/LeftMenu.vue"

export default {
  name: 'app',
  components: {
    'umenu': menu,
    'lmenu': lm,
  }
}
</script>

<style scoped>
.mainContent {
  display: flex;
  width: 100%;
  height: 100%;
  min-height: 100%;
  margin-top: 60px;
}
</style>
