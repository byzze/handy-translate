<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { LogPrint } from '../../wailsjs/runtime/runtime'
import { EventsOn } from '../../wailsjs/runtime/runtime'

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
.content {
  position: absolute;
  top: 0;
  right: 0;
  font-size: 1.0rem;
  cursor: pointer;
  width: 8%;
  height: 8%;
  background-color: red;
}


.scroll-container {
  margin-top: 50px;
  margin-left: 15px;
  margin-right: 15px;
}

.result {
  line-height: 28px;
  font-size: 18px;
  /* margin: 1.5rem auto; */
}

.query {
  line-height: 20px;
  margin: 1.5rem auto;
}

.explian {
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
