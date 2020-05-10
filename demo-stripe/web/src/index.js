import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Products from "components/Products";
import Checkout from "components/Checkout";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/checkout/:productId">
          <Checkout />
        </Route>
        <Route path="/products">
          <Products />
        </Route>
      </Switch>
    </Router>
  );
}

ReactDOM.render(<App />, document.getElementById("root"));
