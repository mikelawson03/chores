import AppLayout from "./layouts/AppLayout";
import Dashboard from "./pages/Dashboard";
import Sidebar from "./components/Sidebar";
import WeeklyPlanner from "./pages/WeeklyPlanner";
import TaskDetails from "./components/TaskDetails"
import { Box, Stack } from "@mui/material";
import { useState } from "react";

function App() {

  const [tasks, setTasks] = useState([
    {
      id: 1,
      title: "Feed dog",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      scheduledFor: "2026-07-04",
      completed: false
    },
    {
      id: 2,
      title: "Make dinner",
      assignee: "Mike",
      cadence: "daily",
      duration: "1 hour",
      scheduledFor: "2026-07-04",
      completed: false
    },
    {
      id: 3,
      title: "Wash dishes",
      assignee: "Mike",
      cadence: "daily",
      duration: "30 mins",
      scheduledFor: "2026-07-04",
      completed: false
    },
    {
      id: 4,
      title: "Vacuum rugs",
      assignee: "Mike",
      cadence: "weekly",
      duration: "30 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 5,
      title: "Clean fish tank",
      assignee: "Mike",
      cadence: "weekly",
      duration: "30 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 6,
      title: "Clean cat litter",
      assignee: "Mike",
      cadence: "weekly",
      duration: "20 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 7,
      title: "Change CPAP filter",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 8,
      title: "Change house water filter",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 9,
      title: "Test",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      scheduledFor: "",
      completed: false
    },
    {
      id: 10,
      title: "Clean out refrigerator",
      assignee: "Mike",
      cadence: "monthly",
      duration: "45 mins",
      scheduledFor: "",
      completed: true
    },
    {
      id: 11,
      title: "Put away leftovers",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      scheduledFor: "2026-07-02",
      completed: true
    },
    {
      id: 12,
      title: "Get water tested",
      assignee: "Mike",
      cadence: "one-off",
      duration: "10 mins",
      scheduledFor: "2026-07-02",
      completed: true
    },
    {
      id: 13,
      title: "Test",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      scheduledFor: "2026-07-01",
      completed: true
    },
  ])

  function onToggleComplete(id) {
    setTasks(
      tasks.map(task => {
        if (task.id == id) {
          return {
            ...task,
            completed: !task.completed
          };
        }
        return task;
      }
      )
    );
  }

  return (
    // <AppLayout>
    //     <Stack direction="row">
    //       <Sidebar />
    //       <Box sx={{ flexGrow: 1 }}>
    //         <Dashboard tasks={tasks} onToggleComplete={onToggleComplete}/>
    //       </Box>
    //     </Stack>
    // </AppLayout>
    <AppLayout maxWidth={false}>
       {/* Temporary until routing determines layout variant. */}
      <Stack direction="row">
        <Sidebar />
        <Box sx={{ flexGrow: 1 }}>
          <WeeklyPlanner tasks={tasks} onToggleComplete={onToggleComplete}/>
        </Box>
        <TaskDetails />
      </Stack>
    </AppLayout>
  );
}

export default App;

