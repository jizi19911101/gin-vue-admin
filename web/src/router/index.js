import { createRouter, createWebHashHistory } from 'vue-router'
import Devices from '@/view/monkeyTest/devices.vue'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/devices',
  name: 'Devices',
  iframeComponent: Devices
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
