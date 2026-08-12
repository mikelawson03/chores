import { Box } from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import dayjs from "dayjs";
import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import AdminRoute from "./auth/AdminRoute";
import ProtectedRoute from "./auth/ProtectedRoute";
import { useAuth } from "./auth/useAuth";
import TaskDetails from "./components/TaskDetails";
import AppLayout from "./layouts/AppLayout";
import Calendar from "./pages/Calendar";
import Chores from "./pages/Chores";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import Settings from "./pages/Settings";
import WeeklyPlanner from "./pages/WeeklyPlanner";
import { queryClient } from "./query/queryClient";
import { updateAssignment } from "./utils/assignmentHelpers";
import { tasksEqual } from "./utils/taskHelpers";

function App() {

  const { user } = useAuth();
  const [taskDetailsOpen, setTaskDetailsOpen] = useState(false);
  const [selectedTask, setSelectedTask] = useState();

  const updateTaskMutation = useMutation({
    mutationFn: updateAssignment,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["assignments", user.id],
      });
    },
  });

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

    const savedTask = await updateTaskMutation.mutateAsync(updatedTask);

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
                toggleTaskComplete={toggleTaskComplete}
                openTaskDetails={openTaskDetails}
              />} />

              <Route path="/calendar" element={<Calendar 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              />} />

              <Route path="/planner" element={<WeeklyPlanner 
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

