<template>
    <div>
        <n-grid :x-gap="12" :y-gap="12" :cols="1" layout-shift-disabled>
            <n-gi>
                <n-space item-style="display: flex;" vertical>
                    <n-checkbox-group :value="cities" @update:value="handleUpdateValue">
                        <n-space item-style="display: flex;" align="center">
                            <n-checkbox value="ctrl" label="Ctrl" />
                            <n-checkbox value="shift" label="Shift" />
                            <n-input style="width: 50%" placeholder="基本按钮" @keydown="captureKeyPress"
                                v-model:value="shortcutKey" />
                        </n-space>
                    </n-checkbox-group>
                </n-space>
            </n-gi>
        </n-grid>
        <n-grid :x-gap="12" :y-gap="12" :cols="1" layout-shift-disabled>
            <n-space>
                <n-button type="warning" @click="reset" ghost>
                    重置
                </n-button>
                <n-button type="success" @click="save" ghost>
                    保存
                </n-button>
            </n-space>
        </n-grid>
    </div>
</template>
  
<script lang="ts">
import { onMounted, defineComponent, ref } from 'vue'; // 导入 Vue 3 的 ref
import { GetKeyBoard, SetKeyBoard } from '../../wailsjs/go/main/App'
import { useMessage } from 'naive-ui'
export default defineComponent({
    setup() {
        const shortcutKey = ref('');
        const shortcutInput = ref(null);
        const citiesRef = ref(["", "", ""]);
        const showKeyModalRef = ref(false);
        const message = useMessage()
        const captureKeyPress = (event) => {
            // 获取按下的键
            if (event.key == "Backspace") {
                return
            }
            const validKeys = /^[a-zA-Z0-9]+$/;
            if (event.key == " ") {
                shortcutKey.value = event.code;
            }
            if (validKeys.test(event.key) && event.key.length === 1) {
                shortcutKey.value = event.key;
            }
            // 阻止事件继续传播，防止多次触发
            event.preventDefault();
            // 禁止使用组合键，以防止误触发
            // shortcutInput.value.blur();
        };

        // 模拟数据获取
        onMounted(() => {
            GetKeyBoard().then(res => {
                citiesRef.value = res
                console.log(res)
                shortcutKey.value = citiesRef.value[2]
            })
        })

        return {
            cities: citiesRef,
            showKeyModal: showKeyModalRef,
            shortcutKey,
            shortcutInput,
            captureKeyPress,
            handleUpdateValue(value) {
                citiesRef.value = value;
            },
            reset() {
                citiesRef.value = ["", "", ""]
                shortcutKey.value = ""
                message.success('已重置')
            },
            save() {
                let ctrl = false;
                let shift = false;

                if (shortcutKey.value == "") {
                    message.warning('快捷键配置为空')
                    return
                }
                for (const key of citiesRef.value) {
                    if (key === "ctrl") {
                        ctrl = true;
                    }
                    if (key === "shift") {
                        shift = true;
                    }
                }

                if (!ctrl && !shift) {
                    SetKeyBoard("center", "", "");
                } else {
                    SetKeyBoard(ctrl ? "ctrl" : "center", shift ? "shift" : "", shortcutKey.value);
                }
                message.success('保存成功')
            },
        };
    },
})

</script>
  