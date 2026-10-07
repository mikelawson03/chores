import TaskCard from "../components/TaskCard";
import TaskListCard from "../components/TaskListCard";
import { Box, CircularProgress, Container, Grid, Stack, Typography } from "@mui/material"
import { getActiveTasks, getScheduledTasksForDay, getUnscheduledTasks, getCompletedTasks } from "../utils/taskHelpers";
import dayjs from "dayjs";
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "../auth/useAuth";
import { getAssignments } from "../utils/assignmentHelpers";
import { getActiveHouseholdUsers } from "../api/users";
import isoWeek from "dayjs/plugin/isoWeek";


export default function Dashboard({ toggleTaskComplete }) {
  dayjs.extend(isoWeek);
  const { user } = useAuth();
  const { 
    data: tasks = [],
    isPending,
  } = useQuery({
    queryKey: ["assignments", user?.user.id],
    queryFn: () => getAssignments(user),
    enabled: !!user,
  });

  const {
    data: householdUsers = [],
  } = useQuery({
    queryKey: ["activeHouseholdUsers", user?.householdId],
    queryFn: () => getActiveHouseholdUsers(user.householdId),
    enabled: !!user?.householdId && user?.role === "admin",
  });

  const dashboardUsers = user.role === "admin"
    ? householdUsers
    : [user];

  const householdUsersById = new Map(
    dashboardUsers.map(hhUser => [
      hhUser.user.id,
      hhUser
    ])
  )

  const dashboardTasks = tasks.map(task => {
    const hhUser = householdUsersById.get(task.userId);

    return {
      ...task,
      userDisplayName: hhUser?.displayName ?? task.userFirstName,
      userColorOption: hhUser?.colorOption ?? null,
    }
  })

  const weekStart = dayjs().startOf("isoWeek")
  const weekEnd = dayjs().endOf("isoWeek")
  const monthRange = {
    lower: weekStart.startOf("month").subtract(1,"day").endOf("day"),
    upper: weekEnd.endOf("month").add(1, "day").startOf("day")
  }

  const activeTasks = getActiveTasks(dashboardTasks)
  const activeAndUnscheduledTasks = getUnscheduledTasks(activeTasks)
  
  const weeklyTasks = activeAndUnscheduledTasks.filter(
    task => task.cadence === "weekly"
    && dayjs(task.dueDate).isAfter(weekEnd.subtract(1, "week"))
    && dayjs(task.dueDate).isBefore(weekStart.add(1, "week"))
  );

  const monthlyTasks = activeAndUnscheduledTasks.filter(
    task => task.cadence === "monthly"
    && dayjs(task.dueDate).isAfter(monthRange.lower)
    && dayjs(task.dueDate).isBefore(monthRange.upper)
  );

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
    <Stack sx={{p: 3, justifyContent: "center", alignItems:"center", width: "100%"}}>
    <Typography variant="h4" >Dashboard</Typography>
    </Stack>
    <Stack spacing={2} direction={"row"} sx={{ marginBottom: 4, marginTop: 8}}>
      <TaskListCard 
        cardTitle={"Weekly Tasks"}
        subHead={`${weekStart.format("MMM D")} - ${weekEnd.format("MMM D")}`}
        tasks={weeklyTasks}
        maxItems={3}
        footerText="View planner →"
        route="/planner"
        toggleTaskComplete={toggleTaskComplete}
      />
      <TaskListCard 
        cardTitle="Monthly Tasks"
        subHead={dayjs().format("MMMM")}
        tasks={monthlyTasks}
        maxItems={3}
        footerText="View planner →"
        route="/planner"
        toggleTaskComplete={toggleTaskComplete}
      />
      <TaskListCard 
        cardTitle="Completed Tasks"
        tasks={completedTasks}
        maxItems={3}
        footerText="View planner →"
        route="/planner"
        toggleTaskComplete={toggleTaskComplete}
      />
    </Stack>
    <Typography variant="h4" sx={{ marginBottom: 4, marginTop: 8}}>Today's Tasks</Typography>
    <Grid container spacing={5}>
    {todaysTasks.map(task => (
      <Grid size={3} key={task.id}>
        <TaskCard 
          task = {task}
          toggleTaskComplete={toggleTaskComplete}
          />
        </Grid>
    ))}
    </Grid>
  </Container>
  )
}