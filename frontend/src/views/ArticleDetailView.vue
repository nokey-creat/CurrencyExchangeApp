<template>
  <div class="article-detail-container">
    <div v-if="isLoading" class="loading">
      <p>加载中...</p>
    </div>
    
    <div v-else-if="error" class="error-message">
      <p>{{ error }}</p>
      <button @click="fetchArticle" class="btn">重试</button>
      <router-link to="/" class="btn btn-secondary">返回首页</router-link>
    </div>
    
    <div v-else class="article-detail">
      <h1 class="article-title">{{ article.title }}</h1>
      
      <div class="article-meta">
        <span class="article-author">作者: {{ article.author }}</span>
        <span class="article-date">发布于: {{ formatDate(article.CreatedAt) }}</span>
      </div>
      
      <div class="article-content">
        {{ article.content }}
      </div>
      
      <div class="article-actions">
        <div class="like-section">
          <button @click="likeArticle" class="like-button" :class="{ liked: hasLiked }">
            <span class="heart-icon">❤️</span>
            <span>{{ likes }} 赞</span>
          </button>
        </div>
        
        <router-link to="/" class="btn">返回首页</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { articleAPI } from '../api/api';

export default {
  name: 'ArticleDetailView',
  setup() {
    const route = useRoute();
    const articleId = route.params.id;
    
    const article = ref({});
    const likes = ref(0);
    const hasLiked = ref(false);
    const isLoading = ref(true);
    const error = ref('');
    
    // 格式化日期
    const formatDate = (dateString) => {
      const date = new Date(dateString);
      return date.toLocaleDateString('zh-CN', { 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });
    };
    
    // 获取文章内容
    const fetchArticle = async () => {
      isLoading.value = true;
      error.value = '';
      
      try {
        const response = await articleAPI.getArticleById(articleId);
        article.value = response.data;
        await fetchLikes();
      } catch (err) {
        console.error('获取文章详情失败:', err);
        error.value = '获取文章详情失败，请重试';
      } finally {
        isLoading.value = false;
      }
    };
    
    // 获取文章点赞数
    const fetchLikes = async () => {
      try {
        const response = await articleAPI.getArticleLikes(articleId);
        likes.value = response.data.likes;
      } catch (err) {
        console.error('获取点赞数失败:', err);
      }
    };
    
    // 点赞文章
    const likeArticle = async () => {
      if (hasLiked.value) return;
      
      try {
        const response = await articleAPI.likeArticle(articleId);
        likes.value = response.data.likes;
        hasLiked.value = true;
        
        // 存储用户已点赞状态
        const likedArticles = JSON.parse(localStorage.getItem('likedArticles') || '{}');
        likedArticles[articleId] = true;
        localStorage.setItem('likedArticles', JSON.stringify(likedArticles));
      } catch (err) {
        console.error('点赞失败:', err);
      }
    };
    
    onMounted(() => {
      fetchArticle();
      
      // 检查用户是否已点赞
      const likedArticles = JSON.parse(localStorage.getItem('likedArticles') || '{}');
      hasLiked.value = likedArticles[articleId] || false;
    });
    
    return {
      article,
      likes,
      hasLiked,
      isLoading,
      error,
      formatDate,
      fetchArticle,
      likeArticle
    };
  }
};
</script>

<style scoped>
.article-detail-container {
  padding: 20px 0;
}

.loading, .error-message {
  text-align: center;
  padding: 40px 0;
}

.error-message {
  color: #dc3545;
}

.error-message .btn {
  margin: 0 5px;
}

.article-detail {
  background-color: #fff;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  padding: 30px;
}

.article-title {
  margin-bottom: 15px;
  font-size: 28px;
}

.article-meta {
  display: flex;
  gap: 15px;
  font-size: 14px;
  color: #666;
  margin-bottom: 25px;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.article-content {
  margin-bottom: 30px;
  line-height: 1.8;
  white-space: pre-line;
}

.article-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.like-button {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 8px 16px;
  background-color: transparent;
  border: 1px solid #ddd;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s;
}

.like-button:hover {
  background-color: #f0f0f0;
}

.like-button.liked {
  background-color: #ffecec;
  border-color: #ffb3b3;
  cursor: default;
}

.heart-icon {
  font-size: 18px;
}
</style>
