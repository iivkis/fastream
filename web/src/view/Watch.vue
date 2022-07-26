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
            <Chat></Chat>
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
</style>

<script lang="ts">
import { defineComponent } from "vue";
import { WebsocketStreamWatcher } from "../service/WebsocketStream/WebsocketStreamWatcher";

import Chat from "../components/chat/Chat.vue";

export default defineComponent({
    name: "Watch",
    components: { Chat },


    mounted() {
        const video = document.getElementById("watch-video");
        if (video) {
            new WebsocketStreamWatcher(video as HTMLMediaElement);
        } else {
            console.error("video tag undefined");
        }
    },
});
</script>
