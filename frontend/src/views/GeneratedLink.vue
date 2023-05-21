<script setup lang="ts">
import Button from "@/components/Button.vue";

import { onMounted, ref } from "vue";

import { useRoute, useRouter } from "vue-router";

const fullLink = ref("");
const router = useRouter();
const route = useRoute();

const copyLink = () => {
  navigator.clipboard.writeText(fullLink.value);
};

onMounted(() => {
  if (!route.params?.id) {
    router.push({
      name: "home",
    });
  } else {
    fullLink.value = `${window.location.origin}/${route.params?.id}`;
  }
});
</script>

<template>
  <main class="h-[calc(100vh-70px)] px-4 md:px-0">
    <section
      class="container mx-auto justify-center items-center h-full flex flex-col"
    >
      <h1 class="text-lg uppercase font-bold">Link generate</h1>
      <p class="py-5 text-center">
        This is the link you can share with your friends and team workers.
      </p>

      <input
        type="text"
        :value="fullLink"
        class="p-3 mb-5 outline-none text-gray-500 bg-gray-100 border rounded-md overflow-hidden md:w-[70%] lg:w-[40%] w-full flex justify-between items-center"
        readonly
      />

      <Button :text="'Copy Link'" @click="copyLink" />
    </section>
  </main>
</template>

<style scoped></style>
