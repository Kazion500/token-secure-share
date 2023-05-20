<script setup lang="ts">
import Button from "@/components/Button.vue";
import TextArea from "@/components/TextArea.vue";

import { ref } from "vue";

const data = ref("");
</script>

<template>
  <main class="h-[calc(100vh-70px)] grid place-items-center px-4 md:px-0">
    <div class="container mx-auto grid lg:grid-cols-2 items-center gap-16">
      <section class="h-full py-5">
        <p>
          Saving password in plain text is not a good idea. This app will help
          you to encrypt your password and save it in your local storage.
        </p>
        <TextArea :id="'create'" v-model="data" />
        <Button
          :text="'Generate Link'"
          @click="
            () =>
              $router.push({
                name: 'generate-link',
                params: {
                  id: 2,
                },
              })
          "
        />
      </section>

      <section class="h-full py-5">
        <h3 class="font-bold uppercase mb-2">How it works</h3>
        <p>
          This app will help you to securely share your secrets with your
          friends and team workers.
        </p>
        <p class="mt-4">
          One thing to note is once the link is generated and shared and it is
          opened, it will be deleted from the server and you will not be able to
          use it or anyone with access . So, make sure you save the link
          somewhere safe.
        </p>
      </section>
    </div>
  </main>
</template>
