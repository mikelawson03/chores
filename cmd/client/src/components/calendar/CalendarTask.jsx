import { Box, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { useTaskStore } from "../../stores/taskStore";

export default function CalendarTask( {task} ) {
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    return(
        <Box 
            key={task.id} 
            onClick = {(e) => {e.stopPropagation(); openTaskDetails(task)}} 
            
            sx={ 
                [clickableSurface,
                {justifyContent: "center", 
                    
                }]
            }
        >
            <Typography 
                variant="body2" 
                noWrap={true}
                sx={{
                    textOverflow: "ellipsis",
                    overflow: "hidden",
                }}
            >
                    {task.templateName}
            </Typography>
        </Box>
    )
}