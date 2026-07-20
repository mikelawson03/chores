import { Stack } from "@mui/material";
import CalendarGridHeader from "./CalendarGridHeader";
import CalendarCells from "./CalendarCells";

export default function CalendarGrid({ currentDate, days, maxDayItems, onOverflowClick, openTaskDetails }) {
    return(
        <Stack 
            direction="column" 
            sx = {{ 
                flex: 1,
                minHeight: 0,
                border: 1,
                borderColor: "divider" ,
            }}
        >
            <CalendarGridHeader />
            <CalendarCells 
                currentDate={currentDate} 
                days={days} 
                maxDayItems={maxDayItems} 
                onOverflowClick={onOverflowClick}
                openTaskDetails={openTaskDetails}
                sx = {{ flex: 1 }} />
        </Stack>
    )
}