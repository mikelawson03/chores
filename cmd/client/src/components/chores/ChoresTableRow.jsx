import { Stack } from "@mui/material";
import ChoreCell from "./ChoreCell";
import { formatCadence, formatDuration, formatTimestamp } from "../../utils/formatters";



export default function ChoresTableRow({ choreTemplate, openEditChore, users }) {
  const getUserName = (userId) => {
        const user = users.find((user) => user.id === userId);
        return user?.firstName ?? "—";
    }

  return (
    <Stack direction="row" onClick={() => {openEditChore(choreTemplate)}} sx ={{ width: "100%", borderLeft: 1, borderRight:1, borderColor: "divider", justifyContent:"space-between"}}>
      <ChoreCell item={choreTemplate.name} />
      <ChoreCell item={formatCadence(choreTemplate.cadence)} />
      <ChoreCell item={getUserName(choreTemplate.assignee)} />
      <ChoreCell item={formatDuration(choreTemplate.duration)} />
      <ChoreCell item={formatTimestamp(choreTemplate.updated_at)} />
  </Stack>
  )
}