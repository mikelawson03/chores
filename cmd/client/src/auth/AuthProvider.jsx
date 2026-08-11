import { createContext, useEffect, useState } from "react";
import { getMe, login } from "../api/users";

export const AuthContext = createContext(null);

export default function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const restoreSession = async () => {
            const token = localStorage.getItem("token");

            if (!token) {
                setLoading(false);
                return;
            }

            try {
                const user = await getMe();
                setUser(user);
            } catch {
                localStorage.removeItem("token");
            } finally {
                setLoading(false);
            }
        };

        restoreSession();
    }, []);

    const logout = () => {
        localStorage.removeItem("token");
        setUser(null);
    };

    const loginUser = async (username, password) => {
        const result = await login(username, password);

        localStorage.setItem("token", result.token)
        
        const user = await getMe();
        setUser(user);        
    };

    return(
        <AuthContext.Provider value= {{ user, loading, loginUser, logout }}>
            {children}
        </AuthContext.Provider>
    )
}