import { Stack, Typography } from "@mui/material"
import StagingSection from "./StagingSection"

export default function StagingArea({ weeklyTasks, monthlyTasks, onToggleComplete }) {
    return (
        <Stack direction="column">
            {weeklyTasks.length > 0 && (
             <StagingSection 
                title="This Week" 
                text="Test" 
                tasks={weeklyTasks} 
                onToggleComplete={onToggleComplete}
            />
            )}
             {monthlyTasks.length > 0 && (
                <StagingSection 
                    title="This Month" 
                    text="Test" 
                    tasks={monthlyTasks} 
                    onToggleComplete={onToggleComplete}
                />
            )}
        </Stack>
    )
}