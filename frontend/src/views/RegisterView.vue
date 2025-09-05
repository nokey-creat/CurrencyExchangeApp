<template>
  <div class="register-container">
    <div class="register-form card">
      <h1 class="section-title">注册</h1>
      
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
        
        <div class="form-group">
          <label for="confirmPassword">确认密码</label>
          <input 
            type="password" 
            id="confirmPassword" 
            v-model="confirmPassword" 
            class="form-control" 
            required
          />
          <div v-if="!passwordsMatch" class="form-error">
            两次输入的密码不一致
          </div>
        </div>
        
        <div class="form-actions">
          <button type="submit" class="btn" :disabled="isLoading || !passwordsMatch">
            {{ isLoading ? '注册中...' : '注册' }}
          </button>
        </div>
      </form>
      
      <div class="form-footer">
        <p>已有账号? <router-link to="/login">立即登录</router-link></p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
  name: 'RegisterView',
  setup() {
    const store = useStore();
    const router = useRouter();
    
    const username = ref('');
    const password = ref('');
    const confirmPassword = ref('');
    const error = ref('');
    const isLoading = ref(false);
    
    const passwordsMatch = computed(() => {
      return password.value === confirmPassword.value;
    });
    
    const handleSubmit = async () => {
      // 检查密码是否匹配
      if (!passwordsMatch.value) {
        return;
      }
      
      error.value = '';
      isLoading.value = true;
      
      try {
        await store.dispatch('auth/register', {
          username: username.value,
          password: password.value
        });
        
        // 注册成功，跳转到首页
        router.push('/');
      } catch (err) {
        // 注册失败，显示错误消息
        if (err.response && err.response.data) {
          error.value = err.response.data.error || '注册失败，请重试';
        } else {
          error.value = '注册失败，请检查网络连接';
        }
      } finally {
        isLoading.value = false;
      }
    };
    
    return {
      username,
      password,
      confirmPassword,
      passwordsMatch,
      error,
      isLoading,
      handleSubmit
    };
  }
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.register-form {
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

.form-error {
  color: #dc3545;
  font-size: 14px;
  margin-top: 5px;
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
