import { Box, Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";
import DraggableTask from "../dragAndDrop/DraggableTask";
import dayjs from "dayjs";

export default function DayColumn({ day, tasks, toggleTaskComplete, isLast}) {
  const isCurrentDate = day.isSame(dayjs(), "day");
  return (

    <Stack 
      direction="column" 
      sx={{
        flex: 1,
        height: "100%",
        borderRight: isLast ? 0 : 1,
        borderColor: "divider",
        minWidth: 0,
      }}
    >
      <Box
        sx ={{
          borderBottom: 1,
          borderColor: "divider",
          py: 1,
          display: "flex",
          alignItems: "center",
          flexDirection: "column"
        }}
      >
        <Typography variant="body2">
          {day.format("ddd")}
        </Typography>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            width: 35,
            height: 35,
            borderRadius: "50%",
            backgroundColor: isCurrentDate ? "#99daff" : "transparent",
          }}
        >
          <Typography variant="h6" color="text.secondary">
              {day.format("D")}
          </Typography>
        </Box>
      </Box>
    <Box
      sx={{
        flex: 1,
        p: 1,
        display: "flex",
        flexDirection: "column",
        gap: 0.75,
        minWidth: 0
      }}
    >
      {tasks.map( task => (
        <DraggableTask task={task} key={task.id}>
          <PlannerTaskCard   
            task={task}
            toggleTaskComplete={toggleTaskComplete}
            width="100%"
            />
        </DraggableTask>
      ) )}
    </Box>
    </Stack>
    )
  }