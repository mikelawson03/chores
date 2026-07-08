import { Stack, Typography } from "@mui/material"
import StagingSection from "./StagingSection"

export default function StagingArea({ weeklyTasks, monthlyTasks, toggleTaskComplete, openTaskDetails }) {
    return (
        <Stack direction="column">
            {weeklyTasks.length > 0 && (
             <StagingSection 
                title="This Week" 
                text="Test" 
                tasks={weeklyTasks} 
                toggleTaskComplete={toggleTaskComplete}
                openTaskDetails={openTaskDetails}
            />
            )}
             {monthlyTasks.length > 0 && (
                <StagingSection 
                    title="This Month" 
                    text="Test" 
                    tasks={monthlyTasks} 
                    toggleTaskComplete={toggleTaskComplete}
                    openTaskDetails={openTaskDetails}
                />
            )}
        </Stack>
    )
}