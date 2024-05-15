<script lang="ts">
  import { NButton, NIcon } from 'naive-ui'
  import { AddFilled } from '@vicons/material'

  import ProjectCard from './ProjectCard.vue'

  export default {
    components: {
      ProjectCard,
      NButton,
      NIcon,
      AddFilled
    }
  }
</script>

<template>
  <div class="workspace">
    <div class="cards">
      <ProjectCard -id="asdasaas" -name="Project 1" -time="2021-10-01 11:12"/>
      <ProjectCard -id="asdasaas" -name="Project 1" -time="2021-10-01 11:12"/>
    </div>
    <div class="new-button">
      <n-button secondary>
        <template #icon>
            <n-icon><add-filled/></n-icon>
        </template>
        Новый проект
      </n-button>
    </div>
  </div>
</template>

<style scoped>
.workspace {
  flex-grow: 1;
  overflow-y: auto;
}

.cards {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 16px;
}

.new-button {
  display: flex;
  justify-content: center;
}
</style>