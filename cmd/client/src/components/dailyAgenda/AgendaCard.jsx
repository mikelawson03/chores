import { Box, Card, Checkbox, Stack, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";


export default function AgendaCard({ task, openTaskDetails, toggleTaskComplete }) {
    return (
        <Card sx={clickableSurface} onClick={() => openTaskDetails(task)}> 
            <Stack spacing={3} sx={{ p: 2 }}>
                <Stack 
                    direction="row"
                    sx={{ 
                        width: "100%", 
                        justifyContent: "space-between", 
                        alignItems: "center"}}
                >
                    
                    <Typography variant="h6">
                        {task.title}
                    </Typography>
                    <Checkbox 
                        size="medium" 
                        onChange={() => toggleTaskComplete(task) } 
                        onClick={(event) => {event.stopPropagation();}}
                        sx={{ pr: 2, pl: 0 }}
                    />
                    
                </Stack>
                <Typography variant="body1">
                    {task.instuctions}
                </Typography>
                <Stack 
                    direction="row"
                    sx={{ width: "100%", justifyContent: "space-between"}}
                >
                    <Typography variant="body1">Assignee: {task.assignee}</Typography>
                    <Typography variant="body1">Duration: {task.duration}</Typography>

                </Stack>
            </Stack>
        </Card>
    )
}