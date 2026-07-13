import { Stack, Typography } from "@mui/material"
import TableHeader from "./TableHeader";
import TableToolbar from "./TableToolbar";

import dayjs from "dayjs";

export default function ChoresTable() {
  const tableHeadings=["Name", "Frequency", "Duration", "Default Assignee", "Last Updated"]
  return (
    <Stack direction="column" sx={{flex: 1, width: "100%"}}>
      <TableToolbar />
      <TableHeader headings={tableHeadings} />
      <Stack direction="row" sx ={{ flex: 1, width: "100%", borderLeft: 1, borderRight:1, borderColor: "divider", justifyContent:"space-between"}}>
          <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>Wash dishes</Typography>
          <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>Daily</Typography>
          <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>30 mins</Typography>
          <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>—</Typography>
          <Typography sx={{flex: 1, borderBottom: 1, borderColor: "divider", p: 2}}>{dayjs().format("MMM DD YYYY @ HH:MM")}</Typography>
      </Stack>
    </Stack>
  )
}