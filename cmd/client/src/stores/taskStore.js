import { create } from "zustand";

export const useTaskStore = create((set) => ({
    selectedTask: null,
    taskDetailsOpen: false,
    taskErrors: {},

    openTaskDetails: (task) => set({
        selectedTask: task,
        taskDetailsOpen: true,
        taskErrors: {},
    }),

    closeTaskDetails: () => set({
        selectedTask: null,
        taskDetailsOpen: false,
        taskErorrs: {},
    }),

    setTaskError: (field, error) => set((state) => ({
        taskErrors: {
            ...state.taskErrors,
            [field]: error,
        },
    })),

    clearTaskErrors: () => set({
        taskErrors: {},
    }),
}));