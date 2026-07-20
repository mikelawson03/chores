import { Box, Button, Drawer, Stack, Typography } from "@mui/material";
import DetailRow from "../details/DetailRow";
import DetailRowSelect from "../details/DetailRowSelect";
import { CADENCES } from "../../constants/cadences";
import { USERS } from "../../config/dev";
import DetailRowNumber from "../details/DetailRowNumber";
import DetailRowLargeText from "../details/DetailRowLargeText";
import DetailRowTitle from "../details/DetailRowTitle";
import { formatTimestamp } from "../../utils/formatters";

export default function EditChore({ open, chore, closeEditChore, createNewChore, editChoreMode, onChoreDetailChange, saveChore, deleteChore}) {
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
    return (
        <Drawer variant="temporary" anchor="right" open={open} onClose={closeEditChore} sx={{
            "& .MuiDrawer-paper": {
                width: DRAWER_DETAIL_WIDTH,
                boxSizing: "border-box",
            }
        }}>
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
                <DetailRowTitle field="name" value={chore.name} placeholder="Enter chore name..." onValueChange={onChoreDetailChange} />
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
                    <Button variant="contained" onClick={() => {editChoreMode === "edit" ? saveChore() : createNewChore(); closeEditChore();}}>{editChoreMode === "edit" ? "Save" : "Create"}</Button>
                    {/* <Button variant="text" onClick={() => {editChoreMode === "edit" ? deleteChore(chore) : undefined; closeEditChore();}}>{editChoreMode === "edit" ? "Delete" : "Discard"}</Button> */}
                    {editChoreMode === "edit" && <Button variant="text" onClick={() => {deleteChore(); closeEditChore();}}>Delete</Button>}
                    {editChoreMode === "create" && <Button variant="text" onClick={() => {closeEditChore();}}>Discard</Button>}
                </Stack>
                {editChoreMode === "edit" && <Stack direction="column" spacing={0.25} sx={{borderBottom: 1, borderColor: "divider", pb: 4}}>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Last updated at {formatTimestamp(chore.updated_at)}</Typography>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Created at {formatTimestamp(chore.created_at)}</Typography>
                </Stack >}
                {editChoreMode === "edit" && <Stack spacing={1}>
                    <Typography variant="h5">Upcoming Tasks:</Typography>
                    <Stack sx={{border: 1, borderColor: "divider", minHeight: 200, width: "100%"}}>
                        <Stack direction="row" sx={{borderBottom: 2, borderColor: "divider", p: 1}}>
                            <Typography variant="h6" sx={{width: "50%"}}>Date</Typography>
                            <Typography variant="h6">Assigned to</Typography>
                        </Stack>
                        <Stack direction="row" sx={{borderBottom: 2, borderColor: "divider", p: 1}}>
                            <Typography sx={{width: "50%"}}>Jul 20, 2026</Typography>
                            <Typography >Mike</Typography>
                        </Stack>
                        <Stack direction="row" sx={{borderBottom: 2, borderColor: "divider", p: 1}}>
                            <Typography sx={{width: "50%"}}>Jul 20, 2026</Typography>
                            <Typography >Mike</Typography>
                        </Stack>
                        <Stack direction="row" sx={{borderBottom: 2, borderColor: "divider", p: 1}}>
                            <Typography sx={{width: "50%"}}>Jul 20, 2026</Typography>
                            <Typography >Mike</Typography>
                        </Stack>
                        <Stack direction="row" sx={{borderBottom: 2, borderColor: "divider", p: 1}}>
                            <Typography sx={{width: "50%"}}>Jul 20, 2026</Typography>
                            <Typography >Mike</Typography>
                        </Stack>
                        <Stack direction="row" sx={{p: 1}}>
                            <Typography sx={{width: "50%"}}>Jul 20, 2026</Typography>
                            <Typography >Mike</Typography>
                        </Stack>
                    </Stack>
                </Stack>}
            </Stack>
        </Drawer>
    )
    
}