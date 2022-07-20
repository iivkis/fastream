<template>
    <div class="container mx-auto">
        <div class="wrap">
            <div class="welcome">
                <h2>Добро пожаловать в Fastream!</h2>
                <img class="wrap-icon" :src="icon" />
            </div>
            <ul class="about">
                <li>
                    <h3>Создай стрим</h3>
                    <P>
                        Начни трансляцию своего экрана, вкладки или отдельного
                        приложения в HD формате на одно или несколько устройств
                    </P>
                    <button @click="$router.push({name:'Stream'})">Начать трансляцию</button>
                </li>
                <div class="w-0.5 bg-slate-600"></div>
                <li>
                    <h3>Подключись к трансляции</h3>
                    <P>
                        Присоединяйся к стриму, чтобы следить за происходящим в
                        реальном времени с минимальными задержками
                    </P>
                    <button @click="$router.push({name:'Watch'})">Присоединиться</button>
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

.wrap-icon {
    @apply w-10 ml-3;
}

.welcome {
    @apply w-full flex justify-center items-center;
}

.welcome > h2 {
    @apply text-2xl font-light tracking-wider;
}

ul.about {
    @apply flex mt-5;
}

ul.about > li {
    @apply flex flex-col justify-between items-center;
    @apply w-1/2 mx-5;
}

ul.about > li > h3 {
    @apply text-2xl font-light;
}

ul.about > li > p {
    @apply text-center text-sm;
    @apply py-3 text-slate-700;
}

ul.about > li > button {
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
