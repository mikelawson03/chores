import { Card, Stack, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";


export default function AgendaCard({ task }) {
    
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    const ownerColor = HOUSEHOLD_USER_COLORS[task.userColorOption]
    return (
        <Card sx={[clickableSurface, {
            position: "relative",
            overflow: "hidden",
            "&::before": {
                content: '""',
                position: "absolute",
                left: 0,
                top: 6,
                bottom: 6,
                width: 4,
                borderRadius: "0 4px 4px 0",
                backgroundColor: ownerColor,
            }
        }]} onClick={() => openTaskDetails(task)}> 
            <Stack spacing={3} sx={{ p: 2 }}>
                <Stack 
                    direction="row"
                    sx={{ 
                        width: "100%", 
                        justifyContent: "space-between", 
                        alignItems: "center"}}
                >
                    
                    <Typography variant="h6">
                        {task.templateName}
                    </Typography>
                    <CompletionCheckbox 
                        checked={task.completed}
                        taskId={task.id}
                    />
                    
                </Stack>
                {task.instructions != "" && <Typography variant="body1">
                    {task.instuctions}
                </Typography>}
                <Stack 
                    direction="row"
                    sx={{ width: "100%", justifyContent: "space-between"}}
                >
                    <Typography variant="body1">Assignee: {task.userDisplayName}</Typography>
                    <Typography variant="body1">Duration: {task.duration}</Typography>

                </Stack>
            </Stack>
        </Card>
    )
}