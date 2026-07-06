import dayjs from "dayjs";
import PlannerToolbar from "../components/planner/PlannerToolbar";
import WeeklyGrid from "../components/planner/WeeklyGrid";
import { useState } from "react";
import isoWeek from "dayjs/plugin/isoWeek";
import { Stack } from "@mui/material";
import StagingArea from "../components/planner/StagingArea";


export default function WeeklyPlanner({ tasks, onToggleComplete }) {
  dayjs.extend(isoWeek);
  const [currentWeek, setCurrentWeek] = useState(dayjs());
  
  const weekStart = currentWeek.startOf("isoWeek");
  const weekEnd = currentWeek.endOf("isoWeek");

  const days = Array.from(
     { length: 7 },
     (_, i) => weekStart.add(i, "day")
  );

  const plannerDays = days.map(day => ({
    day,
    tasks: tasks.filter(
      task => task.scheduledFor === day.format("YYYY-MM-DD")
    ),
  }));

  const weeklyBacklog = tasks.filter(
    task => task.cadence === "weekly" && !task.scheduledFor && !task.completed
  );

  const monthlyBacklog = tasks.filter(
    task => task.cadence === "monthly" && !task.scheduledFor && !task.completed
  );

  const handlePreviousWeek = () => {
    setCurrentWeek(currentWeek.subtract(1, "week"))
  };

  const handleNextWeek = () => {
    setCurrentWeek(currentWeek.add(1, "week"))
  };

  return (
    <Stack spacing={4} direction="column" sx={{height: "100%"}}>
      <PlannerToolbar 
        weekStart={weekStart} 
        weekEnd={weekEnd}
        onPreviousWeek={handlePreviousWeek}
        onNextWeek={handleNextWeek}
      />
      <WeeklyGrid
        weekStart={weekStart}
        weekEnd={weekEnd}
        plannerDays = {plannerDays}
        onToggleComplete={onToggleComplete}
        sx = {{
          flex: 1,
          minHeight:600
        }}
      />
      <StagingArea 
        weeklyTasks = {weeklyBacklog}
        monthlyTasks = {monthlyBacklog}
        onToggleComplete={onToggleComplete}
        />
    </Stack>
  )
}