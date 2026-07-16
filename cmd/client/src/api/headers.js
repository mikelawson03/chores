import { DEV_USER_ID } from "../config/dev";

export function getHeaders() {
    return {
        "Content-Type": "application/json",
        "X-User-ID": DEV_USER_ID,
    };
}