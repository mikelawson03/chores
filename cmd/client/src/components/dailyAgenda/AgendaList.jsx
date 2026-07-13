import { Box, Fade, Stack, Typography } from "@mui/material";
import AgendaCard from "./AgendaCard";
import { getScheduledTasksForDay } from "../../utils/taskFilters";

export default function AgendaList({ tasks, openTaskDetails, toggleTaskComplete }) {
    return (
        <Stack spacing= {2} sx={{flex: 1}}>
                {tasks.length === 0 && <Typography sx={{fontStyle: "italic"}}>No tasks scheduled for this date...</Typography>}
                {tasks.map( task => (
                    <AgendaCard 
                        key={task.id}
                        task={task}
                        openTaskDetails={openTaskDetails}
                        toggleTaskComplete={toggleTaskComplete}
                    />
                ))}
        </Stack>
    )
}