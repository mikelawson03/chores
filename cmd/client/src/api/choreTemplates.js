import { DEV_USER_ID, API_HOST } from "../config/dev"
import { getHeaders } from "./headers"

export async function getChoreTemplates() {
    const response = await fetch(`${API_HOST}/chore-templates`,{
        method: "GET",
        headers: getHeaders(),
    })

    const data = await response.json();

    return data;
};