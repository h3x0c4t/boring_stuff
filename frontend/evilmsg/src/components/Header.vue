<script lang="ts">
  import { useRouter } from 'vue-router'

  export default {
    setup() {
      const router = useRouter()

      const onLogoClick = () => {
        console.log(`Clicked on logo`)
        router.push(`/`)
      }

      return {
        onLogoClick
      }
    }
  }
</script>

<template>
  <div class="header">
    <span class="logo-text" @click="onLogoClick()">Фишинг хуйня</span>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: left;
  min-height: 48px;
  background-color: #f0f0f0;
  font-size: 20px;
  font-weight: bold;
  color: #333;
}

.logo-text {
  margin-left: 18px;
  font-size: 20px;
  font-weight: bold;
  color: #333;
  align-self: center;
}
</style>