import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Checkout from "components/Checkout";

import "css/normalize.css";
import "css/global.css";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/success.html">
          <div />
        </Route>
        <Route path="/canceled.html">
          <div></div>
        </Route>
        <Route path="/">
          <Checkout />
        </Route>
      </Switch>
    </Router>
  );
}

ReactDOM.render(<App />, document.getElementById("root"));
