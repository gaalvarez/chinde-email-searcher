<template>
  <div class="bg-white p-20 border-t border-gray-300">
    <div class="flex flex-col items-center justify-center">
      <div>
        <img
          alt="chinde logo"
          src="@/assets/images/empty.png"
          width="80"
          height="80"
        />
      </div>
      <p class="text-center text-gray-500 font-medium">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  messageEmailSearch: { message: string };
}>();

const message = computed(() => {
  return props.messageEmailSearch.message === "noresult"
    ? "No results found for your search"
    : "Please input a search term";
});
</script>
