import dayjs from "dayjs";
import isoWeek from "dayjs/plugin/isoWeek";
import CalendarToolbar from "../components/calendar/CalendarToolbar";
import CalendarContent from "../components/calendar/CalendarContent";
import DailyAgenda from "../components/dailyAgenda/DailyAgenda";
import { useState } from "react";
import { Box, CircularProgress, Stack } from "@mui/material";
import { getActiveTasks, getMonthlyTasks, getUnscheduledTasks, getTasksDueInMonth, getWeeklyTasks } from "../utils/taskHelpers";
import { useAuth } from "../auth/useAuth";
import { useMutation, useQuery } from "@tanstack/react-query";
import { getAssignments, rescheduleTask } from "../utils/assignmentHelpers";
import { DragDropProvider, DragOverlay } from "@dnd-kit/react";
import { useNotificationStore } from "../stores/notificationStore";
import { queryClient } from "../query/queryClient";
import { parseApiError } from "../utils/errorHelpers";
import CalendarTask from "../components/calendar/CalendarTask";

export default function Calendar({ toggleTaskComplete }) {
  dayjs.extend(isoWeek);
  const { user } = useAuth();
  const { 
    data: tasks = [],
    isPending,
  } = useQuery({
    queryKey: ["assignments", user?.id],
    queryFn: () => getAssignments(user),
    enabled: !!user,
  });

  const showErrorNotification = useNotificationStore(
    (state) => state.showNotification
  )

  const MAX_CALENDAR_DAY_ITEMS = 4;
  const [currentDate, setCurrentDate] = useState(dayjs());
  const [agendaDate, setAgendaDate] = useState(dayjs())
  const [dailyAgendaOpen, setDailyAgendaOpen] = useState(false)
  const [dragTask, setDragTask] = useState(null)

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

  const rescheduleMutation = useMutation({
    mutationFn: rescheduleTask,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["assignments", user.id],
      });
      setDragTask(null);
    },

    onError: (error) => {
      handleRescheduleError(error);
      setDragTask(null);
    }
  })

  function handleRescheduleError(error) {
    let parsed, message
    switch (error.status) {
      case 400:
        parsed = parseApiError(error)
        message = parsed.message[0].toUpperCase() + parsed.message.slice(1)
        showErrorNotification(message)
    }
  }

  function handleTaskDrop(event) {
    const newScheduledFor = event.operation.target?.id;
    const newScheduledDate = dayjs(newScheduledFor);
    const oldScheduledDate = dayjs(dragTask.scheduledFor);
    const dueDate = dayjs(dragTask.dueDate);

    if (!dragTask.id || !newScheduledFor) {
      return;
    }

    if (!dragTask) {
      return;
    }

    if (oldScheduledDate.isSame(newScheduledDate, "day")) {
      return;
    }

    if (dragTask.cadence === "daily") {
      showErrorNotification("Daily tasks cannot be rescheduled.")
      return;
    }

    if (newScheduledDate.isAfter(dueDate, "day")) {
      showErrorNotification("Cannot schedule after due date.")
      return;
    }

    rescheduleMutation.mutate({
      id: dragTask.id,
      scheduledFor: newScheduledDate.startOf("day").format()
    });
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
    <DragDropProvider
      onDragStart = {(event) => {
        const taskId = event.operation.source?.id;
        const task = tasks.find(task => task.id === taskId)

        setDragTask(task)
      }}

      onDragEnd={(event) => {
        if (event.canceled) return;

        handleTaskDrop(event);
        
      }}
    >
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
      <DragOverlay dropAnimation={null}>
        {dragTask && ( 
          <CalendarTask
            task={dragTask}
            width="100%"
          />
        )}
      </DragOverlay>
    </DragDropProvider>
  );
}