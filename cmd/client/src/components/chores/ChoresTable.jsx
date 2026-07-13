import { Stack, Typography } from "@mui/material"
import dayjs from "dayjs";

export default function ChoresTables() {
    return (
        <Stack direction="column" sx={{flex: 1, width: "95%"}}>
            <Stack direction="row" sx={{width: "100%", border: 1, borderColor: "divider", justifyContent: "space-between"}}>
                <Typography variant="h6" sx={{flex: 1, borderColor: "divider", p: 2}}>Name</Typography>
                <Typography variant="h6" sx={{flex: 1, borderColor: "divider", p: 2}}>Frequency</Typography>
                <Typography variant="h6" sx={{flex: 1, borderColor: "divider", p: 2}}>Duration</Typography>
                <Typography variant="h6" sx={{flex: 1, borderColor: "divider", p: 2}}>Default Assignee</Typography>
                <Typography variant="h6" sx={{flex: 1, borderColor: "divider", p: 2}}>Last Updated</Typography>
            </Stack>
            <Stack direction="row" sx ={{ flex: 1, width: "100%", borderLeft: 1, borderRight:1, borderColor: "divider", justifyContent:"space-between"}}>
                <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>Wash dishes</Typography>
                <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>Daily</Typography>
                <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>30 mins</Typography>
                <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>N/A</Typography>
                <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>{dayjs().format("MMM DD YYYY - HH:MM")}</Typography>
            </Stack>
        </Stack>)
}