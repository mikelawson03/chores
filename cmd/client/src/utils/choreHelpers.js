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
    return await createChoreTemplate(prepareChoreForSave(chore));
}

export async function updateChore(chore) {
    return await updateChoreTemplate(prepareChoreForSave(chore));
}

export async function deleteChore(id) {
    return await deleteChoreTemplate(id);
}

export const choreTemplateValidationErrors = {
    "name required": {
        field: "name",
        message: "Chore name is required."
    },
    "cadence required": {
        field: "cadence",
        message: "Cadence is required."
    },
    "duration must be greater than zero": {
        field: "duration",
        message: "Duration must be greater than zero."
    },
    "assigned user not found": {
        field: "assignee",
        message: "User not found."
    },
}

export function getChoreFieldError(field, value) {
    switch (field) {
      case "name":
        if (!value.trim()) {
          return "Chore name is required.";
        }
        break;
      
        case "cadence":
        if (!value) {
          return "Frequency is required.";
        }
        break;
      
        case "duration":
        if (value <= 0) {
          return "Duration must be greater than zero.";
        }
        break;

      default:
        return null;
    }
  }