import AppLayout from "./layouts/AppLayout";
import Calendar from "./pages/Calendar";
import Chores from "./pages/Chores";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import Settings from "./pages/Settings"
import WeeklyPlanner from "./pages/WeeklyPlanner";
import TaskDetails from "./components/TaskDetails"
import { Box, CssBaseline, Stack } from "@mui/material";
import { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom"
import dayjs from "dayjs";
import { getTasks, getTasksForUser, updateTask } from "./api/tasks";
import { tasksEqual } from "./utils/taskHelpers";
import { useAuth } from "./auth/useAuth";
import ProtectedRoute from "./auth/ProtectedRoute";
import AdminRoute from "./auth/AdminRoute";


function App() {

  const { user } = useAuth();
  const [taskDetailsOpen, setTaskDetailsOpen] = useState(false);
  const [selectedTask, setSelectedTask] = useState();
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    if (!user) {
      return;
    }

    const load = async() => {
      const fetchedTasks = await loadTasks();
      console.log(fetchedTasks);
      setTasks(fetchedTasks);
    };

    load();
  }, [user])

  const loadTasks = async () => {
    if (user.role === "admin") {
      return await getTasks();
    }

    return await getTasksForUser(user.id);
  }

  function openTaskDetails(task) {
    setSelectedTask(task);
    setTaskDetailsOpen(true);
  }

  async function finishTaskEditing(task) {
    await saveTask(task, selectedTask)
    setTaskDetailsOpen(false);
  }

  async function saveTask(updatedTask, originalTask = null) {
    if (tasksEqual(updatedTask, originalTask)){
      return;
    }

    const savedTask = await updateTask(updatedTask)

    setTasks(current =>
      current.map(t =>
        t.id === savedTask.id ? savedTask : t
      )
    );

    return savedTask;
  }

  // TODO: 
  // - expand to add activity log entry, 
  // - save to backend
  // - render undo toast
  function toggleTaskComplete(task) {   
    task.completed = !task.completed;

    if (!task.scheduledFor && task.completed) {
      task.scheduledFor = dayjs().format();
    }

    saveTask(task);
  }

  return (
    <>
        <Box 
          sx={{
            flex: 1,
            minHeight: 1,
            display: "flex",
            flexDirection: "column"
          }}
        >
        <Routes>
          <Route path ="/login" element={<Login />} />
          <Route element={<ProtectedRoute />}>
            <Route element={<AppLayout />}>
              <Route path="/" element={<Dashboard 
                tasks={tasks} 
                toggleTaskComplete={toggleTaskComplete}
                openTaskDetails={openTaskDetails}
              />} />

              <Route path="/calendar" element={<Calendar 
                tasks={tasks} 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              />} />

              <Route path="/planner" element={<WeeklyPlanner 
                tasks={tasks} 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              />} />
              <Route element={<AdminRoute />}>
                <Route path="/chores" element={<Chores />} />
                <Route path="/admin" element={<Settings />} />
              </Route>
            </Route>
          </Route>
        </Routes>
        
        {selectedTask && (
          <TaskDetails 
            task={selectedTask} 
            open={taskDetailsOpen}
            finishTaskEditing={finishTaskEditing}
          />
        )}
        </Box>
    </>

  );
}

export default App;

