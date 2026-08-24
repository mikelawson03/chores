import { Button, Dialog, DialogContent, DialogTitle, Stack } from "@mui/material"


export default function CloseDetailsAlert( {task, dialogOpen, setDialogOpen, onSave, onEdit, onDiscard} ) {
    
    return(
        <Dialog
            open={dialogOpen}
            onClose={() => setDialogOpen(false)}
            sx={{
            zIndex: (theme) => theme.zIndex.modal + 2,
            }}
            slotProps={{
            paper: {
                sx: {
                backgroundColor: "#fff",
                },
            },
            }}
        >
            <DialogTitle>Save changes?</DialogTitle>
            <DialogContent>You have unsaved changes to this task.</DialogContent>
            <Stack direction="row" spacing={3} sx={{ px: 2, pb: 2 }}>
                <Button onClick={() => onSave(task)}>Save</Button>
                <Button onClick={onEdit}>Keep Editing</Button>
                <Button onClick={onDiscard}>Discard Changes</Button>
                
            </Stack>
        </Dialog>
    )
}