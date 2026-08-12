import { getAllUsers } from "../api/users";

export async function getUsers() {
    return await getAllUsers()
}