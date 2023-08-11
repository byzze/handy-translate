<template>
  <n-config-provider :theme-overrides="darkTheme">
    <n-loading-bar-provider>
      <n-message-provider>
        <n-notification-provider>
          <n-dialog-provider>
            <Home />
          </n-dialog-provider>
        </n-notification-provider>
      </n-message-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>
  
<script>
import { defineComponent } from "vue";
import Home from "./components/Home.vue";

import { darkTheme } from 'naive-ui'


export default defineComponent({
  setup() {
    return {
      darkTheme
    }
  },
  components: {
    Home
  },
});
</script>