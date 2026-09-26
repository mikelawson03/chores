import { API_HOST } from "../config/dev"
import { getHeaders } from "./headers"
import { ApiError } from "./apiError";

export async function getChoreTemplates() {
    const response = await fetch(`${API_HOST}/chore-templates`,{
        method: "GET",
        headers: getHeaders(),
        credentials: "include",
    })

    return await response.json();

};

export async function createChoreTemplate(template) {
  console.log(template);
  const response = await fetch(`${API_HOST}/chore-templates`,{
    method: "POST",
    body: JSON.stringify(template),
    headers: getHeaders(),
    credentials: "include",
  })

  if (!response.ok) {
    const error = await response.json();
    throw new ApiError(response.status, error.error)
  }

  return await response.json();

}

export async function updateChoreTemplate(template) {
  const response = await fetch(`${API_HOST}/chore-templates/${template.id}`, {
    method: "PUT",
    body: JSON.stringify(template),
    headers: getHeaders(),
    credentials: "include",
  })

  if (!response.ok) {
    const error = await response.json();
    throw new ApiError(response.status, error.error)
  }

  return await response.json();
}

export async function deleteChoreTemplate(id) {
  const response = await fetch(`${API_HOST}/chore-templates/${id}`, {
    method: "DELETE",
    headers: getHeaders(),
    credentials: "include",
  })

  if (!response.ok) {
    const error = await response.text();
    throw new Error(`Server returned ${response.status}: ${error}`);
  }
}