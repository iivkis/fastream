<template>
    <suspense>
        <router-view></router-view>
    </suspense>
</template>

<style lang="postcss">
body {
    @apply bg-slate-200;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    name: "App",
});
</script>
