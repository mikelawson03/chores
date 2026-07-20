import AppLayout from "./layouts/AppLayout";
import Calendar from "./pages/Calendar";
import Chores from "./pages/Chores";
import Dashboard from "./pages/Dashboard";
import Sidebar from "./components/Sidebar";
import WeeklyPlanner from "./pages/WeeklyPlanner";
import TaskDetails from "./components/TaskDetails"
import { Box, CssBaseline, Stack } from "@mui/material";
import { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom"
import dayjs from "dayjs";
import { getTasks, updateTask } from "./api/tasks";
import { tasksEqual } from "./utils/taskHelpers";


function App() {

  const [taskDetailsOpen, setTaskDetailsOpen] = useState(false);
  const [selectedTask, setSelectedTask] = useState();
  const [tasks, setTasks] = useState([])
  

  useEffect(() => {
    async function loadTasks() {
      const fetchedTasks = await getTasks();
      setTasks(fetchedTasks);
    }
    loadTasks();
  }, [])

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

    console.log(task.scheduledFor)

    saveTask(task);
  }

  return (
    <>
      <CssBaseline />
      <AppLayout>
        <Box 
          sx={{
            flex: 1,
            minHeight: 1,
            display: "flex",
            flexDirection: "column"
          }}
        >
        <Routes>
          <Route path="/"
              element={
                <Dashboard 
                  tasks={tasks} 
                  toggleTaskComplete={toggleTaskComplete}
                  openTaskDetails={openTaskDetails}
              />
            }
          />

          <Route 
            path="/calendar"
            element={
              <Calendar 
                tasks={tasks} 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              />
            } 
          />

          <Route 
            path="/chores"
            element={
              <Chores />
            }
          />

          <Route 
            path="/planner"
            element={
              <WeeklyPlanner 
                tasks={tasks} 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              /> 
            }
          />

        </Routes>
        
        {selectedTask && (
          <TaskDetails 
            task={selectedTask} 
            open={taskDetailsOpen}
            finishTaskEditing={finishTaskEditing}
          />
        )}
        </Box>
        
      </AppLayout>
    </>

  );
}

export default App;

