import { createApp } from 'vue'
import App from './App.vue'

import { createRouter, createWebHistory } from 'vue-router'
import Workspace from './components/Workspace.vue'
import ProjectPage from './components/ProjectPage.vue'

// Create Vue router
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Workspace },
    { path: '/project/:id', component: ProjectPage},
    { path: '/:pathMatch(.*)*', redirect: '/'}
  ],
})

// Create Vue app
const app = createApp(App)
app.use(router)
app.mount('#app')
