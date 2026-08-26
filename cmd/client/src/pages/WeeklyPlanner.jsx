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
import { DragDropProvider } from "@dnd-kit/react";
import { useNotificationStore } from "../stores/notificationStore";
import { queryClient } from "../query/queryClient";

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
  
  const weekStart = currentWeek.startOf("isoWeek");
  const weekEnd = currentWeek.endOf("isoWeek");

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

  const rescheduleMutation = useMutation({
    mutationFn: rescheduleTask,

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["assignments", user.id],
      });
    },

    onError: (error) => {
      console.log(error);
    }
  })

  function handleTaskDrop(event) {
    const taskId = event.operation.source?.id;
    const newScheduledFor = event.operation.target?.id;

    if (!taskId || !newScheduledFor) {
      return;
    }

    const task = tasks.find(task => task.id === taskId);

    if (!task) {
      return;
    }

    const newDate = dayjs(newScheduledFor);
    const dueDate = dayjs(task.dueDate);

    console.log(newDate.isAfter(dueDate));
    
    if (newDate.isAfter(dueDate)) {
      showErrorNotification("Cannot schedule after due date.")
      return;
    }

    rescheduleMutation.mutate({
      id: taskId,
      scheduledFor: newDate.endOf("day").toISOString(),
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
      onDragEnd={(event) => {
        if (event.canceled) return;

        handleTaskDrop(event)
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
    </DragDropProvider>
  )
}