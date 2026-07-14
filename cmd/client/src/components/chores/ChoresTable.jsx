import { Stack, Typography } from "@mui/material"
import TableHeader from "./TableHeader";
import TableToolbar from "./TableToolbar";
import ChoresTableRow from "./ChoresTableRow";
import { useState } from "react";

export default function ChoresTable() {
  const tableHeadings=["Name", "Frequency", "Default Assignee", "Duration", "Last Updated"]
  
  const [choreTemplates, setChoreTemplates]=useState([
    {
      id: 1,
      name: "Make dinner",
      cadence: "daily",
      assignee: "",
      duration: "60 minutes",
      created_at: "2026-11-10T05:26:00",
      updated_at: "2026-11-10T05:26:00"
    },
    {
      id: 2,
      name: "Wash dishes",
      cadence: "daily",
      assignee: "",
      duration: "60 minutes",
      created_at: "2026-11-10T17:26:00",
      updated_at: "2026-11-10T17:26:00"
    },
    {
      id: 3,
      name: "Clean fish tank",
      cadence: "weekly",
      assignee: "Mike",
      duration: "30 minutes",
      created_at: "2026-06-05T17:26:00",
      updated_at: "2026-07-05T17:26:00"
    }
  ])

  function sortByName(templates) { 
    setChoreTemplates(templates.sort((a, b) => {
      const nameA = a.name.toUpperCase()
      const nameB = b.name.toUpperCase()

      if (nameA > nameB) {
        return -1;
      }

      if (nameA > nameB) {
        return 1;
      }

      return 0;
  }))}

  return (
    <Stack 
      direction="column" 
      sx={{
        flex: 1, 
        width: "100%"
      }
    }>
      <TableToolbar />
      <TableHeader sort={sortByName} templates={choreTemplates} headings={tableHeadings} />
      {choreTemplates.map( chore => (
      <ChoresTableRow key={chore.id} choreTemplate={chore}/>))}
    </Stack>
  )
}