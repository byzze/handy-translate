<template>
  <div class="close-button" @click="close">
    <svg viewBox="0 0 40 40">
      <path class="close-x" d="M 10,10 L 30,30 M 30,10 L 10,30" />
    </svg>
  </div>
</template>

<script setup>
import { WindowHide } from '../../wailsjs/runtime/runtime'

function close() {
  // 隐藏逻辑
  WindowHide()
}

</script>

<style>
.close-button {
  position: absolute;
  top: 5px;
  right: 5px;
  width: 18px;
  height: 18px;
  border-radius: 50%;
  cursor: pointer;
}

.close-x {
  stroke: white;
  stroke-width: 4px;
  stroke-linecap: round;
}
</style>