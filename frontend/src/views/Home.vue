<script setup lang="ts">
import Button from "@/components/Button.vue";
import TextArea from "@/components/TextArea.vue";

import { useLinkStore } from "@/stores/linkStore";
import { ref } from "vue";
import { useRouter } from "vue-router";

const secret = ref("");
const linkStore = useLinkStore();
const route = useRouter();

const generateLink = () => {
  linkStore
    .createLink(secret.value)
    .then((res: any) => {
      route.push({
        name: "generate-link",
        params: {
          id: res.data.link,
        },
      });
    })
    .catch((err) => {
      console.log(err);
    });
};
</script>

<template>
  <main class="h-[calc(100vh-70px)] grid place-items-center px-4 md:px-0">
    <div class="container mx-auto grid lg:grid-cols-2 items-center gap-16">
      <section class="h-full py-5">
        <p>
          Saving password in plain text is not a good idea. This app will help
          you to encrypt your password and save it in your local storage.
        </p>
        <TextArea
          :id="'create'"
          v-model="secret"
          placeholder="Enter your secret here..."
        />
        <Button :text="'Generate Link'" @click="generateLink" />
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
