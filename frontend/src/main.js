import { createApp } from 'vue';
import { createStore } from 'vuex';
import App from './App.vue';
import router from './router';
import authModule from './store/auth';

// 创建 Vuex store
const store = createStore({
  modules: {
    auth: authModule
  }
});

// 创建应用实例
const app = createApp(App);

// 注册插件
app.use(router);
app.use(store);

// 挂载应用
app.mount('#app');
