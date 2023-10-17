<template>
    <n-radio-group @change="handleChange" v-model:value="translateWay" name="radiogroup">
        <n-space>
            <n-radio v-for="song in dataList" :key="song.label" :value="song.label">
                {{ song.name }}
            </n-radio>
        </n-space>
    </n-radio-group>
</template>
  
<script lang="ts">
import { h, ref, defineComponent, onMounted } from 'vue'
import { NButton, useMessage, DataTableColumns } from 'naive-ui'
import { GetTransalteMap, SetTransalteWay, GetTransalteWay } from '../../wailsjs/go/main/App'
type RowData = {
    name: string
    label: string
}
const createData = (): RowData[] => []
export default defineComponent({
    setup() {
        const dataListRef = ref(createData());
        const translateWayRef = ref('');
        const message = useMessage()
        onMounted(() => {
            GetTransalteMap().then(result => {
                let res = JSON.parse(result)
                console.log(res)
                dataListRef.value = Object.keys(res).map(key => ({
                    name: res[key].name,
                    label: key,
                }));
                console.log(dataListRef.value)
            })

            GetTransalteWay().then(result => {
                console.log(result)
                translateWayRef.value = result
            })
        });
        return {
            translateWay: translateWayRef,
            dataList: dataListRef,
            handleChange() {
                console.log(translateWayRef.value)
                SetTransalteWay(translateWayRef.value)
            }
        }
    }
})
</script>
