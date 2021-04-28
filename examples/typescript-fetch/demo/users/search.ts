import {
  SearchUsersFieldsEnum,
  SearchUsersOrdersEnum,
  UsersApi,
} from "../../microcms";
import { config } from "../config";

const api = new UsersApi(config);

api
  .searchUsers({
    limit: 1,
    depth: 1,
    fields: [
      SearchUsersFieldsEnum.Name,
      SearchUsersFieldsEnum.Bio,
      SearchUsersFieldsEnum.Icon,
    ],
    q: "lambdasawa",
    orders: [SearchUsersOrdersEnum.publishedAtDescending],
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
