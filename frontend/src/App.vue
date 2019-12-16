<template>
  <div id="app">
    <umenu></umenu>
  </div>
</template>

<script>
import menu from "./components/Menu.vue"

export default {
  name: 'app',
  components: {
    'umenu': menu,
  }
}
</script>

<style>
</style>
