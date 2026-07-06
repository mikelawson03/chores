import dayjs from "dayjs";
import PlannerToolbar from "../components/planner/PlannerToolbar";
import WeeklyGrid from "../components/planner/WeeklyGrid";
import { useState } from "react";
import isoWeek from "dayjs/plugin/isoWeek";
import { Stack } from "@mui/material";


export default function WeeklyPlanner({ tasks, onToggle }) {
  dayjs.extend(isoWeek);
  const [currentWeek, setCurrentWeek] = useState(dayjs())
  
  const weekStart = currentWeek.startOf("isoWeek")
  const weekEnd = currentWeek.endOf("isoWeek")

  const days = Array.from(
     { length: 7 },
     (_, i) => weekStart.add(i, "day")
  )

  const handlePreviousWeek = () => {
    setCurrentWeek(currentWeek.subtract(1, "week"))
  };

  const handleNextWeek = () => {
    setCurrentWeek(currentWeek.add(1, "week"))
  };

  return (
    <Stack spacing={4}>
      <PlannerToolbar 
        weekStart={weekStart} 
        weekEnd={weekEnd}
        onPreviousWeek={handlePreviousWeek}
        onNextWeek={handleNextWeek}
      />
      <WeeklyGrid
        weekStart={weekStart}
        weekEnd={weekEnd}
        days={days}
        tasks={tasks}
        onToggle={onToggle}
      />
    </Stack>
  )
}