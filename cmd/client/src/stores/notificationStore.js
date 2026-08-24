import { create } from "zustand"

export const useNotificationStore = create((set) => ({
    notificationOpen: false,
    notificationMessage: "",

    showNotification: (message) => set({
        notificationOpen: true,
        notificationMessage: message,
    }),

    closeNotification: () => set({
        notificationOpen: false,
        notificationMessage: "",
    }),
}));