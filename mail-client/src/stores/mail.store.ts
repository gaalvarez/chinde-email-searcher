//define store for mail using pinia
import { mailApi } from "@/api/mail.api";
import type {
  EmailResult,
  MailSearchResultState,
  SearchResults,
} from "@/api/mail.model";
import type { AxiosResponse } from "axios";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useMailStore = defineStore("mail", () => {
  //state
  const previousSearches = ref<string[]>([]);
  const emails = ref<EmailResult[]>([]);
  const selectedEmail = ref<EmailResult | null>(null);
  const messageEmailsSearch = ref<MailSearchResultState>({ message: "empty" });
  const totalRecords = ref<number>(0);
  const currentPage = ref<number>(1);

  //actions
  async function searchEmails(
    searchTerm: string,
    pagination: { currentPage: number; pageSize: number }
  ) {
    const newEmails: AxiosResponse<SearchResults> = await mailApi.searchEmails(
      searchTerm,
      pagination.currentPage,
      pagination.pageSize
    );
    if (newEmails.data.data.results.length < 1) {
      messageEmailsSearch.value = { message: "noresult" };
    }
    totalRecords.value = newEmails.data.data.total;
    selectedEmail.value = null;
    emails.value = newEmails.data.data.results;
  }

  function addToPreviousSearches(newSearch: string) {
    previousSearches.value.push(newSearch);
  }

  function selectEmail(messageID: string) {
    const email = emails.value.find((email) => email.messageID === messageID);
    if (email) {
      selectedEmail.value = email;
    }
  }

  function cleanSelectEmail() {
    selectedEmail.value = null;
  }

  function resetEmailList() {
    messageEmailsSearch.value = { message: "empty" };
    emails.value = [];
    totalRecords.value = 0;
    currentPage.value = 1;
    cleanSelectEmail();
  }

  function nextPage() {
    currentPage.value++;
  }

  function previousPage() {
    currentPage.value--;
  }

  //getters
  const getEmails = computed(() => emails.value);
  const getPreviousSearches = computed(() => previousSearches.value);
  const getSelectedEmail = computed(() => selectedEmail.value);
  const getMessageEmailSearch = computed(() => messageEmailsSearch.value);
  const getTotalRecords = computed(() => totalRecords.value);
  const getCurrentPage = computed(() => currentPage.value);

  return {
    getEmails,
    getPreviousSearches,
    getSelectedEmail,
    searchEmails,
    addToPreviousSearches,
    selectEmail,
    resetEmailList,
    getMessageEmailSearch,
    cleanSelectEmail,
    getTotalRecords,
    previousPage,
    nextPage,
    getCurrentPage,
  };
});
