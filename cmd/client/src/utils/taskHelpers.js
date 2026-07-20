import dayjs from "dayjs";

export function getWeeklyTasks(tasks) {
    return tasks.filter(task => task.cadence === "weekly");
}

export function getMonthlyTasks(tasks) {
    return tasks.filter(task => task.cadence === "monthly");
}

export function getCompletedTasks(tasks) {
    return tasks.filter(task => task.completed);
}

export function getScheduledTasksForDay(date, tasks) {
    return tasks.filter(task => dayjs(task.scheduledFor).isSame(date, "day"));
}

export function getTasksDueInMonth(date, tasks) {
    return tasks.filter(task => dayjs(task.dueDate).isSame(date, "month"));
}

export function removeCompletedTasks(tasks) {
    return tasks.filter(task => !task.completed);
}

export function removeCanceledTasks(tasks) {
    return tasks.filter(task => !task.canceled);
}

export function getUnscheduledTasks(tasks) {
    return tasks.filter(task => !task.scheduledFor);
}

export function getActiveTasks(tasks) {
    return removeCanceledTasks(removeCompletedTasks(tasks));
}

export function tasksEqual(task1, task2) {
    return JSON.stringify(task1) === JSON.stringify(task2);
}
