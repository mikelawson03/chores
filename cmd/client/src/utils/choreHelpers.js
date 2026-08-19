import { createChoreTemplate, deleteChoreTemplate, getChoreTemplates, updateChoreTemplate } from "../api/choreTemplates";

function prepareChoreForSave(chore) {
    return {
        ...chore,
        assignee: chore.assignee === "unassigned"
            ? ""
            : chore.assignee
    };
}

export async function getChores() {
    return await getChoreTemplates();
}

export async function createChore(chore) {
    chore.assignee === "unassigned" ? null : chore.assignee
    return await createChoreTemplate(prepareChoreForSave(chore));
}

export async function updateChore(chore) {
    chore.assignee === "unassigned" ? null : chore.assignee
    return await updateChoreTemplate(prepareChoreForSave(chore));
}

export async function deleteChore(id) {
    return await deleteChoreTemplate(id);
}