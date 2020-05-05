import { useState, useEffect } from "react";
import { Magic } from "magic-sdk";

const magic = new Magic("pk_test_C1806F6815229492");

function useAuth() {
  const [loading, setLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState({});

  async function refreshLoggedIn() {
    setLoading(true);
    const loggedIn = await magic.user.isLoggedIn();
    setIsLoggedIn(loggedIn);
    if (loggedIn) {
      setUser(await magic.user.getMetadata());
    } else {
      setUser({});
    }
    setLoading(false);
  }

  async function login(email) {
    if (email) {
      await magic.auth.loginWithMagicLink({ email });
      refreshLoggedIn();
    }
  }

  async function logout() {
    await magic.user.logout();
    refreshLoggedIn();
  }

  async function getDIDToken() {
    const didToken = await magic.user.getIdToken();
    console.log({ didToken });
  }

  useEffect(() => {
    refreshLoggedIn();
  }, []);

  return { loading, isLoggedIn, login, user, getDIDToken, logout };
}

export default useAuth;
