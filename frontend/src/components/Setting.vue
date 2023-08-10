<template>
    <div @click="open = !open" class="dot-icon" ref="dropdown">
        <svg viewBox="0 0 30 30" width="24" height="24">
            <circle cx="8" cy="8" r="2.5" />
            <circle cx="16" cy="8" r="2.5" />
            <circle cx="24" cy="8" r="2.5" />
        </svg>
    </div>
    <ul v-show="open" class="dropdown-menu" @mouseenter="onMouseEnter(item)" @mouseleave="onMouseLeave(item)">
        <li v-for="item in list" :key="list.id" @click="clickOpt(item)">
            {{ item.value }}
        </li>
    </ul>
</template>
<script setup>
import { ref } from 'vue'
import { Quit } from '../../wailsjs/runtime/runtime'

document.onkeydown = (e) => {
    console.log(e)
    if (e.key === 'Escape') {
        Quit()
    }
}

const open = ref(false)
let id = 0
const list = ref([{ id: id++, value: '退出       ESC' }])

function clickOpt(item) {
    console.log(item)
    if (item.id == 0) {
        Quit()
    }
}
function onMouseEnter(item) {
    console.log("进入")
}

function onMouseLeave(item) {
    console.log("离开", open.value)
    open.value = false
}
</script>
<style>
.dropdown-menu {
    position: absolute;
    padding: 0 0;
    background: #fff;
    color: black;
    list-style-type: none;
    border: 1px solid #ddd;
    border-radius: 16%;
    top: 5px;
    font-size: 10px;
}

.dropdown-menu li {
    padding: 5px 10px;
}

.dropdown-menu li:hover {
    background: rgb(159, 159, 177);
}

.dot-icon {
    position: absolute;
    top: 5px;
    left: 5px;
}

circle {
    fill: white;
}

circle:hover {
    fill: #f60;
}
</style>