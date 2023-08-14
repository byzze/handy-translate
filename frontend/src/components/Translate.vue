<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { LogPrint, EventsOn } from '../../wailsjs/runtime/runtime'

const data = reactive({
  name: "",
  queryText: "",
  resultText: "",
  explianText: "",
  showCloseBtn: true,
  showModal: true,
})

EventsOn("query", (result) => {
  data.queryText = result
})

EventsOn("result", (result) => {
  data.resultText = result
})

EventsOn("explian", (result) => {
  data.explianText = result
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <!-- <main> -->
  <div class="scroll-container">
    <div id="query" class="query">{{ data.queryText }}</div>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="explian" class="explian">{{ data.explianText }}</div>
  </div>
  <!-- </main> -->
</template>

<style scoped>
.scroll-container {
  margin-left: 1rem;
  margin-right: 1rem;
}

.result {
  line-height: 1.5rem;
  font-size: 1.5rem;
  /* margin: 1.5rem auto; */
}

.query {
  line-height: 1.5rem;
  font-size: 1.5rem;
  margin: 1.5rem auto;
}

.explian {
  line-height: 1.5rem;
  margin: 1.5rem auto;
}
</style>
