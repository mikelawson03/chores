import { API_HOST } from "../config/dev";
import { getHeaders } from "./headers";

export async function getTasks() {
    const response = await fetch(`${API_HOST}/assignments`,{
        method: "GET",
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.text();
        throw new Error(`Server returned ${response.status}: ${error}`);
    }

    return await response.json();
}

export async function getTasksForUser(id) {
    const response = await fetch(`${API_HOST}/assignments?user_id=${id}`, {
        method: "GET",
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.text();
        throw new Error(`Server returned ${response.status}: ${error}`);
    }

    return await response.json();
}

export async function updateTask(task) {
    const response = await fetch(`${API_HOST}/assignments/${task.id}`,{
        method: "PUT",
        body: JSON.stringify(task),
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.text();
        throw new Error(`Server returned ${response.status}: ${error}`)
    }

    return await response.json();
}

