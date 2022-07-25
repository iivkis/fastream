<template>
    <div class="container mx-auto">
        <div class="wrap">
            <div class="broadcast">
                <video
                    id="watch-video"
                    class="broadcast__video"
                    muted
                    autoplay
                    controls
                ></video>
            </div>
            <div class="chat">
                <div class="chat-head">
                    <img class="chat-head__icon" :src="icon" />
                    <h3 class="chat-head__title">Fastream | Чат</h3>
                </div>
                <div class="chat-messages">
                    <ul class="messages">
                        <li
                            class="messages-item"
                            v-for="(message, index) in messages"
                            :key="index"
                        >
                            <div class="messages-item-wrap">
                                <span class="messages-item-wrap__username">
                                    {{ message.username }}
                                </span>
                                <p class="messages-item-wrap__content">
                                    {{ message.content }}
                                </p>
                            </div>
                        </li>
                    </ul>
                </div>
                <div class="chat-settings">
                    <div class="chat-settings-switcher">
                        <div class="chat-settings-switcher__icon">
                            <i
                                class="bi bi-chevron-up"
                                v-show="settings.show"
                            ></i>
                            <i
                                class="bi bi-chevron-down"
                                v-show="!settings.show"
                            ></i>
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
        </div>
    </div>
</template>

<style scoped lang="postcss">
.wrap {
    @apply flex flex-col lg:flex-row;
    @apply p-2 mx-auto lg:mt-32;
    @apply w-full lg:w-2/3;
    @apply bg-white;
}

.broadcast {
    @apply w-full lg:w-4/6;
}

.broadcast__video {
    @apply w-full;
}

.chat {
    @apply w-full lg:w-2/6 flex flex-col;
    @apply mt-3 px-2 lg:mt-2 lg:ml-2;
}

.chat-head {
    @apply flex justify-center items-center;
}

.chat-head__icon {
    @apply w-7 mr-2;
}

.chat-head__title {
    @apply text-xl font-light;
}

.chat-messages {
    @apply flex items-end;
    @apply mt-3 h-60;
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
    @apply p-3 bg-transparent;
    @apply w-full outline-none;
    @apply overflow-clip text-sm;
}
</style>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { WebsocketStreamWatcher } from "../service/WebsocketStream/WebsocketStreamWatcher";
import { WebsocketMessageData } from "../service/WebsocketChat";

import icon from "../assets/icon.png";

export default defineComponent({
    name: "Watch",
    setup() {
        const textareaModel = ref("");

        const settings = ref({
            show: false,
            fieldName: localStorage.getItem("chat.setting.name") || "",
        });

        var messages = ref([
            {
                username: "Система",
                content:
                    "Добро пожаловать в чат! Здесь ты можешь обмениваться сообщениями со зрителями. Подсказка: открой настройки, чтобы изменменить имя, которое будет отображаться в сообщении",
            },
            // {
            //     username: "Georgy",
            //     content: "float lore da bucaib csbia cnac;nisvbdua cnaiopna",
            // },
            // {
            //     username: "Lolita",
            //     content:
            //         "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Odio sint veritatis expedita suscipit! Voluptatem neque consequuntur laudantium. Nihil, odio voluptatum.",
            // },
            // {
            //     username: "Georgy",
            //     content: "https://vk.com/im?sel=c103",
            // },
            // {
            //     username: "Lolita",
            //     content:
            //         "Lorem dolor sit amet, consectetur adipisicing elit. Odio sint veritatis expedita suscipit! Voluptatem neque consequuntur laudantium. Nihil, odio voluptatum.",
            // },
        ] as Array<WebsocketMessageData>);

        function saveToStorage(name: string, value: string): void {
            localStorage.setItem(name, value);
        }

        return {
            icon,
            textareaModel,
            settings,
            messages,
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

        const video = document.getElementById("watch-video");
        if (video) {
            new WebsocketStreamWatcher(video as HTMLMediaElement);
        } else {
            console.error("video tag undefined");
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
