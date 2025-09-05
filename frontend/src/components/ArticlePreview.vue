<template>
  <div class="article-preview card">
    <h2 class="article-title">
      <router-link :to="{ name: 'ArticleDetail', params: { id: article.ID } }">
        {{ article.title }}
      </router-link>
    </h2>
    <div class="article-meta">
      <span class="article-author">作者: {{ article.author }}</span>
      <span class="article-date">发布于: {{ formatDate(article.CreatedAt) }}</span>
    </div>
    <p class="article-preview-text">{{ article.preview }}</p>
    <div class="article-actions">
      <router-link :to="{ name: 'ArticleDetail', params: { id: article.ID } }" class="btn">
        阅读全文
      </router-link>
    </div>
  </div>
</template>

<script>
import { onMounted } from 'vue';

export default {
  name: 'ArticlePreview',
  props: {
    article: {
      type: Object,
      required: true
    }
  },
  setup(props) {
    // 格式化日期
    const formatDate = (dateString) => {
      const date = new Date(dateString);
      return date.toLocaleDateString('zh-CN', { 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
      });
    };

    return {
      formatDate
    };
  }
};
</script>

<style scoped>
.article-preview {
  margin-bottom: 25px;
}

.article-title {
  margin-bottom: 10px;
}

.article-title a {
  color: #333;
  text-decoration: none;
  transition: color 0.3s;
}

.article-title a:hover {
  color: #007bff;
}

.article-meta {
  display: flex;
  gap: 15px;
  font-size: 14px;
  color: #666;
  margin-bottom: 15px;
}

.article-preview-text {
  margin-bottom: 15px;
  color: #555;
}

.article-actions {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
