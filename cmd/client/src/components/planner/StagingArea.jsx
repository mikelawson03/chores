import { Stack, Typography } from "@mui/material"
import StagingSection from "./StagingSection"

export default function StagingArea({ weeklyTasks, monthlyTasks, onToggleComplete, openTaskDetails }) {
    return (
        <Stack direction="column">
            {weeklyTasks.length > 0 && (
             <StagingSection 
                title="This Week" 
                text="Test" 
                tasks={weeklyTasks} 
                onToggleComplete={onToggleComplete}
                openTaskDetails={openTaskDetails}
            />
            )}
             {monthlyTasks.length > 0 && (
                <StagingSection 
                    title="This Month" 
                    text="Test" 
                    tasks={monthlyTasks} 
                    onToggleComplete={onToggleComplete}
                    openTaskDetails={openTaskDetails}
                />
            )}
        </Stack>
    )
}