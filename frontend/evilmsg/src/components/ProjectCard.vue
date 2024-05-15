<script lang="ts">
  import { NCard, NButton, NIcon } from 'naive-ui'
  import { CloseFilled } from '@vicons/material'
  import { deleteProject }  from './Workspace.vue'

  export default {
    props: {
      Name: {
        type: String,
        required: true
      },
      Id: {
        type: Number,
        required: true
      },
      Time: {
        type: String,
        required: true
      },
    },
    components: {
      NCard,
      NButton,
      NIcon,
      CloseFilled
    },
    setup() {
      return {
        deleteProject
      }
    }
  }
</script>

<template>
  <NCard>
    <div class="card">
      <span class="time">[{{ Time }}]&nbsp;</span>
      <span class="name">{{ Name }}</span>
      <span class="id">({{ Id }})&nbsp;</span>
      <n-button quaternary circle @click="deleteProject(Id)">
        <template #icon>
          <n-icon><close-filled/></n-icon>
        </template>
      </n-button>
    </div>
  </NCard>
</template>

<style scoped>
.card {
  display: flex;
  justify-content: left;
}

.time {
  color: #999;
  align-self: center;
}

.name {
  font-weight: bold;
  color: #333;
  flex-grow: 1;
  align-self: center;
}

.id {
  color: #999;
  align-self: center;
}

</style>
