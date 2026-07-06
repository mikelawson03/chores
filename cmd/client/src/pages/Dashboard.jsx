import PageHeader from "../components/PageHeader";
import TaskCard from "../components/TaskCard";
import TaskListCard from "../components/TaskListCard";
import { Grid, Stack, Typography } from "@mui/material"

export default function Dashboard({tasks, onToggleComplete}) {

  const weeklyTasks = tasks.filter(
    task => task.cadence === "weekly" && !task.completed
  );

  const monthlyTasks = tasks.filter(
    task => task.cadence === "monthly" && !task.completed
  );

  const completedTasks = tasks.filter(
    task => task.completed
  );

  const todaysTasks = tasks.filter(
    task => task.cadence === "daily" && !task.completed
  );

  return (
  <>
    <PageHeader title="Dashboard" />
    <Stack spacing={2} direction={"row"} sx={{ marginBottom: 4, marginTop: 8}}>
      <TaskListCard 
        cardName="Weekly Tasks"
        tasks={weeklyTasks}
        maxItems={3}
        footerText="View planner →"
        onToggleComplete={onToggleComplete}
      />
      <TaskListCard 
        cardName="Monthly Tasks"
        tasks={monthlyTasks}
        maxItems={3}
        footerText="View planner →"
        onToggleComplete={onToggleComplete}
      />
      <TaskListCard 
        cardName="Completed Tasks"
        tasks={completedTasks}
        maxItems={3}
        footerText="View completed →"
        onToggleComplete={onToggleComplete}
      />
    </Stack>
    <Typography variant="h4" sx={{ marginBottom: 4, marginTop: 8}}>Today's Tasks</Typography>
    <Grid container spacing={5}>
    {todaysTasks.map(task => (
      <Grid size={3} key={task.id}>
        <TaskCard 
          id = {task.id}
          title={task.title}
          assignee={task.assignee}
          duration={task.duration}
          completed={task.completed}
          onToggleComplete={onToggleComplete}
          />
        </Grid>
    ))}
    </Grid>
  </>
  )
}