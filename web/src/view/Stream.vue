<template>
    <div class="container mx-auto">
        <div class="wrap">
            <div class="head">
                <div class="head-info">
                    <img class="head-info__icon" :src="icon" />
                    <h2 class="head-info__text">Трансляция запущена</h2>
                </div>
                <div class="head-notice">
                    <h5 class="head-notice__text">
                        <span>Примечание:</span> трансляция будет прекращена,
                        как только вы закроете текущую вкладку браузера
                    </h5>
                </div>
            </div>

            <div class="main">
                <div class="qr-code">
                    <img
                        src="https://expertnov.ru/800/600/https/upload.wikimedia.org/wikipedia/commons/7/78/Qrcode_wikipedia_fr_v2clean.png"
                    />
                </div>

                <div class="ip-addr">
                    <span class="ip-addr__instruction">
                        Наведи камеру смартфона на QR-код или используй ссылку,
                        чтобы подключиться к трансляции:
                    </span>
                    <div class="ip-addr-link">
                        <router-link
                            class="ip-addr-link__href"
                            to="/watch"
                            target="_blank"
                        >
                            {{ url }}
                        </router-link>
                        <i
                            class="ip-addr-link__copy bi bi-clipboard"
                            title="копировать ссылку"
                            @click="CopyURLToClipboard"
                        ></i>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="postcss">
.wrap {
    @apply bg-white w-4/5 lg:w-2/5 mx-auto mt-40 py-8 rounded-sm;
}

.head {
    @apply flex flex-col;
}

.head-info {
    @apply flex justify-center items-center;
}

.head-info__icon {
    @apply w-8 mr-3;
}

.head-info__text {
    @apply text-3xl font-light;
}

.head-notice__text {
    @apply w-2/3 py-2 mx-auto text-slate-600 text-center font-light text-sm;
}

.head-notice__text span {
    @apply text-orange-600;
}

.main {
    @apply flex flex-col mt-5;
}

.qr-code {
    @apply flex flex-col items-center;
}

.qr-code img {
    @apply w-64;
}

.ip-addr {
    @apply flex flex-col px-10;
}

.ip-addr__instruction {
    @apply py-4 text-center text-slate-600 tracking-tight;
}

.ip-addr-link {
    @apply flex justify-center;
}

.ip-addr-link__href {
    @apply text-center text-blue-600 text-2xl hover:underline;
}

.ip-addr-link__copy {
    @apply text-xl text-slate-500 hover:text-slate-900 cursor-pointer ml-2;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";
import copy from "copy-to-clipboard";

// import img
import icon from "../assets/icon.png";
import { WebsocketStreamCreator } from "../service/WebsocketStreamCreator";

export default defineComponent({
    name: "Stream",
    async setup() {
        const url = "http://192.168.1.1:8080/watch";

        const config: DisplayMediaStreamConstraints = {
            audio: false,
            video: {
                width: {
                    ideal: 1366,
                    max: 1920,
                },
                height: {
                    ideal: 768,
                    max: 1080,
                },
                frameRate: 30,
                latency: 0,
            },
        };

        const stream = await navigator.mediaDevices.getDisplayMedia(config);

        new WebsocketStreamCreator(stream);

        function CopyURLToClipboard() {
            copy(url);
        }

        return {
            icon,
            url,
            CopyURLToClipboard,
        };
    },
});
</script>
