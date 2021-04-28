import { ArticlesApi } from "../../microcms";
import { config } from "../config";

const api = new ArticlesApi(config);

api
  .fetchArticles({
    contentId: "demo",
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
