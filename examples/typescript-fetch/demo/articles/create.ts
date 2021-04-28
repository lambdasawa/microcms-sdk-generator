import { ArticlesApi } from "../../microcms";
import {
  ArticlesCreateRequestCategoryEnum,
  ArticlesCreateRequestTagEnum,
} from "../../microcms/models";
import { config } from "../config";

const api = new ArticlesApi(config);

api
  .createArticles({
    articlesCreateRequest: {
      title: "microcms-sdk-generator",
      content: "microcms-sdk-generator is awesome.",
      author: "demo",
      category: [ArticlesCreateRequestCategoryEnum.Programming],
      tag: [
        ArticlesCreateRequestTagEnum.Golang,
        ArticlesCreateRequestTagEnum.Typescript,
      ],
      deadline: new Date(),
    },
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
