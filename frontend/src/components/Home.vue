<template>
    <n-space vertical size="large">
        <n-layout>
            <n-layout-header>
                <n-grid x-gap="12" :cols="3">
                    <n-gi>
                        <div class="setting">
                            <Setting />
                        </div>
                    </n-gi>
                    <n-gi>
                        <div class="center">
                            <n-icon size="40">
                                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                                    viewBox="0 0 32 32">
                                    <path
                                        d="M17 14.5l6 2.5v-2l-6-3V9a1 1 0 0 0-2 0v3l-6 3v2l6-2.5V20l-3 2v1l4-1l4 1v-1l-3-2z"
                                        fill-rule="evenodd" fill="currentColor"></path>
                                    <path
                                        d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4z"
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
            <n-layout-content>
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
    top: 10px
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