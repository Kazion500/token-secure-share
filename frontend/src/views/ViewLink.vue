<script setup lang="ts">
import Button from "@/components/Button.vue";
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
        class="p-3 mb-5 outline-none border rounded-md overflow-hidden bg-white md:w-[70%] lg:w-[40%] w-full flex justify-between items-center"
        readonly
      />

      <Button :text="'Copy Link'" />
    </section>
  </main>
</template>

<style scoped></style>
