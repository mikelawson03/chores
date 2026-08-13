import { create } from "zustand";

export const useTaskStore = create((set) => ({
    selectedTask: null,
    taskDetailsOpen: false,

    openTaskDetails: (task) => set({
        selectedTask: task,
        taskDetailsOpen: true,
    }),

    closeTaskDetails: () => set({
        selectedTask: null,
        taskDetailsOpen: false,
    }),
}));