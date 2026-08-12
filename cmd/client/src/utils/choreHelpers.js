import { createChoreTemplate, deleteChoreTemplate, getChoreTemplates, updateChoreTemplate } from "../api/choreTemplates";

export async function getChores() {
    return await getChoreTemplates();
}

export async function createChore(chore) {
    return await createChoreTemplate(chore);
}

export async function updateChore(chore) {
    return await updateChoreTemplate(chore);
}

export async function deleteChore(id) {
    return await deleteChoreTemplate(id);
}