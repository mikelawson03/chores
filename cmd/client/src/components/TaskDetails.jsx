import { Box, Button, Drawer, Stack, TextField, Typography } from "@mui/material";
import DetailRow from "./details/DetailRow"
import { formatDuration, formatTimestamp } from "../utils/formatters";
import { useTaskStore } from "../stores/taskStore";
import dayjs from "dayjs";
import DetailRowDate from "./details/DetailRowDate";

export default function TaskDetails({task, saveTask, onTaskClose}) {
  const TASK_DETAIL_WIDTH=680
  const taskDraft = useTaskStore(
    (state) => state.taskDraft
  )
  const updateTaskDraft = useTaskStore(
    (state) => state.updateTaskDraft
  )
  const taskDetailsOpen = useTaskStore(
    (state) => state.taskDetailsOpen
  )
  const taskErrors = useTaskStore(
    (state) => state.taskErrors
  )

  let statusColor;
  let statusName;

  if (taskDraft.canceled) {
    statusColor = "#737373";
    statusName = "Canceled";
  } else if (taskDraft.completed) {
    statusColor = "#d9d9d9";
    statusName  = "Completed"
  } else {
    statusColor = "#99e17a";
    statusName = "Scheduled"
  }

  return (
    <Drawer 
      variant="temporary"
      anchor="right"  
      open={taskDetailsOpen} 
      onClose={onTaskClose}
      sx={{
        zIndex: (theme) => theme.zIndex.modal + 1,
        "& .MuiDrawer-paper": {
          width: TASK_DETAIL_WIDTH,
          boxSizing: "border-box",
        }}}
    >
      <Stack direction="column" spacing={4} sx={{ p: 5 }}>
        <Typography variant="h2">{taskDraft.templateName}</Typography>
        <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
          <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: statusColor}} />
          <Typography variant="h6">{statusName}</Typography>
        </Stack>
        <Stack direction="column" spacing={2}>
          <DetailRow label="Due Date" value={formatTimestamp(taskDraft.dueDate)} />
          <DetailRow label="Instructions" value={taskDraft.instructions} />
          <DetailRow label="Assigned To" value={taskDraft.userDisplayName ? task.userDisplayName : "Unassigned"} />
          <DetailRow label="Duration" value={formatDuration(taskDraft.duration)} />
          <DetailRowDate  
            label="Scheduled For"
            field="scheduledFor"
            value={taskDraft.scheduledFor ? dayjs(taskDraft.scheduledFor) : null}
            maxDate={taskDraft.dueDate}
            onValueChange={updateTaskDraft}
            // validateField={}
          />
        </Stack>
        <Stack>
          <Typography variant="body1">Notes</Typography>
          <TextField 
            id="notes"  
            variant="outlined" 
            multiline={true} 
            rows={5} 
            value={taskDraft.notes} 
            onChange={
              (event) => updateTaskDraft("notes", event.target.value)
            }
          />
        </Stack>
        
        {/* TODO:
          - create conditional logic to not display Complete if canceled == true */}
        <Typography
            variant="body2"
            sx={{
              color: "error.main",
              textAlign: "center",
              mt: 1,
            }}
        >
          {taskErrors.form}
        </Typography>
        <Button variant="contained" onClick={() => {taskDraft.completed=!taskDraft.completed; saveTask(); }}>
          {task.completed ? "Reopen Task" : "Complete Task"}
        </Button>
        <Button variant="text" onClick={() => {taskDraft.canceled=!taskDraft.canceled; saveTask(); }}>{taskDraft.canceled ? "Restore Task" : "Cancel Task"}</Button>
        
      </Stack>
    </Drawer>
  )
}