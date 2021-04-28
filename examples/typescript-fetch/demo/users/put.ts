import { UsersApi } from "../../microcms";
import { config } from "../config";

const api = new UsersApi(config);

api
  .putUsers({
    contentId: "demo",
    usersCreateRequest: {
      name: `${Date.now()}`,
      email: `${Date.now()}@example.com`,
      bio: `${Date.now()}`,
      birthday: new Date(),
      isSnsPublic: true,
    },
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
