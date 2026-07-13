import AppLayout from "./layouts/AppLayout";
import Calendar from "./pages/Calendar";
import Chores from "./pages/Chores";
import Dashboard from "./pages/Dashboard";
import Sidebar from "./components/Sidebar";
import WeeklyPlanner from "./pages/WeeklyPlanner";
import TaskDetails from "./components/TaskDetails"
import { Box, CssBaseline, Stack } from "@mui/material";
import { useState } from "react";
import dayjs from "dayjs";

function App() {

  const [open, setOpen] = useState(false)
  const [activeTask, setActiveTask] = useState()

  const [tasks, setTasks] = useState([
    {
      id: 1,
      title: "Feed dog",
      instructions: "Feed Ruby 1 cup of food every morning",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      notes: "Ruby only ate half her food this morning",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
    {
      id: 2,
      title: "Make dinner",
      assignee: "Mike",
      cadence: "daily",
      duration: "1 hour",
      notes: "",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
    {
      id: 3,
      title: "Wash dishes",
      assignee: "Mike",
      cadence: "daily",
      duration: "30 mins",
      notes: "",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
    {
      id: 4,
      title: "Vacuum rugs",
      assignee: "Mike",
      cadence: "weekly",
      duration: "30 mins",
      notes: "",
      scheduledFor: "2026-07-11",
      dueDate: "2026-07-11",
      completed: false,
      canceled: false,
    },
    {
      id: 5,
      title: "Clean fish tank",
      assignee: "Mike",
      cadence: "weekly",
      duration: "30 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-10",
      completed: false,
      canceled: false,
    },
    {
      id: 6,
      title: "Clean cat litter",
      assignee: "Mike",
      cadence: "weekly",
      duration: "20 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-10",
      completed: false,
      canceled: false,
    },
    {
      id: 7,
      title: "Change CPAP filter",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-31",
      completed: false,
      canceled: false,
    },
    {
      id: 8,
      title: "Change house water filter",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-31",
      completed: false,
      canceled: false,
    },
    {
      id: 9,
      title: "Test",
      assignee: "Mike",
      cadence: "monthly",
      duration: "10 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-31",
      completed: false,
      canceled: false,
    },
    {
      id: 10,
      title: "Clean out refrigerator",
      assignee: "Mike",
      cadence: "monthly",
      duration: "45 mins",
      notes: "",
      scheduledFor: "",
      dueDate: "2026-07-31",
      completed: true,
      canceled: false,
    },
    {
      id: 11,
      title: "Put away leftovers",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      notes: "",
      scheduledFor: "2026-07-02",
      dueDate: "2026-07-31",
      completed: true,
      canceled: false,
    },
    {
      id: 12,
      title: "Get water tested",
      assignee: "Mike",
      cadence: "one-off",
      duration: "10 mins",
      notes: "",
      scheduledFor: "2026-07-02",
      dueDate: "2026-07-02",
      completed: true,
      canceled: false,
    },
    {
      id: 13,
      title: "Test",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      notes: "",
      scheduledFor: "2026-07-01",
      dueDate: "2026-07-01",
      completed: true,
      canceled: false,
    },
    {
      id: 14,
      title: "Eat breakfast",
      instructions: "Eat some oatmeal",
      assignee: "Mike",
      cadence: "daily",
      duration: "10 mins",
      notes: "",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
    {
      id: 15,
      title: "Take a shower",
      instructions: "Clean ya ass",
      assignee: "Mike",
      cadence: "one-off",
      duration: "10 mins",
      notes: "Ruby only ate half her food this morning",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
    {
      id: 16,
      title: "Get a passport photo",
      instructions: "Picture time!",
      assignee: "Mike",
      cadence: "one-off",
      duration: "60 mins",
      notes: "Ruby only ate half her food this morning",
      scheduledFor: "2026-07-04",
      dueDate: "2026-07-04",
      completed: false,
      canceled: false,
    },
  ])

  function openTaskDetails(task) {
    setActiveTask(task)
    setOpen(true);
  }

  function closeTaskDetails(id, notes) {
    setOpen(false);
    saveNotes(id, notes);
  }

  // TODO: 
  // - expand to add activity log entry, 
  // - save to backend
  // - render undo toast
  function toggleTaskComplete(task) {
    if (!task.scheduledFor) {
      updateScheduledFor(task.id, dayjs())
    }
    onToggleComplete(task.id);
  }

  function toggleTaskCancel(task) {
    onToggleCancel(task.id)
  }
  
  // TODO:
  // make cancel API call once back end connected
  // Back end will own canceledAt data; will assign to current state on API return
  function onToggleCancel(id) {
    setTasks(currentTasks => {
      return currentTasks.map(task => {
        if (task.id == id ) {
          return {
            ...task,
            canceled: !task.canceled
          };
        }
        return task;
      })
    });
  }

  // TODO:
  // make complete API call once back end connected
  // Back end will own completedAt data; will assign to current state on API return
  function onToggleComplete(id) {
    setTasks(currentTasks => {
      return currentTasks.map(task => {
        if (task.id == id) {
          return {
            ...task,
            completed: !task.completed
          };
        }
        return task;
      }
      )
    });
  }
  // Back end will own updatedAt data; will assign to current state on API return
  function saveNotes(id, newNotes) {
    // TODO:
    // add Notes section to assignment table
    // persist Notes changes to backend
    setTasks(currentTasks => {
      return currentTasks.map(task => {
        if (task.id == id) {
          if (task.notes !== newNotes) {
            return {
              ...task,
              notes: newNotes
            };
          }
        }
        return task;
      })
    }
    )
  }

  function updateScheduledFor(id, scheduledFor) {
    setTasks(currentTasks => {
      return currentTasks.map(task => {
        if (task.id == id) {
          return {
            ...task,
            scheduledFor: scheduledFor.format("YYYY-MM-DD")
          };
        }
        return task;
      })
    })
  }

  return (
    
    // <AppLayout>
    //     <Stack direction="row">
    //       <Sidebar />
    //       <Box sx={{ flexGrow: 1 }}>
    //         <Dashboard 
    //           tasks={tasks} 
    //           toggleTaskComplete={toggleTaskComplete}
    //           openTaskDetails={openTaskDetails}
    //         />
    //       </Box>
    //       {activeTask && <TaskDetails 
    //       task={activeTask} 
    //       open={open}
    //       closeTaskDetails={closeTaskDetails}
    //       toggleTaskComplete={toggleTaskComplete}
    //       toggleTaskCancel={toggleTaskCancel}
    //     />}
    //     </Stack>
    // </AppLayout>
   
    //    {/* maxWidth below is temporary until routing determines layout variant. */}
    <>
      <CssBaseline />

      <AppLayout maxWidth={false}>
        <Stack direction="row" sx={{ height: "100%" }}>
          <Sidebar />
            <Box sx={{ flexGrow: 1 }}>
              {/* <Chores /> */}
              <Calendar 
                tasks={tasks} 
                openTaskDetails={openTaskDetails}
                toggleTaskComplete={toggleTaskComplete}
              />
              {/* <WeeklyPlanner 
                tasks={tasks} 
                toggleTaskComplete={toggleTaskComplete} 
                openTaskDetails={openTaskDetails}
              /> */}
            </Box>
          {activeTask && <TaskDetails 
            task={activeTask} 
            open={open}
            closeTaskDetails={closeTaskDetails}
            toggleTaskComplete={toggleTaskComplete}
            toggleTaskCancel={toggleTaskCancel}
          />}
        </Stack>
      </AppLayout>
    </>
  );
}

export default App;

