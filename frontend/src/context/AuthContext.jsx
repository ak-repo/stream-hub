import { useState } from "react";
import { AuthContext } from "./context";
import { loginService, registerService } from "../api/services/authService";

function AuthProvider({ children }) {
  const [user, setUser] = useState(null);

  // useEffect(() => {
  //   (async () => {
  //     try {
  //       await getMe();
  //       setAuthenticated(true);
  //     } catch (error) {
  //       console.log("errr :", error);
  //     }
  //   })();
  // }, []);

  const isAuthenticated = !!user;

  const login = async (email, password) => {
    const data = await loginService(email, password);
    setUser(data.user);
  };

  const register = async (form) => {
    const data = await registerService(form);
    setUser(data.user);
  };

  const logout = () => {
    // logoutService();
    setUser(null);
  };
  return (
    <AuthContext.Provider
      value={{
        user,
        setUser,
        login,
        logout,
        isAuthenticated,
        register,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export default AuthProvider;
