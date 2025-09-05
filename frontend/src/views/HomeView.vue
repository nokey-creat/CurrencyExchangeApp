<template>
  <div class="home-container">
    <h1 class="section-title">最新文章</h1>
    
    <div v-if="isLoading" class="loading">
      <p>加载中...</p>
    </div>
    
    <div v-else-if="error" class="error-message">
      <p>{{ error }}</p>
      <button @click="fetchArticles" class="btn">重试</button>
    </div>
    
    <div v-else-if="articles.length === 0" class="empty-state">
      <p>还没有文章，快来发布第一篇吧！</p>
      <router-link to="/create-article" class="btn">写文章</router-link>
    </div>
    
    <div v-else class="article-list">
      <ArticlePreview 
        v-for="article in sortedArticles" 
        :key="article.ID" 
        :article="article" 
      />
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue';
import { articleAPI } from '../api/api';
import ArticlePreview from '../components/ArticlePreview.vue';

export default {
  name: 'HomeView',
  components: {
    ArticlePreview
  },
  setup() {
    const articles = ref([]);
    const isLoading = ref(true);
    const error = ref('');
    
    // 按ID降序排列文章（新文章在上面）
    const sortedArticles = computed(() => {
      return [...articles.value].sort((a, b) => b.ID - a.ID);
    });
    
    const fetchArticles = async () => {
      isLoading.value = true;
      error.value = '';
      
      try {
        const response = await articleAPI.getArticles();
        articles.value = response.data;
      } catch (err) {
        console.error('获取文章列表失败:', err);
        error.value = '获取文章列表失败，请重试';
      } finally {
        isLoading.value = false;
      }
    };
    
    onMounted(() => {
      fetchArticles();
    });
    
    return {
      articles,
      sortedArticles,
      isLoading,
      error,
      fetchArticles
    };
  }
};
</script>

<style scoped>
.home-container {
  padding: 20px 0;
}

.loading, .error-message, .empty-state {
  text-align: center;
  padding: 40px 0;
}

.error-message {
  color: #dc3545;
}

.empty-state {
  color: #6c757d;
}

.empty-state p {
  margin-bottom: 20px;
}
</style>
