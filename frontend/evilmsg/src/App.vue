<script lang="ts">
import Header from './components/Header.vue'
import Workspace from './components/Workspace.vue'
import Footer from './components/Footer.vue'

export default {
  components: {
    Header,
    Workspace,
    Footer
  }
}
</script>

<template>
  <Header/>
  <router-view/>
</template>

<style>
html, body {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

#app {
  height: 100%;
  display: flex;
  flex-direction: column;
}
</style>
