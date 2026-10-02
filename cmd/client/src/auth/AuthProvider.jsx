import { createContext, useEffect, useState } from "react";
import { getMe, login, logout } from "../api/users";

export const AuthContext = createContext(null);

export default function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const restoreSession = async () => {
            try {
                const user = await getMe();
                setUser(user);
            } catch {
                setUser(null);
            } finally {
                setLoading(false);
            }
        };

        restoreSession();
    }, []);

    const logoutUser = async () => {
        await logout();
        setUser(null);
    };

    const loginUser = async (username, password) => {
        await login(username, password);
        
        const user = await getMe();
        setUser(user);        
    };

    const refreshUser = async () => {
        const refreshedUser = await getMe();
        setUser(refreshedUser);

        return refreshedUser;
    }

    return(
        <AuthContext.Provider value= {{ user, loading, loginUser, logoutUser, refreshUser }}>
            {children}
        </AuthContext.Provider>
    )
}