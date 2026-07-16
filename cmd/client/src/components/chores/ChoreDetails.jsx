import { Box, Button, Drawer, Stack, Typography } from "@mui/material";
import DetailRow from "../../details/DetailRow";
import DetailRowSelect from "../../details/DetailRowSelect";
import { CADENCES } from "../../constants/cadences";
import { USERS } from "../../constants/temp";
import DetailRowNumber from "../../details/DetailRowNumber";
import DetailRowLargeText from "../../details/DetailRowLargeText";
import DetailRowTitle from "../../details/DetailRowTitle";
import { formatTimestamp } from "../../utils/formatters";

export default function ChoreDetails({ open, chore, closeChoreDetails, onChoreDetailChange, saveChore, deleteChore}) {
    console.log(chore)
    const DRAWER_DETAIL_WIDTH=680
    // TODO: Verify behavior for nullable assignee ID
    const user_options=[
        { value: "", label: "Unassigned"},
        ...USERS.map(user => ({
            key: user.id,
            value: user.id,
            label: user.name,
        })),
    ]
    console.log(user_options)
    return (
        <Drawer variant="temporary" anchor="right" open={open} onClose={closeChoreDetails} sx={{
            "& .MuiDrawer-paper": {
                width: DRAWER_DETAIL_WIDTH,
                boxSizing: "border-box",
            }
        }}>
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
                <DetailRowTitle field="name" value={chore.name} onValueChange={onChoreDetailChange} />
                {/* <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
                    <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: statusColor}} />
                    <Typography variant="h6">{statusName}</Typography>
                </Stack> */}
                <Stack direction="column" spacing={2}>
                    <DetailRowSelect label="Frequency" field="cadence" value={chore.cadence} options={CADENCES} onValueChange={onChoreDetailChange}/>
                    <DetailRowSelect label="Assigned To" field="assignee" value={chore.assignee} options={user_options} onValueChange={onChoreDetailChange}/>
                    <DetailRowNumber label="Duration" field="duration" value={chore.duration} onValueChange={onChoreDetailChange} units="mins" />
                </Stack>
                <Stack>
                    <DetailRowLargeText label="Instructions" field="instructions" value={chore.instructions} onValueChange={onChoreDetailChange} />
                </Stack>
                <Stack direction="column" spacing={2}>
                    <Button variant="contained" onClick={() => {saveChore(chore); closeChoreDetails();}}>Save</Button>
                    <Button variant="text" onClick={() => {deleteChore(chore); closeChoreDetails();}}>Delete</Button>
                </Stack>
                <Stack direction="column" spacing={0.25}>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Last updated at {formatTimestamp(chore.updated_at)}</Typography>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Created at {formatTimestamp(chore.created_at)}</Typography>
                </Stack>
                <Typography variant="h6">Upcoming Tasks:</Typography>
                
            </Stack>
        </Drawer>
    )
    
}