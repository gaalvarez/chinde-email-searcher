export interface SearchResults {
  data: Data;
}

export interface Data {
  total: number;
  results: EmailResult[];
}

export interface EmailResult {
  bcc: string[];
  body: string;
  cc: string[];
  date: string;
  from: string;
  messageID: string;
  subject: string;
  to: string[];
}

export interface EmailInfo {
  messageID: string;
  from: string;
  to: string[];
  date: string;
  subject: string;
  body: string;
}

export interface SearchConfig {
  search_type:
    | "match"
    | "matchall"
    | "matchphrase"
    | "term"
    | "querystring"
    | "prefix"
    | "wildcard"
    | "fuzzy"
    | "daterange";
  query: Query;
  from: number;
  max_results: number;
  _source: string[];
  highlight: Highlight;
}

export interface Highlight {
  fields: {
    [key: string]: {};
  };
}

export interface Query {
  term: string;
}

export interface MailSearchResultState {
  message: "empty" | "noresult";
}
