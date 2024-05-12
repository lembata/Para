import { createRouter, createWebHistory } from 'vue-router'
//import HomeView from '../views/HomeView.vue'
import DashboardView from '../views/DashboardView.vue'
import AppLayout from '@/layout/AppLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          name: 'dashboard',
          component: () => import('@/views/DashboardView.vue')
        },
        {
          path: '/accounts',
          name: 'accounts',
          component: () => import('@/views/AccountsView.vue')
        },
        {
          path: '/accounts/add',
          name: 'accounts-add',
          component: () => import('@/views/AccountsView.vue')
        },
        {
          path: '/accounts/:id',
          name: 'accounts-details',
          component: () => import('@/views/AccountsView.vue')
        },
        {
          path: '/accounts/add',
          name: 'accounts-add',
          component: () => import('@/views/AccountsEditView.vue')
        },
        {
          path: '/accounts/edit/:id',
          name: 'accounts-edit',
          component: () => import('@/views/AccountsEditView.vue')
        },
      ]
    }
    // {
    //   path: '/accounts',
    //   name: 'accounts',
    //   component: () => import('../views/AccountsView.vue')
    // },
    // {
    //   path: '/accounts/add',
    //   name: 'accounts-add',
    //   component: () => import('../views/AccountsView.vue')
    // },
    // {
    //   path: '/settings',
    //   name: 'settings',
    //   component: () => import('../views/SettingsView.vue')
    // }
  ]
})

export default router
