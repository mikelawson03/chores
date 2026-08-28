import dayjs from "dayjs";
import PlannerToolbar from "../components/planner/PlannerToolbar";
import WeeklyGrid from "../components/planner/WeeklyGrid";
import { useState } from "react";
import isoWeek from "dayjs/plugin/isoWeek";
import { Box, CircularProgress, Stack } from "@mui/material";
import StagingArea from "../components/planner/StagingArea";
import { useAuth } from "../auth/useAuth";
import { getAssignments, rescheduleTask } from "../utils/assignmentHelpers";
import { useMutation, useQuery } from "@tanstack/react-query";
import { DragDropProvider, DragOverlay } from "@dnd-kit/react";
import { useNotificationStore } from "../stores/notificationStore";
import { queryClient } from "../query/queryClient";
import { parseApiError } from "../utils/errorHelpers";
import PlannerTaskCard from "../components/planner/PlannerTaskCard";

export default function WeeklyPlanner({ toggleTaskComplete }) {
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

  const [currentWeek, setCurrentWeek] = useState(dayjs());
  const [dragTask, setDragTask] = useState(null);
  
  const weekStart = currentWeek.startOf("isoWeek");
  const weekEnd = currentWeek.endOf("isoWeek");
  const monthRange = {
    lower: weekStart.startOf("month").subtract(1,"day").endOf("day"),
    upper: weekEnd.endOf("month").add(1, "day").startOf("day")
  }

  const days = Array.from(
     { length: 7 },
     (_, i) => weekStart.add(i, "day")
  );
  
  const plannerDays = days.map(day => ({
    day,
    tasks: tasks.filter(
      task => dayjs(task.scheduledFor).isSame(day, "day")
    ),
  }));

  const weeklyBacklog = tasks.filter(
    task => task.cadence === "weekly" 
      && !task.scheduledFor 
      && !task.completed
      && dayjs(task.dueDate).isAfter(weekEnd.subtract(1, "week"))
      && dayjs(task.dueDate).isBefore(weekStart.add(1, "week"))
  );

  const monthlyBacklog = tasks.filter(
    task => task.cadence === "monthly" 
    && !task.scheduledFor 
    && !task.completed
    && dayjs(task.dueDate).isAfter(monthRange.lower)
    && dayjs(task.dueDate).isBefore(monthRange.upper)
  );

  const handlePreviousWeek = () => {
    setCurrentWeek(currentWeek.subtract(1, "week"))
  };

  const handleNextWeek = () => {
    setCurrentWeek(currentWeek.add(1, "week"))
  };

  const rescheduleMutation = useMutation({
    mutationFn: rescheduleTask,

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["assignments", user.id],
      });
    },

    onError: (error) => {
      handleRescheduleError(error)
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

    if (!dragTask.id || !newScheduledFor) {
      return;
    }

    if (oldScheduledDate.isSame(newScheduledDate, "day")) {
      return;
    }
    
    const dueDate = dayjs(dragTask.dueDate);
    
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
      scheduledFor: newScheduledDate.startOf("day").format(),
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
        const task = tasks.find(task => task.id === taskId);

        setDragTask(task ?? null);
      }}

      onDragEnd={(event) => {

        if (event.canceled) return;

        handleTaskDrop(event)
        setDragTask(null);
      }}
    >
      <Stack spacing={4} direction="column" sx={{height: "100%"}}>
        <PlannerToolbar 
          weekStart={weekStart} 
          weekEnd={weekEnd}
          onPreviousWeek={handlePreviousWeek}
          onNextWeek={handleNextWeek}
        />
        <WeeklyGrid
          plannerDays = {plannerDays}
          toggleTaskComplete={toggleTaskComplete}
          sx = {{
            flex: 1,
            minHeight:600
          }}
        />
        <StagingArea 
          weeklyTasks = {weeklyBacklog}
          monthlyTasks = {monthlyBacklog}
          toggleTaskComplete={toggleTaskComplete}
          />
      </Stack>
      <DragOverlay dropAnimation={null}>
        {dragTask && (
          <PlannerTaskCard 
            task={dragTask}
            toggleTaskComplete={toggleTaskComplete}
            width="100%"
          />
        )}
      </DragOverlay>
    </DragDropProvider>
  )
}