import { Box, Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";
import DraggableTask from "../dragAndDrop/DraggableTask";

export default function DayColumn({ day, tasks, toggleTaskComplete, isLast}) {

  return (

    <Stack 
      direction="column" 
      sx={{
        flex: 1,
        height: "100%",
        borderRight: isLast ? 0 : 1,
        borderColor: "divider"
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
        <Typography variant="h6">
          {day.format("ddd")}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {day.format("MMM D")}
        </Typography>
      </Box>
    <Box
      sx={{
        flex: 1,
        p: 1,
        display: "flex",
        flexDirection: "column",
        gap: 0.75,
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