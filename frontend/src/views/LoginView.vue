<template>
  <div class="login-container">
    <div class="login-form card">
      <h1 class="section-title">登录</h1>
      
      <div v-if="error" class="alert alert-error">
        {{ error }}
      </div>
      
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="username">用户名</label>
          <input 
            type="text" 
            id="username" 
            v-model="username" 
            class="form-control" 
            required
          />
        </div>
        
        <div class="form-group">
          <label for="password">密码</label>
          <input 
            type="password" 
            id="password" 
            v-model="password" 
            class="form-control" 
            required
          />
        </div>
        
        <div class="form-actions">
          <button type="submit" class="btn" :disabled="isLoading">
            {{ isLoading ? '登录中...' : '登录' }}
          </button>
        </div>
      </form>
      
      <div class="form-footer">
        <p>还没有账号? <router-link to="/register">立即注册</router-link></p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
  name: 'LoginView',
  setup() {
    const store = useStore();
    const router = useRouter();
    
    const username = ref('');
    const password = ref('');
    const error = ref('');
    const isLoading = ref(false);
    
    const handleSubmit = async () => {
      error.value = '';
      isLoading.value = true;
      
      try {
        await store.dispatch('auth/login', {
          username: username.value,
          password: password.value
        });
        
        // 登录成功，跳转到首页
        router.push('/');
      } catch (err) {
        // 登录失败，显示错误消息
        if (err.response && err.response.data) {
          error.value = err.response.data.error || '登录失败，请重试';
        } else {
          error.value = '登录失败，请检查网络连接';
        }
      } finally {
        isLoading.value = false;
      }
    };
    
    return {
      username,
      password,
      error,
      isLoading,
      handleSubmit
    };
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.login-form {
  width: 100%;
  max-width: 400px;
}

.alert {
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
}

.alert-error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.form-footer {
  margin-top: 20px;
  text-align: center;
  font-size: 14px;
}

.form-footer a {
  color: #007bff;
  text-decoration: none;
}

.form-footer a:hover {
  text-decoration: underline;
}
</style>
