import axios from "axios";
import type { AxiosInstance, AxiosRequestConfig } from "axios";

export interface ConfigRepoAction {
  token: string;
  repoPath: string;
  workflowIdPath: string;
}

export abstract class HttpClient {
  protected readonly http: AxiosInstance;
  // protected readonly accessToken: string

  public constructor(baseURL: string) {
    // this.accessToken = accessToken
    this.http = axios.create({
      baseURL,
    });
    // this.addInterceptor()
  }

  //   this.http.interceptors.request.use(
  //     (config: AxiosRequestConfig) => {
  //       return {
  //         ...config,
  //         headers: {
  //           Authorization: `Bearer ${this.accessToken}`
  //         }
  //       }
  //     },
  //     (error) => {
  //       return Promise.reject(error)
  //     }
  //   )
  // }
}
