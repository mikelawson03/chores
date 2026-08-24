import { Box } from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import dayjs from "dayjs";
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
import { useTaskStore } from "./stores/taskStore";
import CloseDetailsAlert from "./components/CloseDetailsAlert";
import { useState } from "react";

function App() {

  const { user } = useAuth();

  const selectedTask = useTaskStore(
    (state) => state.selectedTask
  );

  const taskDraft = useTaskStore(
    (state) => state.taskDraft
  )

  const closeTaskDetails = useTaskStore(
    (state) => state.closeTaskDetails
  );

  const setTaskError = useTaskStore(
    (state) => state.setTaskError
  );

  const [dialogOpen, setDialogOpen] = useState(false)

  const updateTaskMutation = useMutation({
    mutationFn: updateAssignment,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["assignments", user.id],
      });
      closeTaskDetails();
    },
    onError: (error) => {
      handleTaskError(error);
    }
  });

  async function saveTask() {
    if (tasksEqual(taskDraft, selectedTask)){
      return;
    }

    const savedTask = await updateTaskMutation.mutateAsync(taskDraft);

    return savedTask;
  }

  function toggleTaskComplete(task) { 
    
    task.completed = !task.completed;

    if (!task.scheduledFor && task.completed) {
      task.scheduledFor = dayjs().format();
    }

    saveTask(task);
  }

  function handleTaskError(error) {
    console.log("Error status: ", error.status)
    switch (error.status){
      case 403:
        setTaskError("form", "You do not have permission to modify this assignment.")
        break;

      case 404:
        setTaskError("form", "This assignment no longer exists.")
        break;

      default:
        setTaskError("form", "An unexpected error occurred. Please try again.")
        break;
    }
  }

  function onTaskClose() {
    if (tasksEqual(taskDraft, selectedTask)){
      closeTaskDetails();
      return;
    }

    setDialogOpen(true);
    console.log(dialogOpen)
  }

  function onSave(task) {
    console.log(task);
    setDialogOpen(false);
    saveTask(task)
  }

  function onEdit() {
    setDialogOpen(false);
  }

  function onDiscard() {
    setDialogOpen(false);
    closeTaskDetails();
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
              />} />

              <Route path="/calendar" element={<Calendar 
                toggleTaskComplete={toggleTaskComplete}
              />} />

              <Route path="/planner" element={<WeeklyPlanner 
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
            saveTask={saveTask}
            onTaskClose={onTaskClose}
          />
        )}
        </Box>
        <CloseDetailsAlert 
          dialogOpen={dialogOpen}
          setDialogOpen={setDialogOpen}
          onSave={onSave}
          onEdit={onEdit}
          onDiscard={onDiscard}
          task={selectedTask}
        />
    </>

  );
}

export default App;

