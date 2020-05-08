import React, { useEffect } from "react";
import useAuth from "hooks/use-auth";

function App() {
  const { loading, isLoggedIn, login, logout, user, getDIDToken } = useAuth();

  useEffect(() => {
    console.log({ loading, isLoggedIn, user });
  });

  return (
    <div>
      <button onClick={() => login("daongocthanh98hy@gmail.com")}>Login</button>
      <button onClick={getDIDToken}>Get DIDToken</button>
      <button onClick={logout}>Logout</button>
    </div>
  );
}

export default App;
