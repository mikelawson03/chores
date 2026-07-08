import AppLayout from "./layouts/AppLayout";
import Dashboard from "./pages/Dashboard";
import Sidebar from "./components/Sidebar";
import WeeklyPlanner from "./pages/WeeklyPlanner";
import TaskDetails from "./components/TaskDetails"
import { Box, Stack } from "@mui/material";
import { useState } from "react";

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
      scheduledFor: "",
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
      completed: true,
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
    onToggleComplete(task.id);
  }

  function toggleTaskCancel(task) {
    onToggleCancel(task.id)
  }
  
  // TODO:
  // make cancel API call once back end connected
  function onToggleCancel(id) {
    return setTasks(currentTasks => {
      currentTasks.map(task => {
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

  return (
    // <AppLayout>
    //     <Stack direction="row">
    //       <Sidebar />
    //       <Box sx={{ flexGrow: 1 }}>
    //         <Dashboard tasks={tasks} onToggleComplete={onToggleComplete}/>
    //       </Box>
    //     </Stack>
    // </AppLayout>
   
    //    {/* Temporary until routing determines layout variant. */}
    //   
    <AppLayout maxWidth={false}>
      <Stack direction="row">
        <Sidebar />
          <Box sx={{ flexGrow: 1 }}>
            <WeeklyPlanner 
              tasks={tasks} 
              onToggleComplete={onToggleComplete} 
              openTaskDetails={openTaskDetails}
            />
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
    // 
        

  );
}

export default App;

