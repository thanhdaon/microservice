import "styles/tailwind.css";

import { Provider } from "react-redux";

import store from "app-redux";

function App({ Component, pageProps }) {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default App;
