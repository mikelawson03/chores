import { Box, Stack, Typography } from "@mui/material";
import dayjs from "dayjs";

export default function TaskShelfPanel({ title, tasks, openTaskDetails }) {
    return(
        <Stack spacing={2} 
            sx={{pt: 2}}
        >
            {tasks.map( task => 
                <Box 
                    key={task.id}
                    onClick = {() => openTaskDetails(task)} 
                    sx={{ 
                        backgroundColor: "grey.100", 
                        justifyContent: "center", 
                        cursor: "pointer",
                        p: 1,
                        '&:hover': {
                            backgroundColor: "grey.200"
                        }
                    }}
                >
                    <Typography variant="body1">{dayjs(task.dueDate).format("MMM DD")} - {task.templateName}</Typography>
                </Box>
            )}
        </Stack>
    )
}