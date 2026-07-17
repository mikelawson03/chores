import { API_HOST } from "../config/dev"
import { getHeaders } from "./headers"

export async function getChoreTemplates() {
    const response = await fetch(`${API_HOST}/chore-templates`,{
        method: "GET",
        headers: getHeaders(),
    })

    const data = await response.json();

    console.log(data);

    return await data;

};

export async function createChoreTemplate(template) {
  const response = await fetch(`${API_HOST}/chore-templates`,{
    method: "POST",
    body: JSON.stringify(template),
    headers: getHeaders(),
  })

  if (!response.ok) {
    throw new Error("Failed to create chore template");
  }

  return await response.json();

}

export async function updateChoreTemplate(template) {
  const response = await fetch(`${API_HOST}/chore-templates/${template.id}`, {
    method: "PUT",
    body: JSON.stringify(template),
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.text();
    throw new Error(`Server returned ${response.status}: ${error}`);
  }

  return await response.json();
}

export async function deleteChoreTemplate(id) {
  const response = await fetch(`${API_HOST}/chore-templates/${id}`, {
    method: "DELETE",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.text();
    throw newError(`Server returned ${response.status}: ${error}`);
  }
}