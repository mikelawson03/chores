import { Box, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";

export default function Chores() {
  return (
      <Stack 
        spacing={4}
        sx={{
          flex: 1,
          alignItems: "center",
        }}
      >
        <Box 
          sx={{
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
          
          <ChoresTable />
        </Stack>
      </Stack>
)
}