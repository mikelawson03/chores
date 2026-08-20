export function parseApiError(error) {
    const [type, ...messageParts] = error.message.split(":");

    console.log(type)
    console.log(messageParts.join(":").trim())
    return {
        type: type.trim(),
        message: messageParts.join(":").trim(),
    };
}