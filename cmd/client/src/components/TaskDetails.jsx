import { Box, Button, Drawer, Stack, TextField, Typography } from "@mui/material";
import DetailRow from "../components/DetailRow"

export default function TaskDetails() {
    const TASK_DETAIL_WIDTH=680
    return (
        <Drawer anchor="right"  open={true} sx={{
            width: TASK_DETAIL_WIDTH,
            "& .MuiDrawer-paper": {
                width: TASK_DETAIL_WIDTH,
                boxSizing: "border-box",
            }}}>
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
              <Typography variant="h2">Feed dog</Typography>
              <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
                <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: "#99e17a"}} />
                <Typography variant="h6">Scheduled</Typography>
              </Stack>
              <Stack direction="column" spacing={2}>
                <DetailRow label="Due Date" value="July 6, 2026" />
                <DetailRow label="Instructions" value="Feed Ruby 1 cup of food every morning" />
                <DetailRow label="Assigned To" value="Mike" />
                <DetailRow label="Duration" value="10 minutes" />
              </Stack>
              <Stack>
                <Typography variant="body1">Notes</Typography>
                <TextField id="notes"  variant="outlined" multiline={true} rows={5} />
              </Stack>
              <Button variant="contained">Complete Task</Button>
              <Button variant="text">Cancel Task</Button>
              
            </Stack>
        </Drawer>
    )
}