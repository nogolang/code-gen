import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'layout',
			component: () => import('@/views/layout/layoutContainer.vue'),
			redirect: '/fileGen/database',
			children: [
				{
					path: '/file', children: [
						{
							path: '/fileGen/database',
							name: 'database',
							component: () => import('@/views/fileGen/databaseManager.vue')
						},
						{
							path: '/fileGen/mapping',
							name: 'mapping',
							component: () => import('@/views/fileGen/mappingManger.vue')
						},
						{
							path: '/fileGen/fileManager',
							name: 'productInfo',
							component: () => import('@/views/fileGen/fileManager.vue')
						},
						{
							path: '/fileGen/fileGroup',
							name: 'fileGroup',
							component: () => import('@/views/fileGen/groupManager.vue')
						},
						{
							path: '/fileGen/fileGen',
							name: 'fileGen',
							component: () => import('@/views/fileGen/fileGenManger.vue')
						},

					]
				},
			]
		},
		//{path: '/login', name: 'login', component: () => import('@/views/login/login.vue')},
	]
})

//路由导航守卫
//router.beforeEach((to, from,next) => {
//  let passUrl=["/login"]
//
//  //如果不是放行的目录，则需要判定token
//  //token为空，则跳转到登录页面
//  if (!passUrl.includes(to.path)){
//    if (userStore().getToken()===""){
//      ElMessage.warning("请先登录")
//      return  next("/login")
//    }
//  }
//  return next()
//})


export default router
