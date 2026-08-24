import { create } from "zustand";

export const useTaskStore = create((set) => ({
    selectedTask: null,
    taskDetailsOpen: false,
    taskDraft: null,
    taskErrors: {},

    openTaskDetails: (task) => set({
        selectedTask: task,
        taskDraft: {...task},
        taskDetailsOpen: true,
        taskErrors: {},
    }),

    closeTaskDetails: () => set({
        selectedTask: null,
        taskDraft: null,
        taskDetailsOpen: false,
        taskErorrs: {},
    }),

    updateTaskDraft: (field, value) => set((state) => ({
        taskDraft: {
            ...state.taskDraft,
            [field]: value,
        },
    })),

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