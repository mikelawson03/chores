import { Box, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import ChoreDetails from "../components/chores/ChoreDetails";
import { useState } from "react";

export default function Chores() {
  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [choreDetailsOpen, setChoreDetailsOpen] = useState(false);
  

  const [choreTemplates, setChoreTemplates]=useState([
      {
        id: 1,
        name: "Make dinner",
        cadence: "daily",
        assignee: "",
        duration: 60,
        created_at: "2026-11-10T05:26:00",
        updated_at: "2026-11-10T05:26:00"
      },
      {
        id: 2,
        name: "Wash dishes",
        cadence: "daily",
        assignee: "",
        duration: 75,
        created_at: "2026-11-13T17:26:00",
        updated_at: "2026-07-13T17:26:00"
      },
      {
        id: 3,
        name: "Clean fish tank",
        cadence: "weekly",
        assignee: "Mike",
        duration: 20,
        created_at: "2026-06-05T17:26:00",
        updated_at: "2026-07-05T17:26:00"
      },
      {
        id: 4,
        name: "Mow yard",
        cadence: "weekly",
        assignee: "",
        duration: 90,
        created_at: "2026-07-14T16:11:00",
        updated_at: "2026-07-14T16:11:00"
      }
    ])

  function openChoreDetails(choreTemplate) {
    setEditedChoreTemplate({...choreTemplate});
    setChoreDetailsOpen(true);
  }

  return (
      <Stack spacing={4} sx={{
          flex: 1,
          alignItems: "center",
        }}
      >
        <Box sx={{
            width: "100%", 
            pb: 3, 
            borderBottom: 1, 
            borderColor: "divider", 
            display:"flex", 
            flexDirection: "column", 
            alignItems: "center"
          }}
        >
          <PageHeader title="Chore Management" />
        </Box>
        <Stack direction="column" sx={{flex: 1, width: "95%"}}>
          <ChoresTable choreTemplates={choreTemplates} openChoreDetails={openChoreDetails}/>
        </Stack>
          {editedChoreTemplate && <ChoreDetails open={choreDetailsOpen} chore={editedChoreTemplate} />}
      </Stack>
)
}