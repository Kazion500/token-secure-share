<script setup lang="ts">
defineProps({
  id: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <textarea
    :id="id"
    cols="30"
    rows="10"
    class="border w-full rounded-md resize-none my-4 p-4 shadow"
  />
</template>
