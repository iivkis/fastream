<template>
    <div class="chat-input">
        <div class="chat-input-typing">
            <textarea
                id="chat-message-input"
                rows="1"
                placeholder="Введите сообщение..."
                maxlength="200"
                class="chat-input-typing__message"
                v-model="message"
            ></textarea>
        </div>
    </div>
</template>

<style scoped lang="postcss">
.chat-input {
    @apply w-full;
}

.chat-input-typing {
    @apply w-full  bg-slate-100 rounded-sm overflow-hidden;
}

.chat-input-typing__message {
    @apply p-2 bg-transparent;
    @apply w-full outline-none;
    @apply overflow-clip text-sm;
}
</style>

<script lang="ts">
import { defineComponent, PropType, ref } from "vue";

export default defineComponent({
    name: "ChatInput",

    props: {
        send: {
            type: Function,
        }
    },

    setup() {
        const message = ref("");
        return { message };
    },

    data: function () {
        return {} as {
            textarea: HTMLElement;
        };
    },

    mounted() {
        const textarea = document.getElementById("chat-message-input");
        if (textarea) {
            this.$data.textarea = textarea;
        } else {
            throw "undefined textarea";
        }
    },

    watch: {
        message() {
            this.textarea.style.height = "auto";
            this.textarea.style.height = `${this.textarea.scrollHeight}px`;
            window.scroll(0, document.documentElement.scrollHeight);
        },
    },
});
</script>
