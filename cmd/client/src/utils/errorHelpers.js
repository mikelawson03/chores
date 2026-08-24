export function parseApiError(error) {
    const [type, ...messageParts] = error.message.split(":");

    return {
        type: type.trim(),
        message: messageParts.join(":").trim(),
    };
}