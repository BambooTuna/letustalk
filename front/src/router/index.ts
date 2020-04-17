import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/content',
    name: 'Content',
    component: () => import('@/views/Content.vue')
  },
  {
    path: '/host',
    name: 'Host',
    component: () => import('@/views/Host.vue')
  },
  {
    path: '/payment/test/:invoiceId',
    name: 'TestPayment',
    component: () => import('@/views/TestPayment.vue')
  },
  {
    path: '/account/:accountId/calendar',
    name: 'ScheduleCalendar',
    component: () => import('@/views/ScheduleCalendar.vue')
  },
  {
    path: '*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
