import { Box, CircularProgress, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import EditChore from "../components/chores/EditChore";
import { EMPTY_CHORE_TEMPLATE } from "../constants/choreTemplate";
import { useState } from "react";
import { getUsers } from "../utils/userHelpers";
import { getChores } from "../utils/choreHelpers";
import { useQuery } from "@tanstack/react-query";

export default function Chores() {
  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [editChoreOpen, setEditChoreOpen] = useState(false);
  const [editChoreMode, setEditChoreMode] = useState(null);

  const usersQuery = useQuery({
    queryKey: ["users"],
    queryFn: () => getUsers(),
  })

  const choreTemplatesQuery = useQuery({
    queryKey: ["choreTemplates"],
    queryFn: () => getChores(),
  })

  const users = usersQuery.data ?? [];
  const choreTemplates = choreTemplatesQuery.data ?? [];

  function openEditChore(choreTemplate) {
    setEditedChoreTemplate({...choreTemplate});
    setEditChoreMode("edit");
    setEditChoreOpen(true);
  }

  function closeEditChore() {
    setEditChoreOpen(false);
  }

  function onChoreDetailChange(field, value) {
    setEditedChoreTemplate(previous => ({
      ...previous,
      [field]: value,
    }))
  }

  function editNewChore() {
    setEditedChoreTemplate({...EMPTY_CHORE_TEMPLATE});
    setEditChoreMode("create");
    setEditChoreOpen(true);
  }

   if (usersQuery.isPending || choreTemplatesQuery.isPending) {
      return( 
        <Box sx ={{
          width: "100%",
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center"
        }}>
          <CircularProgress />
        </Box>
    )
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
          <ChoresTable choreTemplates={choreTemplates} openEditChore={openEditChore} editNewChore={editNewChore} users={users}/>
          {editedChoreTemplate && 
            <EditChore 
              open={editChoreOpen} 
              chore={editedChoreTemplate} 
              closeEditChore={closeEditChore} 
              editChoreMode={editChoreMode}
              onChoreDetailChange={onChoreDetailChange} 
              users={users}
            />}
      </Stack>
)
}