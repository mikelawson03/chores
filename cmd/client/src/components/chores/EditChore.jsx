import { Button, CircularProgress, Drawer, Stack, Typography } from "@mui/material";
import DetailRowSelect from "../details/DetailRowSelect";
import { CADENCES } from "../../constants/cadences";
import DetailRowNumber from "../details/DetailRowNumber";
import DetailRowLargeText from "../details/DetailRowLargeText";
import DetailRowTitle from "../details/DetailRowTitle";
import { formatTimestamp } from "../../utils/formatters";
import { useMutation } from "@tanstack/react-query";
import { deleteChore } from "../../utils/choreHelpers";
import { queryClient } from "../../query/queryClient";



export default function EditChore({ open, chore, closeEditChore, editChoreMode, onChoreDetailChange, users, errors, validateChoreField, handleSave, isSaving}) {
    const DRAWER_DETAIL_WIDTH=680

    const userOptions=[
        { value: "unassigned", label: "Unassigned"},
        ...users.map(user => ({
            value: user.id,
            label: user.firstName,
        })),
    ]

    const deleteChoreMutation = useMutation({
        mutationFn: deleteChore,
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ["choreTemplates"],
            });
            closeEditChore();
        },
    });

    return (
        <Drawer variant="temporary" anchor="right" open={open} onClose={closeEditChore} sx={{
            "& .MuiDrawer-paper": {
                width: DRAWER_DETAIL_WIDTH,
                boxSizing: "border-box",
            }
        }}>
            <Stack direction="column" spacing={4} sx={{ p: 5 }}>
                <DetailRowTitle field="name" value={chore.name} error={errors.name} placeholder="Enter chore name..." onValueChange={onChoreDetailChange} required={true} validateChoreField={validateChoreField} />
                {/* <Stack direction="row" spacing={1} sx={{ alignItems: "center"}}>
                    <Box sx={{ width: 25, height: 25, borderRadius: "50%", backgroundColor: statusColor}} />
                    <Typography variant="h6">{statusName}</Typography>
                </Stack> */}
                <Stack direction="column" spacing={2}>
                    <DetailRowSelect label="Frequency" field="cadence" value={chore.cadence} error={errors.cadence} options={CADENCES} onValueChange={onChoreDetailChange} required={true} validateChoreField={validateChoreField}/>
                    <DetailRowSelect label="Assigned To" field="assignee" value={chore.assignee ?? "unassigned"} error={errors.assignee} options={userOptions} onValueChange={onChoreDetailChange} validateChoreField={validateChoreField} />
                    <DetailRowNumber label="Duration" field="duration" value={chore.duration} error={errors.duration} onValueChange={onChoreDetailChange} units="mins" required={true} validateChoreField={validateChoreField} />
                </Stack>
                <Stack>
                    <DetailRowLargeText label="Instructions" field="instructions" value={chore.instructions} onValueChange={onChoreDetailChange} />
                </Stack>
                <Stack direction="column" spacing={2}>
                    <Typography
                        variant="body2"
                        sx={{
                            color: "error.main",
                            textAlign: "center",
                            mt: 1,
                        }}
                    >
                        {errors.form}
                    </Typography>
                    <Button 
                        variant="contained" 
                        onClick={handleSave} 
                        disabled={isSaving}
                        startIcon={isSaving ? <CircularProgress size={16} /> : null}
                    >
                        {editChoreMode === "edit" ? "Save" : "Create"}
                    </Button>
                    {editChoreMode === "edit" && <Button variant="text" onClick={() => {deleteChoreMutation.mutate(chore.id);}}>Delete</Button>}
                    {editChoreMode === "create" && <Button variant="text" onClick={() => {closeEditChore();}}>Discard</Button>}
                </Stack>
                {editChoreMode === "edit" && <Stack direction="column" spacing={0.25} sx={{borderBottom: 1, borderColor: "divider", pb: 4}}>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Last updated at {formatTimestamp(chore.updatedAt)}</Typography>
                    <Typography variant="body2" sx={{fontStyle: 'italic' }}>Created at {formatTimestamp(chore.createdAt)}</Typography>
                </Stack >}
            </Stack>
        </Drawer>
    )
    
}