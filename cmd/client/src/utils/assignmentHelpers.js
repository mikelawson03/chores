import { getTasks, getTasksForUser, updateTask } from "../api/tasks";
import { toggleCompletion } from "../api/tasks";

export async function getAssignments(user) {
    if (user.role === "admin"){
        return await getTasks();
    }

    return await getTasksForUser(user.id);
}

export async function updateAssignment(assignment) {
    return await updateTask(assignment);
}

export async function toggleTaskCompletion(id) {
    return await toggleCompletion(id);
}