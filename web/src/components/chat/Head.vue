<template>
    <div class="chat-head">
        <img class="chat-head__icon" :src="icon" />
        <h3 class="chat-head__title">Fastream | Чат</h3>
    </div>
</template>

<style scoped lang="postcss">
.chat-head {
    @apply flex justify-center items-center;
}

.chat-head__icon {
    @apply w-7 mr-2;
}

.chat-head__title {
    @apply text-xl font-light;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";
import icon from "../../assets/icon.png";

export default defineComponent({
    name: "ChatHead",
    setup() {
        return { icon };
    },
});
</script>
