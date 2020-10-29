import {
  LAYOUT_TYPE_FULL,
  NAV_STYLE_FIXED,
  THEME_COLOR_SELECTION_PRESET,
  THEME_TYPE_SEMI_DARK,
} from "constants/theme-settings";

const WINDOW_WIDTH = "WINDOW_WIDTH";

const SettingActions = {
  updateWindowWidth: (width) => ({ type: WINDOW_WIDTH, payload: width }),
};

const initialState = {
  navCollapsed: true,
  navStyle: NAV_STYLE_FIXED,
  layoutType: LAYOUT_TYPE_FULL,
  themeType: THEME_TYPE_SEMI_DARK,
  colorSelection: THEME_COLOR_SELECTION_PRESET,

  pathname: "",
  width: typeof window === "undefined" ? 0 : window.innerWidth,
  locale: {
    languageId: "english",
    locale: "en",
    name: "English",
    icon: "us",
  },
};

function reducer(state = initialState, action) {
  switch (action.type) {
    case SettingActions.WINDOW_WIDTH:
      return { ...state, width: action.payload };
    default:
      return state;
  }
}

function getSettingsModule() {
  return {
    id: "settings",
    reducerMap: { settings: reducer },
  };
}

export { SettingActions, getSettingsModule };
