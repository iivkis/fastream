<template>
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
                    @change="toStorage('chat.setting.name', settings.fieldName)"
                />
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
</style>

<script lang="ts">
import { Settings } from "./chat.interface";
import { defineComponent, PropType } from "vue";

export default defineComponent({
    name: "ChatSettings",
    props: {
        settings: {
            required: true,
            type: Object as PropType<Settings>,
        },
    },
    setup() {
        function toStorage(key: string, value: string) {
            localStorage.setItem(key, value);
        }

        return {
            toStorage,
        };
    },
});
</script>
