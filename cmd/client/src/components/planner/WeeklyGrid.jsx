import { Box, Stack } from "@mui/material"
import DayColumn from "./DayColumn"

export default function WeeklyGrid({plannerDays, toggleTaskComplete }) {
  return (    
    <Stack 
      direction="row" 
      sx = {{
        border: 1,
        borderColor: "divider",
        minHeight: 600
      }}>
      {plannerDays.map(({ day, tasks }, index) => (
        <Box
          key={day.format("YYYY-MM-DD")}
          sx={{ 
            flex: 1,
            
          }}
          
        >
          <DayColumn 
            day={day}
            tasks={tasks}
            toggleTaskComplete={toggleTaskComplete}
            isLast={index === plannerDays.length - 1}    
          />
        </Box>
      ))}
    </Stack>
  );
}