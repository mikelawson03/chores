import { Box, IconButton, Stack, Typography } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";

export default function AgendaHeader({ agendaDate, onNextAgendaDay, onPreviousAgendaDay }) {
    return(
        <Stack 
            direction="row" 
            sx={{
                alignItems: "center",
                
            }}
        >
            <IconButton onClick={onPreviousAgendaDay}>
                <ChevronLeftIcon fontSize="large"  />
            </IconButton>
            <Typography variant="h6">
                {agendaDate.format("MMMM DD YYYY")}
            </Typography>
            <IconButton onClick={onNextAgendaDay} >
                <ChevronRightIcon fontSize="large"  />
            </IconButton>
        </Stack>
    )
}