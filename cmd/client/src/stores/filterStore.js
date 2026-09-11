import { create } from "zustand";

const filterStateKeys = {
    users: "hiddenUserIds",
    cadences: "hiddenCadences",
    statuses: "hiddenStatuses",
};

export const useFilterStore = create((set) => ({
    hiddenUserIds: new Set(),
    hiddenCadences: new Set(),
    hiddenStatuses: new Set(),

    toggleFilterItem: (filterType, value) => {
        const stateKey = filterStateKeys[filterType];
        set(state => {
            const next = new Set(state[stateKey]);
            if (next.has(value)) {
                next.delete(value);
            } else {
                next.add(value);
            }
            return {
                [stateKey]: next,
            };
        });
    },

    toggleAllFilters: (filterType, values) => {
        const stateKey = filterStateKeys[filterType];

        set(state => {
            if (state[stateKey].size === 0) {
                return {[stateKey]: new Set(values)};
            } else {
                return {[stateKey]: new Set()};
            }
        });
    },

    initializeFilters: (config) => {
        set(() => ({
            hiddenUserIds: new Set(config.hiddenUserIds),
            hiddenCadences: new Set(config.hiddenCadences),
            hiddenStatuses: new Set(config.hiddenStatuses)
        }))
    }
}))