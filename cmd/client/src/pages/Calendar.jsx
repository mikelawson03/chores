import dayjs from "dayjs";
import isoWeek from "dayjs/plugin/isoWeek";
import CalendarToolbar from "../components/calendar/CalendarToolbar";
import CalendarContent from "../components/calendar/CalendarContent";
import DailyAgenda from "../components/dailyAgenda/DailyAgenda";
import { useState } from "react";
import { Box, CircularProgress, Stack } from "@mui/material";
import { getActiveTasks, getMonthlyTasks, getUnscheduledTasks, getTasksDueInMonth, getWeeklyTasks, removeCompletedTasks } from "../utils/taskHelpers";
import { useAuth } from "../auth/useAuth";
import { useQuery } from "@tanstack/react-query";
import { getAssignments } from "../utils/assignmentHelpers";

export default function Calendar({ toggleTaskComplete }) {
  dayjs.extend(isoWeek);
  const { user } = useAuth();
  const { 
    data: tasks = [],
    isPending,
    isError,
    error,
  } = useQuery({
    queryKey: ["assignments", user?.id],
    queryFn: () => getAssignments(user),
    enabled: !!user,
  });

  const MAX_CALENDAR_DAY_ITEMS = 4;
  const [currentDate, setCurrentDate] = useState(dayjs());
  const [agendaDate, setAgendaDate] = useState(dayjs())
  const [dailyAgendaOpen, setDailyAgendaOpen] = useState(false)

  const monthDisplayStart = currentDate.date(1).startOf("isoWeek")
  
  const days = Array.from(
    { length: 42 },
    (_, i) => monthDisplayStart.add(i, "day")
  )

  const calendarDays = days.map(day => ({
    day,
    tasks: tasks.filter(
      task => dayjs(task.scheduledFor).isSame(day, "day") && !task.completed && !task.canceled
    )
  }));

  
  let filterTasks = getActiveTasks(tasks);
  filterTasks = getUnscheduledTasks(filterTasks);
  filterTasks = getTasksDueInMonth(currentDate, filterTasks);

  const weeklyTasks = getWeeklyTasks(filterTasks);
  const monthlyTasks = getMonthlyTasks(filterTasks);

  const handlePreviousMonth = () => {
    setCurrentDate(currentDate.subtract(1, "month"))
  };

  const handleNextMonth = () => {
    setCurrentDate(currentDate.add(1, "month"));
  }

  const handleNextAgendaDay = () => {
    setAgendaDate(agendaDate.add(1, "day"))
  }

  const handlePreviousAgendaDay = () => {
    setAgendaDate(agendaDate.subtract(1, "day"))
  }

  const handleDailyAgendaClose = () => {
    setDailyAgendaOpen(false);
  }

  function handleDailyAgendaOpen(date) {
    setAgendaDate(date);
    setDailyAgendaOpen(true);
  }

  if (isPending){
    return (
      <Box sx ={{
          width: "100%",
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center"
        }}>
          <CircularProgress />
      </Box>
    )
  }

  return (
    <Stack spacing={0} direction="column" sx ={{ flex: 1, minHeight: 0}}>
      <CalendarToolbar 
        currentDate={currentDate}
        onPreviousMonth={handlePreviousMonth}
        onNextMonth={handleNextMonth}
      />
      <CalendarContent 
        currentDate={currentDate} 
        days={calendarDays} 
        maxDayItems={MAX_CALENDAR_DAY_ITEMS}
        weeklyTasks={weeklyTasks}
        monthlyTasks={monthlyTasks}
        onOverflowClick={handleDailyAgendaOpen}
      />
      <DailyAgenda 
        agendaDate={agendaDate}
        activeTasks={getActiveTasks(tasks)}
        dailyAgendaOpen={dailyAgendaOpen}
        toggleTaskComplete={toggleTaskComplete}
        onDailyAgendaClose={handleDailyAgendaClose}
        onPreviousAgendaDay={handlePreviousAgendaDay}
        onNextAgendaDay={handleNextAgendaDay}
      />
    </Stack>
  );
}