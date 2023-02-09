<template>
  <!-- <div>
    <input type="text" placeholder="Search" v-model="searchInput" />
  </div> -->
  <div class="relative rounded-md shadow-sm">
    <div
      class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
    >
      <svg
        class="h-5 w-5 text-gray-400"
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path
          fill-rule="evenodd"
          d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
          clip-rule="evenodd"
        ></path>
      </svg>
    </div>
    <input
      ref="el"
      v-model="searchInput"
      class="form-input py-2 pl-10 w-full leading-5 rounded-md transition duration-150 ease-in-out bg-white border border-gray-300 focus:outline-none focus:border-indigo-600 focus:shadow-outline-indigo active:bg-gray-50 active:text-indigo-800"
      placeholder="Your search term"
    />
    <div
      :class="searchInput ? 'block' : 'hidden'"
      class="absolute inset-y-0 right-0 pr-1 flex items-center cursos-pointer"
      @click="onClearInput"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="20px"
        width="20px"
        viewBox="0 0 25 25"
      >
        <path
          d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
        />
        <path d="M0 0h24v24H0z" fill="none" />
      </svg>
    </div>
  </div>
</template>

<script setup lang="ts">
import { debounce } from "lodash";
import { ref, watch, onMounted } from "vue";

const el = ref<HTMLInputElement | null>(null);

onMounted(() => {
  el.value?.focus();
});

const searchInput = ref("");

const emit = defineEmits<{
  (e: "search:complete", value: string): void;
  (e: "search:clear"): void;
}>();

function onClearInput(): void {
  searchInput.value = "";
  el.value?.focus();
  emit("search:clear");
}

watch(searchInput, (value: string): void => {
  debouncedEmitValue(value);
});

const debouncedEmitValue = debounce((value: string): void => {
  emit("search:complete", value);
}, 600);
</script>
