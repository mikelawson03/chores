import { Stack } from "@mui/material";
import ChoreCell from "./ChoreCell";
import { formatCadence, formatDuration, formatTimestamp } from "../../utils/formatters";



export default function ChoresTableRow({ choreTemplate, openEditChore }) {
  return (
    <Stack direction="row" onClick={() => {openEditChore(choreTemplate)}} sx ={{ flex: 1, width: "100%", borderLeft: 1, borderRight:1, borderColor: "divider", justifyContent:"space-between"}}>
      <ChoreCell item={choreTemplate.name} />
      <ChoreCell item={formatCadence(choreTemplate.cadence)} />
      <ChoreCell item={choreTemplate.assignee ? choreTemplate.assignee : "—"} />
      <ChoreCell item={formatDuration(choreTemplate.duration)} />
      <ChoreCell item={formatTimestamp(choreTemplate.updated_at)} />
  </Stack>
  )
}