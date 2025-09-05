<template>
  <div class="create-article-container">
    <h1 class="section-title">发布文章</h1>
    
    <div v-if="success" class="alert alert-success">
      <p>文章发布成功！</p>
      <div class="alert-actions">
        <router-link :to="{ name: 'ArticleDetail', params: { id: newArticleId } }" class="btn">
          查看文章
        </router-link>
        <button @click="resetForm" class="btn btn-secondary">
          发布新文章
        </button>
      </div>
    </div>
    
    <div v-else>
      <div v-if="error" class="alert alert-error">
        {{ error }}
      </div>
      
      <ArticleForm @submit="createArticle" />
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { articleAPI } from '../api/api';
import ArticleForm from '../components/ArticleForm.vue';

export default {
  name: 'CreateArticleView',
  components: {
    ArticleForm
  },
  setup() {
    const error = ref('');
    const success = ref(false);
    const newArticleId = ref(null);
    
    const createArticle = async (articleData) => {
      error.value = '';
      
      try {
        const response = await articleAPI.createArticle(articleData);
        success.value = true;
        newArticleId.value = response.data.ID;
      } catch (err) {
        console.error('发布文章失败:', err);
        if (err.response && err.response.data) {
          error.value = err.response.data.error || '发布文章失败，请重试';
        } else {
          error.value = '发布文章失败，请检查网络连接';
        }
      }
    };
    
    const resetForm = () => {
      success.value = false;
      newArticleId.value = null;
      error.value = '';
    };
    
    return {
      error,
      success,
      newArticleId,
      createArticle,
      resetForm
    };
  }
};
</script>

<style scoped>
.create-article-container {
  padding: 20px 0;
}

.alert {
  padding: 15px;
  border-radius: 4px;
  margin-bottom: 20px;
}

.alert-success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.alert-error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.alert-actions {
  margin-top: 15px;
  display: flex;
  gap: 10px;
}
</style>
