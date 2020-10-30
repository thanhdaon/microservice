const SET_WINDOW_WIDTH = "set_window_width";
const TOGGLE_COLLAPSED_NAV = "toggle_collapsed_nav";
const SWITCH_LANGUAGE = "switch_language";

const SettingActions = {
  updateWindowWidth: (width) => ({ type: SET_WINDOW_WIDTH, payload: width }),
  toggleCollapsedSideNav: (navCollapsed) => ({
    type: TOGGLE_COLLAPSED_NAV,
    payload: navCollapsed,
  }),
  switchLanguage: (locale) => ({
    type: SWITCH_LANGUAGE,
    payload: locale,
  }),
};

const initialState = {
  navCollapsed: true,
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
    case SET_WINDOW_WIDTH:
      return { ...state, width: action.payload };
    case TOGGLE_COLLAPSED_NAV:
      return { ...state, navCollapsed: action.payload };
    case SWITCH_LANGUAGE:
      return { ...state, locale: action.payload };
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
