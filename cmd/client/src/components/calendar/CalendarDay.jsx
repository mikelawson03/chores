import { Box, Stack, Typography } from "@mui/material";
import { clickableText } from "../../styles/typography";
import CalendarTask from "./CalendarTask";
import DraggableTask from "../dragAndDrop/DraggableTask";

export default function CalendarDay({day, tasks, maxDayItems, onOverflowClick, isLastColumn, isLastRow, isCurrentMonth, isCurrentDate}) {
    const overflowTaskCount = tasks.length - maxDayItems
    
    return (
        <Box 
            onClick={() => onOverflowClick(day)}
            sx={{
                borderRight: isLastColumn ? 0 : 1, 
                borderBottom: isLastRow ? 0 : 1,
                borderColor:"divider", 
                p: 1,
                backgroundColor: isCurrentMonth ? "background.paper" : "grey.200",
                minWidth: 0,
                height: "100%",
                "&:hover": {
                    border: 1,
                    borderColor: "grey.400",
                }
            }}
        >
            <Stack direction="column">
                <Box sx={{
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "center",
                }}>
                    <Typography sx={{
                        backgroundColor: isCurrentDate ? "#99daff" : "none",
                        px: 0.5,
                        py: 0.25,
                        borderRadius: "50%",
                        
                    }}>
                        {day.format("DD")}
                    </Typography>
                </Box>
                <Stack spacing={0.5} >
                    {tasks
                        .slice(0, maxDayItems)
                        .map(task =>
                            <DraggableTask task={task} key={task.id}>
                                <CalendarTask task={task} />
                            </DraggableTask>
                        )
                    }
                    {overflowTaskCount > 0 && <Typography variant="body2" onClick={(e) => {e.stopPropagation(); onOverflowClick(day)}} sx={clickableText}>+{overflowTaskCount} more...</Typography>}
                </Stack>
            </Stack>
        </Box>
)
}