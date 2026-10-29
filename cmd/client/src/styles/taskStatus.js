export function getTaskStatusStyles(task) {
    if (task.canceled) {
        return {
            bgColor: "grey.100",
            txtColor: "text.disabled",
            txtDecoration: "line-through",
        };
    }

    if (task.completed) {
        return {
            bgColor: "grey.200",
            txtColor: "text.secondary",
            txtDecoration: "none"
        };
    }

    return {
        bgColor: "background.paper",
        txtColor: "text.primary",
        txtDecoration: "none",
    };
}