import { Stack } from "@mui/material"
import StagingSection from "./StagingSection"

export default function StagingArea({ weeklyTasks, monthlyTasks, toggleTaskComplete }) {
    return (
        <Stack direction="column">
            {weeklyTasks.length > 0 && (
             <StagingSection 
                title="This Week" 
                text="Test" 
                tasks={weeklyTasks} 
                toggleTaskComplete={toggleTaskComplete}
            />
            )}
             {monthlyTasks.length > 0 && (
                <StagingSection 
                    title="This Month" 
                    text="Test" 
                    tasks={monthlyTasks} 
                    toggleTaskComplete={toggleTaskComplete}
                />
            )}
        </Stack>
    )
}