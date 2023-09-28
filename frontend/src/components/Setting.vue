<template>
    <n-dropdown :options="options" @select="handleSelect">
        <n-button>
            <n-icon>
                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32">
                    <g fill="none">
                        <path
                            d="M16 11a5 5 0 1 0 0 10a5 5 0 0 0 0-10zm-3 5a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm-.16 13.628c1.035.247 2.096.372 3.16.372a13.643 13.643 0 0 0 3.156-.375a1.478 1.478 0 0 0 1.13-1.276l.234-2.13a1.471 1.471 0 0 1 2.066-1.2l1.955.856a1.472 1.472 0 0 0 1.671-.345a14.245 14.245 0 0 0 3.156-5.443a1.478 1.478 0 0 0-.535-1.627l-1.729-1.275a1.481 1.481 0 0 1 .003-2.396l1.72-1.27a1.474 1.474 0 0 0 .537-1.63a14.199 14.199 0 0 0-3.157-5.443a1.48 1.48 0 0 0-1.674-.345l-1.946.856a1.483 1.483 0 0 1-2.067-1.2l-.236-2.12a1.476 1.476 0 0 0-1.147-1.283a15.123 15.123 0 0 0-3.127-.363a15.395 15.395 0 0 0-3.146.363a1.469 1.469 0 0 0-1.147 1.28l-.237 2.122a1.493 1.493 0 0 1-2.073 1.206l-1.946-.857a1.493 1.493 0 0 0-1.67.35a14.245 14.245 0 0 0-3.16 5.446a1.478 1.478 0 0 0 .536 1.625l1.725 1.272a1.488 1.488 0 0 1 0 2.397L3.167 18.47a1.477 1.477 0 0 0-.535 1.63a14.253 14.253 0 0 0 3.16 5.45a1.458 1.458 0 0 0 1.077.465c.203 0 .404-.042.591-.123l1.955-.859a1.485 1.485 0 0 1 2.065 1.2l.235 2.126a1.476 1.476 0 0 0 1.125 1.27zm5.501-1.866a11.638 11.638 0 0 1-4.677 0l-.195-1.74a3.48 3.48 0 0 0-1.14-2.208a3.534 3.534 0 0 0-3.718-.6l-1.606.7a12.237 12.237 0 0 1-2.348-4.05l1.424-1.052a3.488 3.488 0 0 0 0-5.616L4.66 12.147a12.243 12.243 0 0 1 2.348-4.046l1.6.7a3.45 3.45 0 0 0 1.4.294a3.5 3.5 0 0 0 3.467-3.108l.194-1.747c.774-.15 1.56-.23 2.347-.24c.782.01 1.562.09 2.33.24l.186 1.74a3.48 3.48 0 0 0 1.137 2.216a3.525 3.525 0 0 0 3.727.6l1.6-.7a12.212 12.212 0 0 1 2.35 4.047l-1.423 1.046a3.48 3.48 0 0 0 0 5.62l1.422 1.05A12.273 12.273 0 0 1 25 23.901l-1.6-.7a3.473 3.473 0 0 0-4.866 2.81l-.193 1.75z"
                            fill="currentColor"></path>
                    </g>
                </svg>
            </n-icon>
        </n-button>
    </n-dropdown>

    <n-modal v-model:show="showKeyModal">
        <n-card style="width: auto" title="快捷键" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <p>触发按键，点击保存, 默认是鼠标中键</p>
            <ShortKey />
        </n-card>

    </n-modal>

    <n-modal v-model:show="showModal">
        <n-card style="width: auto" title="翻译" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <TranslateWay />
        </n-card>
    </n-modal>

    <n-modal v-model:show="showModalAbout">
        <n-card style="width: auto; text-align: left;" title="便捷翻译" :bordered="false" size="huge" role="dialog"
            aria-modal="true">
            这是一款简单-便捷的翻译工具
            <div style="display: flex; align-items: center;">
                <n-icon size="27">
                    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24">
                        <path
                            d="M9 19c-4.3 1.4-4.3-2.5-6-3m12 5v-3.5c0-1 .1-1.4-.5-2c2.8-.3 5.5-1.4 5.5-6a4.6 4.6 0 0 0-1.3-3.2a4.2 4.2 0 0 0-.1-3.2s-1.1-.3-3.5 1.3a12.3 12.3 0 0 0-6.2 0C6.5 2.8 5.4 3.1 5.4 3.1a4.2 4.2 0 0 0-.1 3.2A4.6 4.6 0 0 0 4 9.5c0 4.6 2.7 5.7 5.5 6c-.6.6-.6 1.2-.5 2V21"
                            fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round">
                        </path>
                    </svg>
                </n-icon>
                <span style="margin-left: 10px;">github.com/byzze/handy-translate</span>
            </div>
            <div style="display: flex; align-items: center;">
                <n-icon size="27">
                    <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32">
                        <path
                            d="M28 6H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm-2.2 2L16 14.78L6.2 8zM4 24V8.91l11.43 7.91a1 1 0 0 0 1.14 0L28 8.91V24z"
                            fill="currentColor"></path>
                    </svg>
                </n-icon>
                <span style="margin-left: 10px;">luoyd@163.com</span>
            </div>
            <template #footer>

            </template>
        </n-card>
    </n-modal>
