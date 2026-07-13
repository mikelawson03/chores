import { Box, Stack, Typography } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { clickableText } from "../../styles/typography";

export default function CalendarDay({currentDate, day, tasks, maxDayItems, onOverflowClick, openTaskDetails, isLastColumn, isLastRow, isCurrentMonth, isCurrentDate}) {
    const overflowTaskCount = tasks.length - maxDayItems
    return (
            <Box 
                onClick={() => onOverflowClick(day)}
                sx={{
                    borderRight: isLastColumn ? 0 : 1, 
                    borderBottom: isLastRow ? 0 : 1,
                    borderColor:"divider", p: 1,
                    backgroundColor: isCurrentMonth ? "background.paper" : "grey.200",
                    minWidth: 0,
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
                                            {task.title}
                                    </Typography>
                                </Box>
                            )
                        }
                        {overflowTaskCount > 0 && <Typography variant="body2" onClick={(e) => {e.stopPropagation(); onOverflowClick(day)}} sx={clickableText}>+{overflowTaskCount} more...</Typography>}
                    </Stack>
                </Stack>
            </Box>
    )
}