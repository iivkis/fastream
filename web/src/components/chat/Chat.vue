<template>
    <div class="chat">
        <ChatHead></ChatHead>
        <ChatMessages :messages="messages"></ChatMessages>
        <ChatSettings :settings="settings"></ChatSettings>
        <ChatInput v-show="!settings.show"></ChatInput>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { Settings } from "./types";

import ChatHead from "./Head.vue";
import ChatMessages from "./Messages.vue";
import ChatSettings from "./Settings.vue";
import ChatInput from "./Input.vue";
import { WebsocketChat } from "../../service/WebsocketChat";
import { WebsocketMessageData } from "../../service/WebsocketChat/types";

export default defineComponent({
    name: "Chat",
    components: { ChatHead, ChatMessages, ChatSettings, ChatInput },
    setup() {
        const settings = ref<Settings>({
            show: false,
            fieldName: localStorage.getItem("chat.setting.name") || "",
        });

        const messages = ref<WebsocketMessageData[]>([
            {
                username: "Система",
                content: `Добро пожаловать в чат! Здесь ты можешь обмениваться сообщениями со зрителями. 
                    Подсказка: открой настройки, чтобы изменменить имя, которое будет отображаться в сообщении`,
            },
        ]);

        const chat = new WebsocketChat();

        chat.OnMessage = function (message) {
            messages.value.push(message);
        };

        

        return {
            settings,
            messages,
        };
    },
});
</script>
