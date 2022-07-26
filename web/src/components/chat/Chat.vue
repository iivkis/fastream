<template>
    <div class="chat">
        <ChatHead></ChatHead>
        <ChatMessages></ChatMessages>
        <div class="chat-settings">
            <div class="chat-settings-switcher">
                <div class="chat-settings-switcher__icon">
                    <i class="bi bi-chevron-up" v-show="settings.show"></i>
                    <i class="bi bi-chevron-down" v-show="!settings.show"></i>
                </div>
                <span
                    class="chat-settings-switcher__title"
                    @click="settings.show = !settings.show"
                >
                    настройки
                </span>
            </div>
            <div class="chat-settings-dropdown" v-show="settings.show">
                <div class="chat-settings-dropdown-item">
                    <label>Ваше имя:</label>
                    <input
                        class="chat-settings-dropdown-item__input"
                        autocomplete="off"
                        placeholder="Имя"
                        v-model.lazy.trim="settings.fieldName"
                        @change="
                            saveToStorage(
                                'chat.setting.name',
                                settings.fieldName
                            )
                        "
                    />
                </div>
            </div>
        </div>
        <div class="chat-input" v-show="!settings.show">
            <div class="chat-input-typing">
                <textarea
                    id="chat-message-input"
                    rows="1"
                    placeholder="Введите сообщение..."
                    maxlength="200"
                    class="chat-input-typing__message"
                    v-model="textareaModel"
                ></textarea>
            </div>
        </div>
    </div>
</template>

<style scoped lang="postcss">
.chat-settings {
    @apply flex flex-col;
}

.chat-settings-switcher {
    @apply flex items-center;
}

.chat-settings-switcher__icon {
    @apply text-sm p-1;
}

.chat-settings-switcher__title {
    @apply cursor-pointer;
}

.chat-settings-dropdown {
    @apply flex flex-col;
}

.chat-settings-dropdown-item {
    @apply flex flex-col;
    @apply mt-2 text-sm;
}

.chat-settings-dropdown-item__input {
    @apply w-1/3 lg:w-full bg-slate-100;
    @apply px-2 py-3 rounded-sm outline-none;
}

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
import { defineComponent, ref } from "vue";
import ChatHead from "./Head.vue";
import ChatMessages from "./Messages.vue";

export default defineComponent({
    name: "Chat",
    components: { ChatHead, ChatMessages },
    setup() {
        const textareaModel = ref("");

        const settings = ref({
            show: false,
            fieldName: localStorage.getItem("chat.setting.name") || "",
        });

        function saveToStorage(name: string, value: string): void {
            localStorage.setItem(name, value);
        }

        return {
            textareaModel,
            settings,
            saveToStorage,
        };
    },

    data: () =>
        ({} as {
            textarea: HTMLElement;
        }),

    mounted() {
        const textarea = document.getElementById("chat-message-input");
        if (textarea) {
            this.$data.textarea = textarea;
        } else {
            throw "undefined textarea";
        }
    },

    watch: {
        textareaModel() {
            this.textarea.style.height = "auto";
            this.textarea.style.height = `${this.textarea.scrollHeight}px`;
            window.scroll(0, document.documentElement.scrollHeight);
        },
    },
});
</script>
