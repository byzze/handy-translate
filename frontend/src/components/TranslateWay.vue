<template>
    <n-space vertical :size="6">
        <n-data-table :bordered="false" :single-line="false" :columns="columns" :data="dataList" />
    </n-space>
</template>
  
<script lang="ts">
import { h, ref, defineComponent, onMounted } from 'vue'
import { NButton, useMessage, DataTableColumns } from 'naive-ui'
import { GetTransalteList, SetTransalteList } from '../../wailsjs/go/main/App'
import {
    ArrowUpCircle,
    ArrowDownCircle
} from '@vicons/ionicons5'

type RowData = {
    index: number
    name: string
}

const createColumns = ({
    up,
    down
}: {
    up: (rowData: RowData) => void
    down: (rowData: RowData) => void
}): DataTableColumns<RowData> => {
    return [
        {
            title: '序号',
            key: 'index'
        },
        {
            title: '翻译方式',
            key: 'name'
        },
        {
            title: '操作',
            key: 'opt',
            render(row) {
                let str = [
                    h(
                        NButton,
                        {
                            style: {
                                marginRight: '3px'
                            },
                            bordered: false,
                            size: 'small',
                            onClick: () => up(row),
                        },
                        { icon: () => h(ArrowUpCircle) }
                    ),
                    h(
                        NButton,
                        {
                            style: {
                                marginRight: '3px'
                            },
                            bordered: false,
                            size: 'small',
                            onClick: () => down(row),
                        },
                        { icon: () => h(ArrowDownCircle) }
                    )]
                const list = str.map((obj) => {
                    return obj
                })
                return list
            },
        }
    ]
}

const createData = (): RowData[] => []

export default defineComponent({
    setup() {
        const DataListRef = ref(createData());
        const message = useMessage()
        // 模拟数据获取
        onMounted(() => {
            GetTransalteList().then(result => {
                let res = JSON.parse(result)
                console.log(res)
                var i = 1
                DataListRef.value = Object.keys(res).map(key => ({
                    index: i++,
                    name: res[key].name,
                    secret: res[key].secret,
                    key: res[key].key,
                    token: res[key].token,
                }));
            })
        });
        return {
            dataList: DataListRef,
            columns: createColumns({
                up(rowData) {
                    console.log(rowData)
                    console.log(DataListRef.value)

                    let newArray = [...DataListRef.value];
                    let index = -1
                    for (let i = 0; i < newArray.length; i++) {
                        if (newArray[i].name == rowData.name) {
                            index = i
                        }
                    }

                    if (index === 0) {
                        message.warning('已经最顶了')
                        return
                    } else {
                        let tmp = newArray[index];
                        newArray[index] = newArray[index - 1];
                        newArray[index - 1] = tmp;
                    }

                    for (let i = 0; i < newArray.length; i++) {
                        newArray[i].index = i + 1
                    }
                    let res = JSON.stringify(newArray)
                    DataListRef.value = newArray
                    // SetTransalteList(res)
                },
                down(rowData) {
                    for (let i = 0; i < DataListRef.value.length; i++) {
                        if (DataListRef.value[i].name == rowData.name) {
                            let tmp: number
                            if (i == DataListRef.value.length - 1) {
                                tmp = 0
                                message.warning('已经最底了')
                                return
                            } else {
                                tmp = i + 1
                            }
                            let tmpData = DataListRef.value[tmp]
                            console.log(tmpData)
                            DataListRef.value[tmp] = rowData
                            DataListRef.value[i] = tmpData
                            break
                        }
                    }
                    for (let i = 0; i < DataListRef.value.length; i++) {
                        DataListRef.value[i].index = i + 1
                    }
                    // SetTransalteList(DataListRef.value.toString())
                }
            }),
            pagination: {
                pageSize: 10
            }
        }
    }
})
</script>