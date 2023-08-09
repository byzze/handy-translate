<template>
    <div class="app-container">
        <!-- Your application content -->

        <!-- Minimize button -->
        <span class="minimize-button" @click="minimizeApp">-</span>
    </div>
</template>
  
<script>
export default {
    methods: {
        minimizeApp() {
            // Logic to minimize the application
            // You can modify the CSS styles, use a Vue state to track the minimized state, etc.
        }
    }
}
</script>
  
<style>
.app-container {
    position: absolute;
    /* Ensure no extra margins or padding */
    margin: 0;
    padding: 0;
    top: 0;
    /* Button will be at the top edge of the container */
    right: 0;
}

.minimize-button {
    padding: 5px 10px;
    font-size: 20px;
    color: white;
    background-color: black;
    border: none;
    cursor: pointer;
}
</style>
  