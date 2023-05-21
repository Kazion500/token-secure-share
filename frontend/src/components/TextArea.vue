<script setup lang="ts">
defineProps({
  id: {
    type: String,
    required: true,
  },
  modelValue: {
    type: String,
  },
});
defineEmits(["update:modelValue"]);
</script>

<template>
  <textarea
    :id="id"
    cols="30"
    rows="10"
    class="border w-full rounded-md resize-none my-4 p-4 shadow"
    :value="modelValue"
    @input="$emit('update:modelValue', ($event.target as any).value)"
  />
</template>
