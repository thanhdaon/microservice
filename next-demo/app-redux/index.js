import { createStore } from "redux-dynamic-modules";
import { getThunkExtension } from "redux-dynamic-modules-thunk";

import { getSettingsModule } from "app-redux/settings";

const store = createStore(
  {
    extensions: [getThunkExtension()],
  },
  getSettingsModule()
);

export default store;
