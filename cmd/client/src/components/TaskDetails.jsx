import { Box, Button, Drawer, Stack, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import DetailRow from "./details/DetailRow"
import { formatDuration, formatTimestamp } from "../utils/formatters";

export default function TaskDetails({task, open, finishTaskEditing}) {
  const TASK_DETAIL_WIDTH=680
  const [taskDraft, setTaskDraft] = useState(() => ({...task}))
  let statusColor;
  let statusName;

  useEffect(() => {
    setTaskDraft({...task});
    }, [task]
  );

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
          open={open} 
          onClose={() => finishTaskEditing(taskDraft)}
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
                <DetailRow label="Assigned To" value={taskDraft.userFirstName ? task.userFirstName : "Unassigned"} />
                <DetailRow label="Duration" value={formatDuration(taskDraft.duration)} />
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
                    (event) => setTaskDraft(
                      {...taskDraft,
                      notes: event.target.value}
                    )
                  }
                />
              </Stack>
              
               {/* TODO:
               - create conditional logic to not display Complete if canceled == true */}

              <Button variant="contained" onClick={() => {taskDraft.completed=!taskDraft.completed; finishTaskEditing(taskDraft); }}>
                {task.completed ? "Reopen Task" : "Complete Task"}
              </Button>
              <Button variant="text" onClick={() => {taskDraft.canceled=!taskDraft.canceled; finishTaskEditing(taskDraft); }}>{taskDraft.canceled ? "Restore Task" : "Cancel Task"}</Button>
            </Stack>
        </Drawer>
    )
}