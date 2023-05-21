<script setup lang="ts">
import Button from "@/components/Button.vue";
import { useRouter } from "vue-router";

const router = useRouter();

const goBackHome = () => {
  router.push({
    name: "home",
  });
};
</script>

<template>
  <main class="h-[calc(100vh-100px)] grid place-items-center text-center px-5 md:px-0">
    <section class="container mx-auto">
      <div class="border px-5 py-10 lg:w-[50%] mx-auto rounded-md">
        <h1 class="font-bold text-3xl text-red-500">
          Error: While trying to view link
        </h1>
        <p class="my-7 w-[90%] mx-auto">
          Possible reasons for this error message could be that the link has
          already been accessed or that the link does not exist. In such cases,
          we recommend reaching out to the person who shared the link and kindly
          requesting them to generate a new link for you.
        </p>

        <Button :text="'Back Home'" @click="goBackHome" />
      </div>
    </section>
  </main>
</template>
