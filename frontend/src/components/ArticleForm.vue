<template>
  <form @submit.prevent="submitForm" class="article-form">
    <div class="form-group">
      <label for="title">标题</label>
      <input 
        type="text" 
        id="title" 
        v-model="formData.title" 
        class="form-control" 
        required
      />
    </div>
    
    <div class="form-group">
      <label for="preview">文章预览</label>
      <textarea 
        id="preview" 
        v-model="formData.preview" 
        class="form-control" 
        rows="3"
        required
      ></textarea>
    </div>
    
    <div class="form-group">
      <label for="content">文章内容</label>
      <textarea 
        id="content" 
        v-model="formData.content" 
        class="form-control" 
        rows="10"
        required
      ></textarea>
    </div>
    
    <div class="form-actions">
      <button type="submit" class="btn" :disabled="isSubmitting">
        {{ isSubmitting ? '提交中...' : '发布文章' }}
      </button>
    </div>
  </form>
</template>

<script>
import { reactive, ref } from 'vue';

export default {
  name: 'ArticleForm',
  emits: ['submit'],
  setup(props, { emit }) {
    const formData = reactive({
      title: '',
      preview: '',
      content: ''
    });
    
    const isSubmitting = ref(false);
    
    const submitForm = () => {
      isSubmitting.value = true;
      
      // 提交表单数据
      emit('submit', { ...formData });
      
      // 重置表单
      formData.title = '';
      formData.preview = '';
      formData.content = '';
      
      isSubmitting.value = false;
    };
    
    return {
      formData,
      isSubmitting,
      submitForm
    };
  }
};
</script>

<style scoped>
.article-form {
  background-color: #fff;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  padding: 25px;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
}

.form-actions {
  margin-top: 20px;
}

textarea {
  resize: vertical;
}
</style>
