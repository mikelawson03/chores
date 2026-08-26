import { API_HOST } from "../config/dev";
import { ApiError } from "./apiError";
import { getHeaders } from "./headers";

export async function getTasks() {
    const response = await fetch(`${API_HOST}/assignments`,{
        method: "GET",
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.json();
        throw new ApiError(response.status, error.error);
    }

    return await response.json();
}

export async function getTasksForUser(id) {
    const response = await fetch(`${API_HOST}/assignments?user_id=${id}`, {
        method: "GET",
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.json();
        throw new ApiError(response.status, error.error);
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
        const error = await response.json();
        throw new ApiError(response.status, error.error);
    }

    return await response.json();
}

export async function toggleCompletion(id) {
    const response = await fetch(`${API_HOST}/assignments/${id}/complete`, {
        method: "POST",
        headers: getHeaders(),
    })

    if (!response.ok) {
        const error = await response.json();
        throw new ApiError(response.status, error.error);
    }

    return await response.json();
}

export async function rescheduleAssignment(props) {
    const response = await fetch(`${API_HOST}/assignments/${props.id}/reschedule`, {
        method: "POST",
        headers: getHeaders(),
        body: JSON.stringify({"scheduledFor": props.scheduledFor})
    })

    if (!response.ok) {
        const error = await response.json();
        throw new ApiError(response.status, error.error);
    }

    return await response.json();
}