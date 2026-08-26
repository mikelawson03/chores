import { Stack } from "@mui/material"
import DayColumn from "./DayColumn"
import DroppableDay from "../dragAndDrop/DroppableDay";

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
        <DroppableDay  key={day.format("YYYY-MM-DD")} day={day}>
          <DayColumn 
            day={day}
            tasks={tasks}
            toggleTaskComplete={toggleTaskComplete}
            isLast={index === plannerDays.length - 1}    
          />
        </DroppableDay>
        
      ))}
    </Stack>
  );
}