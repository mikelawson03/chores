import { Stack, Typography } from "@mui/material";
import AgendaCard from "./AgendaCard";


export default function AgendaList({ tasks, toggleTaskComplete }) {
    return (
        <Stack spacing= {2} sx={{flex: 1}}>
                {tasks.length === 0 && <Typography sx={{fontStyle: "italic"}}>No tasks scheduled for this date...</Typography>}
                {tasks.map( task => (
                    <AgendaCard 
                        key={task.id}
                        task={task}
                        toggleTaskComplete={toggleTaskComplete}
                    />
                ))}
        </Stack>
    )
}