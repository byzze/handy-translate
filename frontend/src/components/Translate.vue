<template>
  <!-- 根据是否有数据来显示加载动画或实际内容 -->
  <n-space vertical>
    <n-spin size="small" :show="isLoading">
      <br />
    </n-spin>
    <div class="scroll-container">
      <div id="query" class="query">{{ queryText }}</div>
      <div id="result" class="result">{{ resultText }}</div>
      <div id="explian" class="explian">{{ explianText }}</div>
    </div>
  </n-space>
</template>


<script>
import { defineComponent, ref, reactive, onMounted } from 'vue'
import { LogPrint, Hide, EventsOn, EventsEmit } from '../../wailsjs/runtime/runtime'









export default defineComponent({
  setup() {
    const queryTextRef = ref("");
    const resultTextRef = ref("");
    const explianTextRef = ref("");
    const isLoadingRef = ref(false);
    // 模拟数据获取
    onMounted(() => {
      EventsOn("explian", (result) => {
        explianTextRef.value = result
      })

      EventsOn("result", (result) => {
        resultTextRef.value = result
      })

      EventsOn("query", (result) => {
        queryTextRef.value = result
      })

      EventsOn("loading", (result) => {
        queryTextRef.value = ""
        resultTextRef.value = ""
        explianTextRef.value = ""
        isLoadingRef.value = result
      })

    });
    return {
      queryText: queryTextRef,
      resultText: resultTextRef,
      explianText: explianTextRef,
      isLoading: isLoadingRef, // 初始时显示加载动画
    }
  }
})
</script>

<style scoped>
.scroll-container {
  margin-left: 1rem;
  margin-right: 1rem;
}

.result {
  line-height: 1.5rem;
  font-size: 1.3rem;
  /* margin: 1.5rem auto; */
}

.query {
  line-height: 1.5rem;
  font-size: 1.3rem;
  margin: 1.5rem auto;
}

.explian {
  line-height: 1.5rem;
  margin: 1.5rem auto;
}
</style>
