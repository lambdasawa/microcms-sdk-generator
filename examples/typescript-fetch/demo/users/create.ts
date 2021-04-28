import { UsersApi, UsersSnsFieldIdEnum } from "../../microcms";
import { config } from "../config";

const api = new UsersApi(config);

api
  .createUsers({
    usersCreateRequest: {
      name: `${Date.now()}`,
      email: `${Date.now()}@example.com`,
      bio: `${Date.now()}`,
      birthday: new Date(),
      isSnsPublic: true,
      twitter: {
        fieldId: UsersSnsFieldIdEnum.Sns,
        userId: "lambdasawa",
        url: "https://twitter.com/intent/follow?screen_name=lambdasawa",
      },
    },
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
