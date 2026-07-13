import { IconButton, Stack, Typography } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";

export default function CalendarToolbar({ currentDate, onPreviousMonth, onNextMonth }) {
    return (
        <Stack 
            direction="row" 
            spacing={2} 
            sx ={{justifyContent: "center", 
                alignItems: "center",
                pb:3
            }} 
        >
            <IconButton onClick={onPreviousMonth} >
                <ChevronLeftIcon fontSize="large" />
            </IconButton>
            <Typography variant="h5">
                {currentDate.format("MMMM YYYY")}
            </Typography>
            <IconButton onClick={onNextMonth}>
                <ChevronRightIcon fontSize="large" />
            </IconButton>
        </Stack>
    )
}