import { Box, Stack, Typography } from "@mui/material";
import dayjs from "dayjs";
import { useTaskStore } from "../../stores/taskStore";
import DraggableTask from "../dragAndDrop/DraggableTask";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";
import { clickableSurface } from "../../styles/surfaces";

export default function TaskShelfPanel({ tasks }) {
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    return(
        <Stack spacing={2} 
            sx={{pt: 2}}
        >
            {tasks.map( task => 
                <DraggableTask task={task} key={task.id}>
                    <Box 
                        key={task.id}
                        onClick = {() => openTaskDetails(task)} 
                        sx={[clickableSurface, { 
                            justifyContent: "center", 
                            position: "relative",
                            overflow: "hidden",
                            p: 1,
                            "&::before": {
                                content: '""',
                                position: "absolute",
                                left: 0,
                                top: 1,
                                bottom: 1,
                                width: 4,
                                backgroundColor: HOUSEHOLD_USER_COLORS[task.userColorOption],
                            }
                        }]}
                    >
                        <Typography 
                            variant="body1"
                            noWrap={true}
                            sx={{
                                textOverflow: "ellipsis",
                                overflow: "hidden",
                            }}
                        >
                            {dayjs(task.dueDate).format("MMM DD")} - {task.templateName}
                        </Typography>
                    </Box>
                </DraggableTask>
            )}
        </Stack>
    )
}