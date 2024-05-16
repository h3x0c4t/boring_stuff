<script lang="ts">
import Header from './components/Header.vue'
import Workspace from './components/Workspace.vue'
import Footer from './components/Footer.vue'

import hljs from 'highlight.js/lib/core'
import json from 'highlight.js/lib/languages/json'

import { NConfigProvider } from 'naive-ui'

hljs.registerLanguage('json', json)

export default {
  components: {
    Header,
    Workspace,
    Footer,
    NConfigProvider
  },
  setup() {
    return {
      hljs
    }
  }
}
</script>

<template>
  <n-config-provider :hljs="hljs" class="config">
    <Header/>
    <router-view/>
  </n-config-provider>
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
}

.config {
  height: 100%;
  display: flex;
  flex-direction: column;
}
</style>
