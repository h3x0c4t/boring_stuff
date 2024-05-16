<script lang="ts">
  import { NCard, NButton, NInputGroup, NSelect } from 'naive-ui'
  import { useRoute } from 'vue-router'
  import { ref, onMounted } from 'vue'

  type Project = {
    id: number
    name: string
    status: boolean
    time_stopped: string
    time_started: string
  }

  var CurProject = ref({} as Project)

  const fetchProject = async (id: string) => {
    const response = await fetch(`http://localhost:3000/api/projects/${id}`)
    const data = await response.json()
    CurProject.value = data
    console.log(data)
  }

  const stopProject = async (id: number) => {
    await fetch(`http://localhost:3000/api/projects/${id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ status: 0 })
    })

    fetchProject(id.toString())

    console.log(`Stopped project with id ${id}`)
  }

  export default {
    components: {
      NCard,
      NButton,
      NInputGroup,
      NSelect,
    },
    setup() {
      const route = useRoute()
      const id = route.params.id as string

      onMounted(() => {
        fetchProject(id)
      })

      return {
        CurProject,
        stopProject
      }
    }
  }
</script>

<template>
    <div class="project">
      <n-card class="contol-menu">
        <div class="panel">
          <div class="info">
            <div class="project-name">{{ CurProject.name }}</div>
            <div class="description">Начало: {{ CurProject.time_started }}</div>
            <div class="description">Статус: 
              <span v-if="CurProject.status" class="status-active">Активен <span style="font-weight: normal; color: black;">(xxx.xxx.xxx.xxx:xxxx)</span></span>
              <span v-else class="status-disable">Завершен</span>
            </div>
          </div>
          <div class="control">
            <n-button size="small" ghost type="primary">Генерация агентов</n-button>
            <n-button size="small" ghost type="info">Экспорт таблицы</n-button>
            <n-button size="small" ghost type="error" @click="stopProject(CurProject.id)" v-if="CurProject.status">Завершить проект</n-button> 
            <n-button size="small" ghost type="error" v-else disabled>Завершить проект</n-button> 
          </div>
        </div>
      </n-card>
      <n-card class="data-table">
        <n-button>Button</n-button>
      </n-card>
    </div>
</template>

<style scoped>
.project {
  flex-grow: 1;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  align-items: center;  
  gap: 16px;
}

.panel {
  display: flex;
  justify-content: space-between;
}

.control {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  gap: 4px;
  min-width: 25%;
}

.info {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex-grow: 1;
}

.project-name {
  font-size: 20px;
  font-weight: bold;
}

.status-active {
  font-weight: bold;
  color: #00a854;
}

.status-disable {
  color: #f04134;
}

.description {
  font-size: 14px;
  color: #666;
}

.contol-menu {
  margin-top: 16px;
  max-width: 80%;
  /* min-height: 128px; */
}

.data-table {
  flex-grow: 1;
  margin-bottom: 16px;
  width: 80%;
  overflow-y: auto;
}
</style>