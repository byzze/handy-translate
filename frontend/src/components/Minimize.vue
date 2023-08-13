<template>
    <div class="minimize" @click="minimizeApp">
        <n-button>
            <n-icon>
                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32">
                    <path d="M4 18v2h6.586L2 28.582L3.414 30L12 21.414V28h2V18H4z" fill="currentColor">
                    </path>
                    <path d="M30 3.416L28.592 2L20 10.586V4h-2v10h10v-2h-6.586L30 3.416z" fill="currentColor"></path>
                </svg>
            </n-icon>
        </n-button>
    </div>
</template>
  
 
<script lang="ts">
import { defineComponent } from 'vue'
import { WindowHide } from '../../wailsjs/runtime/runtime'
export default defineComponent({
    setup() {
        return {
            minimizeApp() {
                WindowHide()
            }
        }
    }
})
</script>
  
<style>
.minimize {
    float: right;
    /* padding: 10px 5px; */
    color: white;
    cursor: pointer;
}
</style>
  