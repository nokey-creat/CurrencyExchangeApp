<template>
  <nav class="navbar">
    <div class="navbar-logo">
      <router-link to="/">文章网站</router-link>
    </div>
    <div class="navbar-menu">
      <router-link to="/" class="navbar-item">首页</router-link>
      <router-link to="/create-article" class="navbar-item">发布文章</router-link>
    </div>
    <div class="navbar-user">
      <span class="username">{{ username }}</span>
      <button @click="logout" class="btn btn-secondary btn-sm">退出</button>
    </div>
  </nav>
</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
  name: 'Navbar',
  setup() {
    const store = useStore();
    const router = useRouter();
    
    const username = computed(() => store.getters['auth/username']);
    
    const logout = () => {
      store.dispatch('auth/logout');
      router.push('/login');
    };
    
    return {
      username,
      logout
    };
  }
};
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 0;
  margin-bottom: 30px;
  border-bottom: 1px solid #eee;
}

.navbar-logo a {
  font-size: 22px;
  font-weight: bold;
  color: #333;
  text-decoration: none;
}

.navbar-menu {
  display: flex;
  gap: 20px;
}

.navbar-item {
  color: #555;
  text-decoration: none;
  transition: color 0.3s;
}

.navbar-item:hover {
  color: #007bff;
}

.navbar-user {
  display: flex;
  align-items: center;
  gap: 10px;
}

.username {
  font-weight: 500;
}

.btn-sm {
  padding: 5px 10px;
  font-size: 12px;
}
</style>
