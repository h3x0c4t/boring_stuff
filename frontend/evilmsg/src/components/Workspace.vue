<script lang="ts">
  import { NButton, NIcon, NModal, NCard, NInput } from 'naive-ui'
  import { AddFilled } from '@vicons/material'

  import ProjectCard from './ProjectCard.vue'
  import { ref, onMounted } from 'vue'

  type Project = {
    id: number
    name: string
    status: boolean
    time_stopped: string
    time_started: string
  }

  var Projects = ref([] as Project[])
  var showModal =  ref(false)
  var projectName = ref('')

  const fetchProjects = async () => {
    const response = await fetch('/api/projects')
    const data = await response.json()
    Projects.value = data
    console.log(data)
  }

  export const deleteProject = async (id: number) => {
    await fetch(`/api/projects`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ id })
    })

    Projects.value = Projects.value.filter(project => project.id !== id)

    console.log(`Deleted project with id ${id}`)
  }

  const createProject = async (name: string) => {
    await fetch(`/api/projects`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ name })
    })

    projectName.value = ''
    showModal.value = false

    console.log(`Created project with name ${name}`)

    fetchProjects()
  }

  export default {
    components: {
      ProjectCard,
      NButton,
      NIcon,
      AddFilled,
      NModal,
      NCard,
      NInput
    },
    setup() {
      onMounted(() => {
        fetchProjects()
      })

      return {
        Projects,
        showModal,
        projectName,
        createProject
      }
    }
  }
</script>

<template>
  <div class="workspace">
    <div class="cards">
      <project-card v-for="project in Projects" :Id="project.id" :Name="project.name" :Time="project.time_started" :Status="project.status"/>
    </div>
    <div>
      <n-button secondary @click="showModal = true">
        <template #icon>
            <n-icon><add-filled/></n-icon>
        </template>
        Новый проект
      </n-button>
      <n-modal v-model:show="showModal">
        <n-card title="Новый проект" size="medium" style="min-width: 300px; max-width: 500px;">
          <n-input placeholder="Название проекта" v-model:value="projectName"/>
          <div class="buttons-dialog">
            <n-button type="primary" ghost style="flex-grow: 1;" @click="createProject(projectName)">Создать</n-button>
            <n-button type="error" ghost style="flex-grow: 1;" @click="showModal = false">Отмена</n-button>
          </div>
        </n-card>
      </n-modal>
    </div>
  </div>
</template>

<style scoped>
.workspace {
  flex-grow: 1;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.cards {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 16px;
  width: 80%;
}

.buttons-dialog {
  display: flex;
  height: 16px;
  gap: 16px;
  margin-top: 19px;
  margin-bottom: 20px;
}
</style>