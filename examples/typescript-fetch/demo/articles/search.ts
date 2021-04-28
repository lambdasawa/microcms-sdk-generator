import {
  ArticlesApi,
  SearchArticlesFieldsEnum,
  SearchArticlesOrdersEnum,
} from "../../microcms";
import { config } from "../config";

const api = new ArticlesApi(config);

api
  .searchArticles({
    limit: 1,
    depth: 1,
    fields: [SearchArticlesFieldsEnum.Title, SearchArticlesFieldsEnum.Content],
    q: "microcms",
    orders: [SearchArticlesOrdersEnum.publishedAtDescending],
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
