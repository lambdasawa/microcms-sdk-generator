import { SettingApi } from "../../microcms";
import { config } from "../config";

const api = new SettingApi(config);

api
  .updateSetting({
    settingPatchRequest: {
      articlesPerPage: 10 + Math.ceil(Math.random() * 10),
    },
  })
  .then((value) => console.log(JSON.stringify(value, null, 2)))
  .catch(async (error) => {
    console.error(await error.text());
    process.exit(1);
  });
