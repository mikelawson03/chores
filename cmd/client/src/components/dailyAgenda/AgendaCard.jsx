import { Card, Stack, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";


export default function AgendaCard({ task }) {
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );
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
                    <CompletionCheckbox 
                        checked={task.completed}
                        taskId={task.id}
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