<template>
    <div class="chat-messages">
        <ul class="messages">
            <li
                class="messages-item"
                v-for="(message, index) in messages"
                :key="index"
            >
                <div class="messages-item-wrap">
                    <span class="messages-item-wrap__username">
                        {{ message.username}}
                    </span>
                    <p class="messages-item-wrap__content">
                        {{ message.content }}
                    </p>
                </div>
            </li>
        </ul>
    </div>
</template>

<style scoped lang="postcss">
.chat-messages {
    @apply flex items-end;
    @apply mt-3 h-72;
    @apply border;
}

.messages {
    @apply py-2 overflow-y-auto;
    @apply w-full max-h-full flex flex-col;
}

.messages-item {
    @apply flex justify-start;
    @apply px-2;
}

.messages-item-wrap {
    @apply px-4 py-2.5 my-1;
    @apply w-full bg-white;
}

.messages-item:not(:first-child) {
    @apply border-t;
}

.messages-item-wrap__username {
    @apply text-sm font-medium;
}

.messages-item-wrap__content {
    @apply text-sm font-light;
}
</style>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { WebsocketMessageData } from "../../service/WebsocketChat/types";

export default defineComponent({
    name: "ChatMessages",
    props: {
        messages: {
            required: true,
            type: Object as PropType<WebsocketMessageData[]>,
        },
    },
});
</script>
