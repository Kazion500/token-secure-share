<script setup lang="ts">
import Button from "@/components/Button.vue";
import TextArea from "@/components/TextArea.vue";

import { useLinkStore } from "@/stores/linkStore";
import { onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";

const linkStore = useLinkStore();
const router = useRouter();
const route = useRoute();

const copyContent = (content: string) => {
  navigator.clipboard.writeText(content);
};

onMounted(() => {
  if (route.params?.id) {
    linkStore.getLink(route.params?.id as string).catch(() => {
      router.push({
        name: "error",
      });
    });
  }
});
</script>

<template>
  <main class="h-[calc(100vh-70px)] grid place-items-center px-4 md:px-0">
    <div class="container mx-auto grid lg:grid-cols-2 items-center gap-16">
      <section class="h-full py-5">
        <p>
          Below is the secret that was shared with you. Make sure you save it
          somewhere safe. Once you close this page, you will not be able to
          access it again.
        </p>
        <TextArea
          :id="'view'"
          :model-value="linkStore?.link?.content"
          placeholder="Enter your secret here..."
          readonly
          aria-readonly="true"
        />

        <Button
          :text="'Copy Content'"
          @click="() => copyContent(linkStore?.link?.content as string)"
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
