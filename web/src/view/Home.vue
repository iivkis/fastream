<template>
    <div class="container mx-auto">
        <div class="wrap">
            <div class="welcome">
                <h2 class="welcome__title">Добро пожаловать в Fastream!</h2>
                <img class="welcome__icon" :src="icon" />
            </div>
            <ul class="about">
                <li class="about-item">
                    <h3 class="about-item__title">Создай стрим</h3>
                    <P class="about-item__description">
                        Начни трансляцию своего экрана, вкладки или отдельного
                        приложения в HD формате на одно или несколько устройств
                    </P>
                    <button
                        class="about-item__action"
                        @click="$router.push({ name: 'Stream' })"
                    >
                        Начать трансляцию
                    </button>
                </li>
                <div class="w-0.5 bg-slate-600"></div>
                <li class="about-item">
                    <h3 class="about-item__title">Подключись к трансляции</h3>
                    <P class="about-item__description">
                        Присоединяйся к стриму, чтобы следить за происходящим в
                        реальном времени с минимальными задержками
                    </P>
                    <button
                        class="about-item__action"
                        @click="$router.push({ name: 'Watch' })"
                    >
                        Присоединиться
                    </button>
                </li>
            </ul>
        </div>
    </div>
</template>

<style scoped lang="postcss">
.wrap {
    @apply sm:w-full lg:w-1/2 bg-white px-5 py-7 shadow-sm rounded-sm;
    @apply mt-64 mx-auto;
}

.welcome {
    @apply w-full flex justify-center items-center;
}

.welcome__title {
    @apply text-2xl font-light tracking-wider;
}

.welcome__icon {
    @apply w-10 ml-3;
}

ul.about {
    @apply flex mt-5;
}

.about-item {
    @apply flex flex-col justify-between items-center;
    @apply w-1/2 mx-5;
}

.about-item__title {
    @apply text-2xl font-light;
}

.about-item__description {
    @apply text-center text-sm;
    @apply py-3 text-slate-700;
}

.about-item__action {
    @apply bg-orange-600 text-white;
    @apply w-2/3 px-5 py-3 mt-4 rounded-sm;
    @apply hover:bg-orange-700;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";
import icon from "../assets/icon.png";

export default defineComponent({
    name: "Home",
    setup() {
        return { icon };
    },
});
</script>
