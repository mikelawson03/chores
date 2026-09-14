import { Card, Chip, Stack, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";
import { CADENCES } from "../../constants/cadences";


export default function AgendaCard({ task }) {
    
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    const ownerColor = HOUSEHOLD_USER_COLORS[task.userColorOption].bgColor
    return (
        <Card sx={[clickableSurface, {
            backgroundColor: "background.paper",
            position: "relative",
            overflow: "hidden",
            "&::before": {
                content: '""',
                position: "absolute",
                left: 0,
                top: 1,
                bottom: 1,
                width: 4,
                borderRadius: "0 4px 4px 0",
                backgroundColor: ownerColor,
            }
        }]} onClick={() => openTaskDetails(task)}> 
            <Stack sx={{ p: 2 }}>
                <Stack 
                    direction="row"
                    sx={{ 
                        width: "100%", 
                        justifyContent: "space-between", 
                        alignItems: "center"}}
                >   
                    <Typography 
                        variant="h6"
                        noWrap
                        sx = {{
                            overflow: "hidden",
                            textOverflow: "ellipsis",
                            flex: 1,
                            minWidth: 0,
                        }}>
                        {task.templateName}
                    </Typography>
                    <Chip 
                        label={CADENCES[task.cadence].cadenceBadge}
                        size="medium"
                        sx={{ 
                        flexShrink: 0 , 
                        bgcolor: CADENCES[task.cadence].backgroundColor,
                        color: CADENCES[task.cadence].color,
                        }}
                    />
                    
                </Stack>
                <Typography 
                    variant="body1" 
                    sx={{
                        color: "text.secondary"
                        }}
                >
                    Assigned to: {task.userDisplayName}
                </Typography>
                <Stack 
                    direction="row"
                    sx={{ width: "100%", justifyContent: "space-between", color: "text.secondary"}}
                >
                <Typography 
                    variant="body1"
                >
                    Duration: {task.duration}
                </Typography>
                {task.instructions === "" && <CompletionCheckbox
                    checked={task.completed}
                    taskId={task.id}
                />}
                </Stack>
                {task.instructions != "" && <Stack 
                    direction="row"
                    sx={{ width: "100%", justifyContent: "space-between", alignContent:"end", color: "text.secondary"}}
                >
                     <Stack sx={{pt: 1}}>
                        <Typography variant="body1">
                            Instructions: 
                        </Typography>
                        <Typography variant="body1">
                            {task.instructions}
                        </Typography>
                    </Stack>
                    <CompletionCheckbox 
                        checked={task.completed}
                        taskId={task.id}
                    />
                </Stack>}
            </Stack>
        </Card>
    )
}