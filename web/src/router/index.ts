import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Home from "../view/Home.vue";
import Stream from "../view/Stream.vue";
import Watch from "../view/Watch.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path:"/stream",
        name: "Stream",
        component: Stream
    },
    {
        path:"/watch",
        name: "Watch",
        component: Watch
    }
];

export default createRouter({
    history: createWebHistory(),
    routes,
});
