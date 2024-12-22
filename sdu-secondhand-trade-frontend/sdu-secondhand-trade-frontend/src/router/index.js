import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login/index.vue'
import Layout from '@/views/Layout/index.vue'
import Home from '@/views/Home/index.vue'
import Category from '@/views/Category/index.vue'
import SubCategory from '@/views/SubCategory/index.vue'
import Detail from '@/views/Detail/index.vue'
import Order from '@/views/Order/index.vue'
import Pay from '@/views/Pay/index.vue'
import PayBack from '@/views/Pay/PayBack.vue'
import User from '@/views/User/index.vue'
import UserInfo from '@/views/User/components/UserInfo.vue'
import UserBuyOrder from '@/views/User/components/UserBuyOrder.vue'
import UserSellOrder from '@/views/User/components/UserSellOrder.vue'
import Sell from '@/views/Sell/index.vue'
import UserGood from '@/views/User/components/UserGood.vue'
import UserAfterSale from '@/views/User/components/UserAfterSale.vue'
import AdminUser from '@/views/User/components/AdminUser.vue'
import AdminGood from '@/views/User/components/AdminGood.vue'
import AdminAfterSale from '@/views/User/components/AdminAfterSale.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Layout,
      children: [
        {
          path: '',
          name: 'home',
          component: Home
        },
        {
          path: 'category/:id',
          name: 'category',
          component: Category
        },
        {
          path: 'category/sub/:id',
          name: 'sub-category',
          component: SubCategory
        },
        {
          path: 'detail/:id',
          name: 'detail',
          component: Detail
        },
        {
          path: 'order/:id',
          name: 'order',
          component: Order
        },
        {
          path: 'pay/:id',
          name: 'pay',
          component: Pay
        },
        {
          path: 'pay/callback/:id',
          name: 'payBack',
          component: PayBack,
        },
        {
          path: 'user',
          name: 'user',
          component: User,
          children: [
            {
              path: '',
              name: 'info',
              component: UserInfo,
            },
            {
              path: 'buyer_order',
              name: 'buyer_order',
              component: UserBuyOrder,
            },
            {
              path: 'seller_order',
              name: 'seller_order',
              component: UserSellOrder,
            },
            {
              path: 'good',
              name: 'good',
              component: UserGood,
            },
            {
              path: 'problem',
              name: 'problem',
              component: UserAfterSale,
            },
            {
              path: 'admin/user',
              name: 'adminUser',
              component: AdminUser,
            },
            {
              path: 'admin/good',
              name: 'adminGood',
              component: AdminGood,
            },
            {
              path: 'admin/problem',
              name: 'adminProblem',
              component: AdminAfterSale,
            }
          ]
        },
        {
          path: 'sell',
          name: 'sell',
          component: Sell,
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ],
  scrollBehavior() {
    return {
      top: 0
    }
  }
})

export default router
