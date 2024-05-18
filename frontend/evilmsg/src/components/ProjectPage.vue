<script lang="ts">
  import { NCard, NButton, NInputGroup, NSelect, NDataTable, NCode } from 'naive-ui'
  import { useRoute, onBeforeRouteLeave } from 'vue-router'
  import { ref, onMounted, h } from 'vue'
  import type { DataTableColumns } from 'naive-ui'

  type Project = {
    id: number
    name: string
    status: boolean
    time_stopped: string
    time_started: string
  }

  var CurProject = ref({} as Project)

  const fetchProject = async (id: string) => {
    const response = await fetch(`/api/projects/${id}`)
    const data = await response.json()
    CurProject.value = data
    console.log(data)
  }

  const stopProject = async (id: number) => {
    await fetch(`/api/projects/${id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ status: 0 })
    })

    fetchProject(id.toString())

    console.log(`Stopped project with id ${id}`)
  }

  type Hit = {
    id: number
    time: string
    ip: string
    data: string
  }

  var Hits = ref([] as Hit[])

  const fetchHits = async (id: string) => {
    const response = await fetch(`/api/hit/${id}`)
    const data = await response.json()
    Hits.value = data
    console.log(data)
  }

  var intervalId: number

  const startFetchingHits = (id: string) => {
    intervalId = setInterval(() => {
      fetchHits(id)
    }, 1000)
  }

  const stopFetchingHits = () => {
    clearInterval(intervalId)
    Hits.value = []
  }

  const HitsColumns = (): DataTableColumns<Hit> => [
    {
      title: 'ID',
      key: 'id',
      width: 50
    },
    {
      title: 'Время',
      key: 'time',
      width: 200
    },
    {
      title: 'IP',
      key: 'ip',
      width: 150
    },
    {
      title: 'Данные',
      key: 'data',
      render (row, _) {
        var rdata = JSON.stringify(JSON.parse(row.data), null, 2);
        return h(NCode, {
          language: 'json',
          wordWrap: true,
          code: rdata
        })
      }
    }
  ]

  const generateAgents = (id: number) => {
    location.href = `/api/agent/linux/raw/${id}`

    console.log(`Generated agents for project with id ${id}`)
  }

  export default {
    components: {
      NCard,
      NButton,
      NInputGroup,
      NSelect,
      NDataTable
    },
    setup() {
      const route = useRoute()
      const id = route.params.id as string

      onMounted(() => {
        fetchProject(id)
        fetchHits(id)
        startFetchingHits(id)
      })

      onBeforeRouteLeave(() => {
        stopFetchingHits()
      })

      return {
        CurProject,
        stopProject,
        HitsColumns,
        generateAgents,
        Hits,
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
              <span v-if="CurProject.status" class="status-active">Активен</span>
              <span v-else class="status-disable">Завершен ({{ CurProject.time_stopped }})</span>
            </div>
          </div>
          <div class="control">
            <n-button size="small" ghost type="primary" v-if="CurProject.status" @click="generateAgents(CurProject.id)">Генерация агентов</n-button>
            <n-button size="small" ghost type="primary" v-else disabled>Генерация агентов</n-button>
            <n-button size="small" ghost type="info">Экспорт таблицы</n-button>
            <n-button size="small" ghost type="error" @click="stopProject(CurProject.id)" v-if="CurProject.status">Завершить проект</n-button> 
            <n-button size="small" ghost type="error" v-else disabled>Завершить проект</n-button> 
          </div>
        </div>
      </n-card>
      <n-card class="data-table">
        <n-data-table
          :columns=HitsColumns()
          :data="Hits"
          style="height: 100%;"
        />
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