</template>
<style scoped>
.light-green {
    height: 108px;
    background-color: rgba(0, 128, 0, 0.12);
}

.green {
    height: 108px;
    background-color: rgba(0, 128, 0, 0.24);
}
</style>
<script>
import { h, defineComponent, ref, onMounted } from 'vue'
import { NIcon } from 'naive-ui'
import TranslateWay from "./TranslateWay.vue";
import ShortKey from "./ShortKey.vue";

import {
    Pencil as EditIcon,
    LogOutOutline as LogoutIcon
} from '@vicons/ionicons5'
import {
    ExclamationCircleOutlined as AboutIcon,
} from '@vicons/antd'
import {
    Keyboard as Keyboard,
} from '@vicons/tabler'
import { useMessage, useDialog } from 'naive-ui'
import { Hide, Quit } from '../../wailsjs/runtime/runtime'

const renderIcon = (icon) => {
    return () => {
        return h(NIcon, null, {
            default: () => h(icon)
        })
    }
}

export default defineComponent({
    components: {
        TranslateWay,
        ShortKey,
    },
    setup() {
        const showModalRef = ref(false);
        const showModalAboutRef = ref(false);
        const showKeyModalRef = ref(false);
        const transalteWayRef = ref("");
        const editableTextRef = ref([]);
        const transalteWayDataRef = ref([]);
        const message = useMessage()
        const dialog = useDialog()
        const checkedRef = ref(false);
        const citiesRef = ref(["", "", ""]);

        // 模拟数据获取
        onMounted(() => {
            document.onkeydown = function (e) {
                switch (e.key) {
                    case "Escape":
                        Hide()
                }
            }
        })
        return {
            editableText: editableTextRef,
            checkedValue: transalteWayRef,
            showModal: showModalRef,
            showKeyModal: showKeyModalRef,
            showModalAbout: showModalAboutRef,
            translateList: transalteWayDataRef,
            checked: checkedRef,
            cities: citiesRef,
            selectedValue: ref(''),
            options: [
                {
                    label: '配置翻译',
                    key: 'config',
                    icon: renderIcon(EditIcon)
                },
                {
                    label: '快捷键',
                    key: 'keyboard',
                    icon: renderIcon(Keyboard)
                },
                {
                    label: '退出',
                    key: 'logout',
                    icon: renderIcon(LogoutIcon)
                },
                {
                    label: '关于',
                    key: 'about',
                    icon: renderIcon(AboutIcon)
                }
            ],

            handleSelect(key) {
                showModalRef.value = false
                showModalAboutRef.value = false
                switch (key) {
                    case 'logout':
                        dialog.warning({
                            title: '警告',
                            content: '确定退出吗？',
                            positiveText: '确定',
                            negativeText: '取消',
                            onPositiveClick: () => {
                                Quit()
                            },
                            onNegativeClick: () => {
                                message.warning('已取消')
                            }
                        })
                        // It's important to add a "break" statement here to exit the switch after this case.
                        break;
                    case 'config':
                        showModalRef.value = true
                        break;
                    case 'keyboard':
                        showKeyModalRef.value = true
                        break;
                    case 'about':
                        showModalAboutRef.value = true
                        break;
                }
            }
        }
    }
})
</script>