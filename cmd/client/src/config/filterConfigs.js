export const PLANNER_FILTER_CONFIG = {
    hiddenUserIds: new Set(),
    hiddenCadences: new Set(),
    hiddenStatuses: new Set(),
};

export const CALENDAR_FILTER_CONFIG = {
    hiddenUserIds: new Set(),
    hiddenCadences: new Set(),
    hiddenStatuses: new Set(["completed", "canceled"]),
}