<script lang="ts">
  import { NCard, NButton, NIcon } from 'naive-ui'
  import { CloseFilled } from '@vicons/material'
  import { deleteProject }  from './Workspace.vue'
  import { useRouter } from 'vue-router'


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
      Status : {
        type: Boolean,
        required: true
      }
    },
    components: {
      NCard,
      NButton,
      NIcon,
      CloseFilled
    },
    setup() {
      const router = useRouter()

      const onCardClick = (id: number) => {
        console.log(`Clicked on project ${id}`)
        router.push(`/project/${id}`)
      }

      return {
        deleteProject,
        onCardClick
      }
    }
  }
</script>

<template>
  <NCard hoverable @click="onCardClick(Id)">
    <div class="card">
      <span class="time">[{{ Time }}]&nbsp;</span>
      <span class="name">{{ Name }}</span>
      <span class="status-on" v-if="Status">(Активен)&nbsp;</span>
      <span class="status-off" v-else>(Завершен)&nbsp;</span>
      <n-button quaternary circle @click.stop="deleteProject(Id)">
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

.status-on {
  color: #00a854;
  align-self: center;
}

.status-off {
  color: #f04134;
  align-self: center;
}

</style>
