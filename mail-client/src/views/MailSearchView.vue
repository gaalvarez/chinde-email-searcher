<template>
  <section class="flex flex-1 overflow-y-scroll">
    <div class="p-4 flex flex-row justify-center flex-1 overflow-x-auto">
      <div
        class="flex flex-row justify-center items-start h-full w-full gap-2 wrapper-content"
      >
        <div
          :class="{
            'w-0': isMobile && getSelectedEmail !== null,
            'w-full': getSelectedEmail === null,
          }"
          class="lg:w-1/3 h-full overflow-y-scroll"
        >
          <!-- Left column content here -->
          <div class="bg-white rounded-lg shadow-md pb-4 h-full">
            <!-- Left column content here -->
            <div class="p-4">
              <InputSearch
                @search:complete="onSearchComplete"
                @search:clear="onSearchClear"
              />
            </div>
            <EmailsPagination
              :page-size="pageSize"
              :total-records="getTotalRecords"
              @pagination:change="onPaginationChange"
            ></EmailsPagination>
            <NoEmailsFounded
              :messageEmailSearch="getMessageEmailSearch"
              v-if="getEmails.length === 0"
            />
            <div v-else>
              <EmailFound
                v-for="email in getEmails"
                :key="email.messageID"
                :emailInfo="email"
                @selectEmail="onSelectEmail"
              ></EmailFound>
            </div>
          </div>
        </div>
        <div
          :class="{
            'w-full': isMobile && getSelectedEmail !== null,
            'w-0': getSelectedEmail === null,
          }"
          class="lg:w-2/3 h-full overflow-y-scroll"
        >
          <!-- Right column content here -->
          <div class="bg-white pb-4 rounded-lg shadow-md h-full">
            <!-- Left column content here -->

            <EmailDetail
              v-if="getSelectedEmail"
              :emailInfo="getSelectedEmail"
              :search-term="searchTerm"
              @emaildetail:close="onCloseEmailDetail"
            />
            <NoSelectedEmail v-else />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import InputSearch from "@/components/InputSearch.vue";
import EmailFound from "@/components/EmailFound.vue";
import NoEmailsFounded from "@/components/NoEmailsFounded.vue";
import EmailDetail from "@/components/EmailDetail.vue";
import NoSelectedEmail from "@/components/NoSelectedEmail.vue";
import { useMailStore } from "@/stores/mail.store";
import { storeToRefs } from "pinia";
import { ref } from "vue";
import EmailsPagination from "@/components/EmailsPagination.vue";

const { searchEmails, resetEmailList, selectEmail, cleanSelectEmail } =
  useMailStore();

const { getEmails, getSelectedEmail, getMessageEmailSearch, getTotalRecords } =
  storeToRefs(useMailStore());

const pageSize = 10;
const searchTerm = ref("");
const isMobile = ref(false);

const onSearchComplete = (value: string): void => {
  console.log("onSearchComplete", value);
  searchTerm.value = value;
  if (value) {
    searchEmails(value, { currentPage: 1, pageSize });
  } else {
    resetEmailList();
  }
};

const onSearchClear = (): void => {
  searchTerm.value = "";
  resetEmailList();
};

const onSelectEmail = (messageID: string): void => {
  console.log("onSelectEmail", messageID);
  isMobile.value = window.matchMedia("(max-width: 1024px)").matches;
  selectEmail(messageID);
};

const onCloseEmailDetail = () => {
  isMobile.value = window.matchMedia("(max-width: 1024px)").matches;
  cleanSelectEmail();
};

const onPaginationChange = (value: { currentPage: number }) => {
  if (searchTerm.value) {
    searchEmails(searchTerm.value, {
      currentPage: value.currentPage,
      pageSize,
    });
  } else {
    resetEmailList();
  }
};
</script>

<style scoped>
.resize-x {
  resize: horizontal;
  overflow: auto;
}

.wrapper-content {
  max-width: 93.75rem;
}

.main-section {
  flex: 1;
  display: flex;
}
</style>
