import { Stack } from "@mui/material";
import ChoreCell from "./ChoreCell";
import StatusCell from "./StatusCell";
import { formatCadence, formatDuration, formatTimestamp } from "../../utils/formatters";




export default function ChoresTableRow({ choreTemplate, openEditChore, users }) {
  const getUserName = (userId) => {
        const user = users.find((hhUser) => hhUser.user.id === userId);
        return user?.displayName ?? "—";
    }

  return (
    <Stack direction="row" onClick={() => {openEditChore(choreTemplate)}} spacing={0.25} 
    sx ={{ 
      px: 1, 
      py: 2, 
      width: "100%", 
      borderLeft: 1, 
      borderRight:1, 
      borderColor: "divider", 
    }}>
      <ChoreCell item={choreTemplate.name} />
      <ChoreCell item={formatCadence(choreTemplate.cadence)} />
      <ChoreCell item={getUserName(choreTemplate.assignee)} />
      <StatusCell isActive={choreTemplate.isActive} />
      <ChoreCell item={formatDuration(choreTemplate.duration)} />
      <ChoreCell item={formatTimestamp(choreTemplate.updated_at)} />
  </Stack>
  )
}