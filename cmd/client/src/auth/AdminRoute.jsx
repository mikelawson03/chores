import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./useAuth";

export default function AdminRoute() {
    const { user } = useAuth();

    if (user?.role !== "admin") {
        return <Navigate to="/" replace />
    }

    return <Outlet />
}