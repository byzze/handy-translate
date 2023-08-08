<template>
  <div class="modal-container">
    <div class="modal">
      <div class="modal-content">
        <span class="close-btn" @click="closeModal">&times;</span>
        <p>This is a modal content.</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    closeModal() {
      console.log('Modal closed.');
    }
  }
}
</script>

<style scoped>
.modal-container {
  position: fixed;
  top: 20px;
  /* 调整距离顶部的距离 */
  right: 20px;
  /* 调整距离右侧的距离 */
  z-index: 1000;
}

.modal {
  display: block;
  position: relative;
  width: 300px;
  background-color: white;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.modal-content {
  padding: 20px;
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 15px;
  font-size: 1.5rem;
  cursor: pointer;
}
</style>
