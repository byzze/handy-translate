<template>
    <div>
        <NList>
            <NListItem v-for="(item, index) in items" :key="index" :draggable="true" @dragstart="startDrag(index)"
                @dragover="allowDrop" @drop="endDrag(index)" @dragenter="setDropTarget(index)" @dragleave="unsetDropTarget"
                :class="{ 'drag-over': isDragOver && index === dropTarget }">
                {{ item }}
            </NListItem>
        </NList>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            items: ['Item 1', 'Item 2', 'Item 3', 'Item 4', 'Item 5'],
            isDragging: false,
            dragIndex: null,
            isDragOver: false,
            dropTarget: null,
        };
    },
    methods: {
        startDrag(index) {
            this.isDragging = true;
            this.dragIndex = index;
        },
        allowDrop(event) {
            event.preventDefault();
        },
        endDrag(index) {
            if (this.isDragging) {
                if (this.dropTarget !== null && this.dragIndex !== this.dropTarget) {
                    // Reorder the items
                    const itemToMove = this.items[this.dragIndex];
                    this.items.splice(this.dragIndex, 1);
                    this.items.splice(this.dropTarget, 0, itemToMove);
                }
                this.isDragging = false;
                this.dragIndex = null;
                this.isDragOver = false;
                this.dropTarget = null;
            }
        },
        setDropTarget(index) {
            if (this.isDragging && index !== this.dragIndex) {
                this.isDragOver = true;
                this.dropTarget = index;
            }
        },
        unsetDropTarget() {
            this.isDragOver = false;
            this.dropTarget = null;
        },
    },
};
</script>
  
<style>
.drag-over {
    border: 2px dashed #007bff;
}
</style>
  