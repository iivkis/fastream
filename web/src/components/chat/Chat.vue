<template>
    <div class="chat">
        <ChatHead></ChatHead>
        <ChatMessages></ChatMessages>
        <ChatSettings :settings="settings"></ChatSettings>
        <ChatInput v-show="!settings.show"></ChatInput>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { Settings } from "./chat.types";

import ChatHead from "./Head.vue";
import ChatMessages from "./Messages.vue";
import ChatSettings from "./Settings.vue";
import ChatInput from "./Input.vue";

export default defineComponent({
    name: "Chat",
    components: { ChatHead, ChatMessages, ChatSettings, ChatInput },
    setup() {
        const settings = ref<Settings>({
            show: false,
            fieldName: localStorage.getItem("chat.setting.name") || "",
        });

        return {
            settings,
        };
    },
});
</script>
