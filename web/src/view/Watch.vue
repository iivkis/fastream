<template>
    <div class="container mx-auto">
        <div class="wrap">
            <div class="broadcast">
                <video id="watch-video" autoplay muted controls></video>
            </div>
            <div class="chat">
                <div class="chat__head">
                    <img :src="icon" />
                    <h3>Fastream | Чат</h3>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="postcss">
.wrap {
    @apply lg:w-2/3 flex mx-auto mt-40;
    @apply bg-white p-2;
}

.broadcast {
    @apply w-4/6;
}

.broadcast video {
    @apply w-full;
}

.chat {
    @apply w-2/6 mt-2 flex flex-col;
}

.chat__head {
    @apply flex justify-center items-center;
}

.chat__head img {
    @apply w-7 mr-2;
}

.chat__head h3 {
    @apply text-xl font-light;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";

// import img
import icon from "../assets/icon.png";

export default defineComponent({
    name: "Watch",
    setup() {
        return {
            icon,
        };
    },
});
</script>
