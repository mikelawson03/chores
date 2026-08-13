import { Dialog, DialogContent, DialogTitle } from "@mui/material";
import AgendaHeader from "./AgendaHeader";
import AgendaList from "./AgendaList";
import { getScheduledTasksForDay } from "../../utils/taskHelpers";


export default function DailyAgenda({ activeTasks, dailyAgendaOpen, onDailyAgendaClose, agendaDate, onNextAgendaDay, onPreviousAgendaDay, toggleTaskComplete }) {
    return(
        <Dialog
            maxWidth="sm"
            fullWidth={true}
            open={dailyAgendaOpen} 
            onClose={onDailyAgendaClose}
            slotProps={{
                paper: {
                    sx: {
                        height: "80%",
                        display: "flex",
                        flexDirection: "column",
                    }
                }
            }}
        >
            <DialogTitle 
                sx={{
                    
                    width: "100%",
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "center",
                    justifyContent: "center",
                    borderBottom: 1,
                    borderColor: "divider",
                    mb: 3
                }}
            >
                <AgendaHeader
                    agendaDate={agendaDate}
                    onNextAgendaDay={onNextAgendaDay}
                    onPreviousAgendaDay={onPreviousAgendaDay}
                 />                
            </DialogTitle>
            <DialogContent>
                <AgendaList 
                    tasks={getScheduledTasksForDay(agendaDate, activeTasks)}
                    toggleTaskComplete={toggleTaskComplete}
                />
            </DialogContent>
        </Dialog>
    )
}