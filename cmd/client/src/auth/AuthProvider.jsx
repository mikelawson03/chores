import { createContext, useState } from "react";
import { getMe, login } from "../api/users";

export const AuthContext = createContext(null);

export default function AuthProvider({ children }) {
    const [user, setUser] = useState(null);

    const loginUser = async (username, password) => {
        const result = await login(username, password);

        localStorage.setItem("token", result.token)
        
        const user = await getMe();
        setUser(user);        
    }

    return(
        <AuthContext.Provider value= {{ user, loginUser }}>
            {children}
        </AuthContext.Provider>
    )
}