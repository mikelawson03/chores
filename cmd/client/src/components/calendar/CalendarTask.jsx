import { Box, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { useTaskStore } from "../../stores/taskStore";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";
import { CADENCES } from "../../constants/cadences"

export default function CalendarTask( {task} ) {
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    const ownerColor = HOUSEHOLD_USER_COLORS[task.userColorOption].bgColor
    const cadenceColor = CADENCES[task.cadence].outlineColor

    return(
        <Box 
            key={task.id} 
            onClick = {(e) => {e.stopPropagation(); openTaskDetails(task)}} 
            
            sx={ 
                [clickableSurface,{
                    justifyContent: "center", 
                    position: "relative",
                    overflow: "hidden",
                    border: 1,
                    borderColor: cadenceColor,
                    borderRadius: 1,
                    

                    "&::before": {
                        content: '""',
                        position: "absolute",
                        left: 0,
                        top: 1,
                        bottom: 1,
                        width: 4,
                        backgroundColor: ownerColor,
                    }
                }]
            }
        >
            <Typography 
                variant="body2" 
                noWrap={true}
                sx={{
                    textOverflow: "ellipsis",
                    overflow: "hidden",
                    px: 1
                }}
            >
                    {task.templateName}
            </Typography>
        </Box>
    )
}