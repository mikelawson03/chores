import { Box, Stack } from "@mui/material";
import { Typography } from "@mui/material";

export default function CalendarGridHeader() {
    const weekdays = ["Mon", "Tues", "Wed", "Thurs", "Fri", "Sat", "Sun"]
    return(
        <Stack 
            direction="row"
            >
            {weekdays.map( day => (
                <Box 
                key={day}
                sx={{ 
                    flex: 1,
                    borderBottom:1, 
                    borderRight: day === "Sun" ? 0 : 1, 
                    borderColor:"divider",
                    py: 2,
                    display: "flex",
                    alignItems: "center",
                    flexDirection: "column"
                    }}>
                        <Typography variant="h6">{day}</Typography>
                </Box>
            ))}
        </Stack>
    )
}