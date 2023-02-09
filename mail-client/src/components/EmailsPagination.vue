<template>
  <div class="flex justify-between">
    <span></span>
    <div class="flex gap-x-1">
      <button @click="onPreviousPage">
        <svg
          class="previous-icon"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
        >
          <path d="M15.41 16.09l-4.58-4.59 4.58-4.59L14 5.5l-6 6 6 6z" />
          <path d="M0-.5h24v24H0z" fill="none" />
        </svg>
      </button>
      <div>Page {{ getCurrentPage }} of {{ totalPages }}</div>
      <button @click="onNextPage">
        <svg
          class="next-icon"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
        >
          <path d="M8.59 16.34l4.58-4.59-4.58-4.59L10 5.75l6 6-6 6z" />
          <path d="M0-.25h24v24H0z" fill="none" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from "pinia";
import { computed } from "vue";
import { useMailStore } from "@/stores/mail.store";

const props = defineProps<{
  totalRecords: number;
  pageSize: number;
}>();

const emit = defineEmits<{
  (e: "pagination:change", value: { currentPage: number }): void;
}>();

const { getCurrentPage } = storeToRefs(useMailStore());
const { nextPage, previousPage } = useMailStore();

const totalPages = computed(() => {
  return props.totalRecords === 0
    ? 0
    : Math.ceil(props.totalRecords / props.pageSize);
});

function onPreviousPage() {
  if (getCurrentPage.value > 1) {
    previousPage();
    emit("pagination:change", { currentPage: getCurrentPage.value });
  }
}

function onNextPage() {
  if (getCurrentPage.value < totalPages.value) {
    nextPage();
    emit("pagination:change", { currentPage: getCurrentPage.value });
  }
}
</script>
