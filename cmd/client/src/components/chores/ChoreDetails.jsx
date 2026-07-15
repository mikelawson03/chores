import { Box, Drawer, Stack, Typography } from "@mui/material";
import DetailRow from "../DetailRow";
import DetailRowSelect from "../DetailRowSelect";
import { CADENCES } from "../../constants/cadences";

export default function ChoreDetails({ open, chore }) {
    const DRAWER_DETAIL_WIDTH=680
    
    return (
        <Drawer variant="temporary" anchor="right" open={open} sx={{
            "& .MuiDrawer-paper": {
                width: DRAWER_DETAIL_WIDTH,
                boxSizing: "border-box",
            }
        }}>
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
        <Typography variant="h2">{chore.name}</Typography>
        {/* <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
            <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: statusColor}} />
            <Typography variant="h6">{statusName}</Typography>
        </Stack> */}
        <Stack direction="column" spacing={2}>
            <DetailRowSelect label="Frequency" value={chore.cadence} options={CADENCES} />
            <DetailRow label="Instructions" value={chore.instructions} />
            <DetailRow label="Assigned To" value={chore.assignee} />
            <DetailRow label="Duration" value={chore.duration} />
        </Stack>
        {/* <Stack>
            <Typography variant="body1">Notes</Typography>
            <TextField id="notes"  variant="outlined" multiline={true} rows={5} value={notes} onChange={(event) => setNotes(event.target.value)}  />
        </Stack> */}
    </Stack>
        </Drawer>
    )
    
}