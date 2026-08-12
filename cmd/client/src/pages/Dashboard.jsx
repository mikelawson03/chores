import PageHeader from "../components/PageHeader";
import TaskCard from "../components/TaskCard";
import TaskListCard from "../components/TaskListCard";
import { Box, CircularProgress, Container, Grid, Stack, Typography } from "@mui/material"
import { getActiveTasks, getScheduledTasksForDay, getUnscheduledTasks, getWeeklyTasks, getMonthlyTasks, removeCompletedTasks, getCompletedTasks } from "../utils/taskHelpers";
import dayjs from "dayjs";
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "../auth/useAuth";
import { getAssignments } from "../utils/assignmentHelpers";

export default function Dashboard({openTaskDetails, toggleTaskComplete}) {
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

  const activeTasks = getActiveTasks(tasks)
  const activeAndUnscheduledTasks = getUnscheduledTasks(activeTasks)
  
  const weeklyTasks = getWeeklyTasks(activeAndUnscheduledTasks);
  const monthlyTasks = getMonthlyTasks(activeAndUnscheduledTasks);
  const todaysTasks = getScheduledTasksForDay(dayjs(), activeTasks);
  const completedTasks = getCompletedTasks(tasks);

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
  <Container maxWidth="lg">
    <PageHeader title="Dashboard" />
    <Stack spacing={2} direction={"row"} sx={{ marginBottom: 4, marginTop: 8}}>
      <TaskListCard 
        cardName="Weekly Tasks"
        tasks={weeklyTasks}
        maxItems={3}
        footerText="View planner →"
        toggleTaskComplete={toggleTaskComplete}
        openTaskDetails={openTaskDetails}
      />
      <TaskListCard 
        cardName="Monthly Tasks"
        tasks={monthlyTasks}
        maxItems={3}
        footerText="View planner →"
        toggleTaskComplete={toggleTaskComplete}
        openTaskDetails={openTaskDetails}
      />
      <TaskListCard 
        cardName="Completed Tasks"
        tasks={completedTasks}
        maxItems={3}
        footerText="View completed →"
        toggleTaskComplete={toggleTaskComplete}
        openTaskDetails={openTaskDetails}
      />
    </Stack>
    <Typography variant="h4" sx={{ marginBottom: 4, marginTop: 8}}>Today's Tasks</Typography>
    <Grid container spacing={5}>
    {todaysTasks.map(task => (
      <Grid size={3} key={task.id}>
        <TaskCard 
          task = {task}
          toggleTaskComplete={toggleTaskComplete}
          openTaskDetails={openTaskDetails}
          />
        </Grid>
    ))}
    </Grid>
  </Container>
  )
}