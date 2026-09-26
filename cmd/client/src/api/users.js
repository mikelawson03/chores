import { API_HOST } from "../config/dev";
import { getHeaders } from "./headers";

export async function login(username, password) {
    const response = await fetch(`${API_HOST}/login`, {
        method: "POST",
        credentials: "include",
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

}

export async function logout() {
    const response = await fetch(`${API_HOST}/logout`, {
        method: "POST",
        credentials: "include",
        headers: getHeaders(),
    })

    if (!response.ok) {
        throw new Error("Unexpected error");
    }
}

export async function getMe() {
    const response = await fetch(`${API_HOST}/me`, {
        method: "GET",
        credentials: "include",
        headers: getHeaders(),
    })

    if (!response.ok) {
        throw new Error("User not found");
    }

    return response.json();
}

export async function getHouseholdUsers(hhid) {
    const response = await fetch(`${API_HOST}/households/${hhid}/users`, {
        method: "GET",
        headers: getHeaders(),
        credentials: "include",
    })
    if (!response.ok) {
        throw new Error("Users not found");
    }

    return response.json();
}