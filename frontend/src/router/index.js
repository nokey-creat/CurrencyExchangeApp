import { createRouter, createWebHistory } from 'vue-router';

// 导入视图组件
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';
import ArticleDetailView from '../views/ArticleDetailView.vue';
import CreateArticleView from '../views/CreateArticleView.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterView
  },
  {
    path: '/articles/:id',
    name: 'ArticleDetail',
    component: ArticleDetailView,
    meta: { requiresAuth: true },
    props: true
  },
  {
    path: '/create-article',
    name: 'CreateArticle',
    component: CreateArticleView,
    meta: { requiresAuth: true }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 导航守卫
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token');
  
  if (to.matched.some(record => record.meta.requiresAuth) && !isAuthenticated) {
    // 需要认证但用户未登录，重定向到登录页
    next({ name: 'Login' });
  } else {
    // 继续导航
    next();
  }
});

export default router;
