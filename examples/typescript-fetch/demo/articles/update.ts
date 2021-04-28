import { ArticlesApi } from "../../microcms";
import { ArticlesPatchRequestCategoryEnum } from "../../microcms/models";
import { config } from "../config";

const api = new ArticlesApi(config);

api
  .updateArticles({
    contentId: "demo",
    articlesPatchRequest: {
      title: "Team Fortress 2",
      content: "https://www.teamfortress.com/",
      author: "demo",
      category: [ArticlesPatchRequestCategoryEnum.Game],
      tag: [],
      deadline: new Date(),
    },
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
