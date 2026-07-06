import { Box, Stack } from "@mui/material"
import DayColumn from "./DayColumn"

export default function WeeklyGrid({weekStart, weekEnd, days, tasks, onToggle}) {
  return (
    <Stack 
      direction="row" 
      sx = {{
        border: 1,
        borderColor: "divider"
      }}>
      {days.map((day, index) => (
        <Box
          key={day.format("YYYY-MM-DD")}
          sx={{ 
            flex: 1
          }}
          
        >
          <DayColumn 
            day={day}
            tasks={tasks}
            onToggle={onToggle}
            isLast={index === days.length - 1}    
          />
        </Box>
      ))}
    </Stack>
  );
}