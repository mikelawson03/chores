import dayjs from "dayjs";
import isoWeek from "dayjs/plugin/isoWeek";
import CalendarToolbar from "../components/calendar/CalendarToolbar";
import CalendarContent from "../components/calendar/CalendarContent";
import DailyAgenda from "../components/dailyAgenda/DailyAgenda";
import { useState } from "react";
import { Stack } from "@mui/material";
import { getActiveTasks, getMonthlyTasks, getUnscheduledTasks, getTasksDueInMonth, getWeeklyTasks, removeCompletedTasks } from "../utils/taskFilters";

export default function Calendar({ tasks, openTaskDetails, toggleTaskComplete }) {
  dayjs.extend(isoWeek);
  const MAX_CALENDAR_DAY_ITEMS = 4;
  const [currentDate, setCurrentDate] = useState(dayjs());
  const [agendaDate, setAgendaDate] = useState(dayjs())
  const [dailyAgendaOpen, setDailyAgendaOpen] = useState(false)

  const monthDisplayStart = currentDate.date(1).startOf("isoWeek")
  const monthDisplayEnd = monthDisplayStart.add(41, "day")
  
  const days = Array.from(
    { length: 42 },
    (_, i) => monthDisplayStart.add(i, "day")
  )

  const calendarDays = days.map(day => ({
    day,
    tasks: tasks.filter(
      task => task.scheduledFor === day.format("YYYY-MM-DD") && !task.completed && !task.canceled
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

  return (
    <Stack spacing={0} direction="column" sx ={{p: 3, height: "100%"}}>
      <CalendarToolbar 
        currentDate={currentDate}
        onPreviousMonth={handlePreviousMonth}
        onNextMonth={handleNextMonth}
      />
      <CalendarContent 
        currentDate={currentDate} 
        days={calendarDays} 
        maxDayItems={MAX_CALENDAR_DAY_ITEMS}
        openTaskDetails={openTaskDetails}
        weeklyTasks={weeklyTasks}
        monthlyTasks={monthlyTasks}
        onOverflowClick={handleDailyAgendaOpen}
        sx={{ flex: 1 }}
      />
      <DailyAgenda 
        agendaDate={agendaDate}
        activeTasks={getActiveTasks(tasks)}
        dailyAgendaOpen={dailyAgendaOpen}
        openTaskDetails={openTaskDetails}
        toggleTaskComplete={toggleTaskComplete}
        onDailyAgendaClose={handleDailyAgendaClose}
        onPreviousAgendaDay={handlePreviousAgendaDay}
        onNextAgendaDay={handleNextAgendaDay}
      />
    </Stack>
  );
}