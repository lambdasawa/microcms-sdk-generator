import fetch from "node-fetch";
import { Configuration } from "../microcms";

export const config = new Configuration({
  apiKey: (name) => {
    switch (name.toLowerCase()) {
      case "x-api-key":
        return process.env.X_API_KEY || "";
      case "x-write-api-key":
        return process.env.X_WRITE_API_KEY || "";
      case "x-global-draft-key":
        return process.env.X_GLOBAL_DRAFT_KEY || "";
      default:
        return "";
    }
  },
  fetchApi: (input, init) => fetch(input as any, init as any) as any, // TODO: fix type error
});
