import { Stack } from "@mui/material";
import ChoreCell from "./ChoreCell";
import { formatCadence } from "../../utils/stringFormatting";


export default function ChoresTableRow({ choreTemplate }) {
  return (
    <Stack direction="row" sx ={{ flex: 1, width: "100%", borderLeft: 1, borderRight:1, borderColor: "divider", justifyContent:"space-between"}}>
      <ChoreCell item={choreTemplate.name} />
      <ChoreCell item={formatCadence(choreTemplate.cadence)} />
      <ChoreCell item={choreTemplate.assignee ? choreTemplate.assignee : "—"} />
      <ChoreCell item={choreTemplate.duration} />
      <ChoreCell item={choreTemplate.updated_at} />
  </Stack>
  )
}