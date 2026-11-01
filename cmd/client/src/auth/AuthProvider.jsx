import { createContext, useEffect, useState } from "react";
import { getMe, login, logout } from "../api/users";
import { parseApiError } from "../utils/errorHelpers";

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
        
        try {
        const user = await getMe();
        setUser(user);
        } catch (err) {
            const { type } = parseApiError(err);
            if (
                type === "no_household_membership" ||
                type === "inactive_household_user"
            ) {
                throw new Error("You are not an active member of a household.");
            }
        }
    };

    const refreshUser = async () => {
        try {
            const refreshedUser = await getMe();
            setUser(refreshedUser);
            return refreshedUser;
        } catch (err) {
            if (err.status === 401) {
                setUser(null);
                return null;
            }
            
            throw err;
        }
        
    }

    return(
        <AuthContext.Provider value= {{ user, loading, loginUser, logoutUser, refreshUser }}>
            {children}
        </AuthContext.Provider>
    )
}