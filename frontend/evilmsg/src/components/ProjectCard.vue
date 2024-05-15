<script lang="ts">
  import { NCard, NButton, NIcon } from 'naive-ui'
  import { CloseFilled } from '@vicons/material'

  export default {
    props: {
      Name: {
        type: String,
        required: true
      },
      Id: {
        type: String,
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
    }
  }
</script>

<template>
  <NCard>
    <div class="card">
      <span class="time">[{{ Time }}]&nbsp;</span>
      <span class="name">{{ Name }}</span>
      <span class="id">({{ Id }})&nbsp;</span>
      <n-button quaternary circle>
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
