<template>
  <div
    class="bg-white p-4 border-t border-gray-300 cursor-pointer"
    @click="() => onClick()"
  >
    <div class="flex">
      <div class="w-1/2">
        <div class="text-sm text-gray-700 font-medium">
          From:
          <span class="text-sm text-black font-medium">{{
            emailInfo.from
          }}</span>
        </div>

        <div class="text-sm text-gray-700 font-medium">
          To:
          <span class="text-sm text-black font-medium">{{ toShort }}</span>
        </div>
      </div>
      <div class="w-1/2 text-right">
        <div class="text-xs text-gray-500 font-xs">{{ emailInfo.date }}</div>
      </div>
    </div>
    <div class="mt-2 text-l text-black font-medium mb-2">
      {{ emailInfo.subject }}
    </div>
    <div class="text-sm text-gray-600 font-medium">
      {{ bodyShort }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { EmailInfo } from "@/api/mail.model";

const props = defineProps<{
  emailInfo: EmailInfo;
}>();

const emit = defineEmits<{
  (event: "selectEmail", messageID: string): void;
}>();

const onClick = (): void => {
  emit("selectEmail", props.emailInfo.messageID);
};

//get first 100 characters of body
const bodyShort = `${props.emailInfo.body.substring(0, 100)}...`;

const toShort = `${props.emailInfo.to.slice(0, 2).join(", ")}${
  props.emailInfo.to.length > 2 ? "..." : ""
}`;
</script>
