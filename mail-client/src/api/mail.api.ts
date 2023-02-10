import { HttpClient } from "./http-client";
import type { SearchConfig, SearchResults } from "./mail.model";

class MailApi extends HttpClient {
  private searchConfig: SearchConfig = {
    search_type: "querystring",
    query: {
      term: "",
    },
    from: 0,
    max_results: 10,
    _source: [],
    highlight: {
      fields: {
        Body: {},
      },
    },
  };

  public constructor() {
    super(
      import.meta.env.VITE_APP_BACKEND_API_URL || `http://${window.location.origin}:8080/mail`
    );
  }

  public searchEmails = (
    searchTerm: string,
    page: number,
    recordsPerPage: number
  ) => {
    return this.http.post<SearchResults>(`/search`, {
      ...this.searchConfig,
      query: {
        term: `+Body:${searchTerm}`,
      },
      from: (page - 1) * recordsPerPage,
      max_results: recordsPerPage,
    });
  };
}

export const mailApi = new MailApi();
