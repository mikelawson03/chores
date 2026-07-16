import { Box, Button, Drawer, Stack, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import DetailRow from "../details/DetailRow"

export default function TaskDetails({task, open, closeTaskDetails, toggleTaskComplete, toggleTaskCancel}) {
    const TASK_DETAIL_WIDTH=680
    let statusColor;
    let statusName;

    if (task.canceled) {
      statusColor = "#737373";
      statusName = "Canceled";
    } else if (task.completed) {
      statusColor = "#d9d9d9";
      statusName  = "Completed"
    } else {
      statusColor = "#99e17a";
      statusName = "Scheduled"
    }
    
    const [notes, setNotes] = useState(task.notes)

    useEffect(() => {
      setNotes(task.notes);
    }, [task]);

    return (
        <Drawer 
          variant="temporary"
          anchor="right"  
          open={open} 
          onClose={() => closeTaskDetails(task.id, notes)}
          sx={{
            zIndex: (theme) => theme.zIndex.modal + 1,
            "& .MuiDrawer-paper": {
                width: TASK_DETAIL_WIDTH,
                boxSizing: "border-box",
            }}}
            >
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
              <Typography variant="h2">{task.title}</Typography>
              <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
                <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: statusColor}} />
                <Typography variant="h6">{statusName}</Typography>
              </Stack>
              <Stack direction="column" spacing={2}>
                <DetailRow label="Due Date" value={task.dueDate} />
                <DetailRow label="Instructions" value={task.instructions} />
                <DetailRow label="Assigned To" value={task.assignee} />
                <DetailRow label="Duration" value={task.duration} />
              </Stack>
              <Stack>
                <Typography variant="body1">Notes</Typography>
                <TextField id="notes"  variant="outlined" multiline={true} rows={5} value={notes} onChange={(event) => setNotes(event.target.value)}  />
              </Stack>
              
               {/* TODO:
               - create conditional logic to not display Complete if canceled == true */}

              <Button variant="contained" onClick={() => {toggleTaskComplete(task); closeTaskDetails(task.id, notes); }}>
                {task.completed ? "Reopen Task" : "Complete Task"}
              </Button>
              <Button variant="text" onClick={() => {toggleTaskCancel(task); closeTaskDetails(task.id, notes); }}>{task.canceled ? "Restore Task" : "Cancel Task"}</Button>
            </Stack>
        </Drawer>
    )
}