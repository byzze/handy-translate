<template>
    <n-space vertical size="large">
        <n-layout position="absolute" style=" background-color: rgba(27, 38, 54, 1);">
            <n-layout-header>
                <n-grid x-gap="12" :cols="3">
                    <n-gi>
                        <div class="setting">
                            <Setting />
                        </div>
                    </n-gi>
                    <n-gi>
                        <div class="center">
                            <n-icon size="35">
                                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                                    viewBox="0 0 24 24">
                                    <path
                                        d="M20 5h-9.12L10 2H4c-1.1 0-2 .9-2 2v13c0 1.1.9 2 2 2h7l1 3h8c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zM7.17 14.59c-2.25 0-4.09-1.83-4.09-4.09s1.83-4.09 4.09-4.09c1.04 0 1.99.37 2.74 1.07l.07.06l-1.23 1.18l-.06-.05c-.29-.27-.78-.59-1.52-.59c-1.31 0-2.38 1.09-2.38 2.42s1.07 2.42 2.38 2.42c1.37 0 1.96-.87 2.12-1.46H7.08V9.91h3.95l.01.07c.04.21.05.4.05.61c0 2.35-1.61 4-3.92 4zm6.03-1.71c.33.6.74 1.18 1.19 1.7l-.54.53l-.65-2.23zm.77-.76h-.99l-.31-1.04h3.99s-.34 1.31-1.56 2.74c-.52-.62-.89-1.23-1.13-1.7zM21 20c0 .55-.45 1-1 1h-7l2-2l-.81-2.77l.92-.92L17.79 18l.73-.73l-2.71-2.68c.9-1.03 1.6-2.25 1.92-3.51H19v-1.04h-3.64V9h-1.04v1.04h-1.96L11.18 6H20c.55 0 1 .45 1 1v13z"
                                        fill="currentColor"></path>
                                </svg>
                            </n-icon>
                        </div>
                    </n-gi>
                    <n-gi>
                        <div class="right">
                            <Minimize />
                        </div>
                    </n-gi>
                </n-grid>
            </n-layout-header>
            <n-layout-content position="absolute" style="bottom: 0;top:10vh">
                <Translate />
            </n-layout-content>

        </n-layout>
        <!-- <n-layout>
                <n-layout-footer>成府路</n-layout-footer>
            </n-layout>
        </n-layout> -->
    </n-space>
</template>
  
<script>
import { defineComponent } from "vue";
import Setting from "./Setting.vue";
import Translate from "./Translate.vue";
import Minimize from "./Minimize.vue";

export default defineComponent({
    components: {
        Setting,
        Translate,
        Minimize
    },
});
</script>
<style scoped>
.setting {
    display: flex;
    align-items: flex-start;
    color: white;
}

.light-green {
    /* height: 108px; */
    /* background-color: rgba(0, 128, 0, 0.12); */
}

.center {
    top: 12px
        /* height: 108px; */
        /* background-color: rgba(0, 128, 0, 0.24); */
}

.right {
    /* top: 10px; */
    /* height: 108px; */
    background-color: rgba(0, 128, 0, 0.24);
}

.n-layout-header,
.n-layout-footer {
    background-color: rgba(27, 38, 54, 1);

    /* padding: 24px; */
}

.n-layout-sider {
    background: rgba(128, 128, 128, 0.3);
}

.n-layout-content {
    font-family: "Nunito", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
        "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
        sans-serif;
    color: white;
    background-color: rgba(27, 38, 54, 1);
}
</style>