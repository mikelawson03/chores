import { API_HOST } from "../config/dev";
import { getHeaders } from "./headers";

export async function login(username, password) {
    const response = await fetch(`${API_HOST}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username,
            password,
        })
    });

    if (!response.ok) {
        throw new Error("Invalid credentials");
    }

    return response.json();
}

export async function getMe() {
    const response = await fetch(`${API_HOST}/me`, {
        method: "GET",
        headers: getHeaders(),
    })

    if (!response.ok) {
        throw new Error("User not found");
    }

    return response.json();
}