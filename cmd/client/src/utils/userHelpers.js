import { getHouseholdUsers } from "../api/users";

export async function getUsers(hhid) {
    return await getHouseholdUsers(hhid)
}