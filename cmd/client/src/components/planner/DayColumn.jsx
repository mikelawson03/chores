import { Box, Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";

export default function DayColumn({ day, tasks, onToggle, isLast}) {
  const dailyTasks = tasks.filter(
    task => task.scheduled_for === day.format("YYYY-MM-DD")
  );

  return (

    <Stack 
      direction="column" 
      sx={{
        flex: 1,
        borderRight: isLast ? 0 : 1,
        borderColor: "divider",
        minHeight: 600
      }}
    >
      <Box
        sx ={{
          borderBottom: 1,
          borderColor: "divider",
          py: 1,
          display: "flex",
          alignItems: "center",
          display: "flex",
          flexDirection: "column"
        }}
      >
        <Typography variant="h6">
          {day.format("ddd")}
        </Typography>
        <Typography variant="body2" color="te xt.secondary">
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
      {dailyTasks.map( task => (
        <PlannerTaskCard 
          key={task.id}
          id={task.id} 
          title={task.title}
          assignee={task.assignee}
          duration={task.duration}
          completed={task.completed}
          onToggle={onToggle}
          />
      ) )}
    </Box>
    </Stack>
    )
  }