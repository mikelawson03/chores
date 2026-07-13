import { Box, Button, IconButton, Stack, TextField, Typography } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import FilterListIcon from "@mui/icons-material/FilterList";
import dayjs from "dayjs";

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
          <Stack 
          direction="row" 
          sx={{
            width: "100%", 
            justifyContent: "space-between",
            pb: 1,
          }}
        >
          <Stack 
            direction="row" 
            spacing={2}
          >
            <Box>
              <TextField size="small"  />
            </Box>
            <Box>
              <FilterListIcon fontSize="large" />
            </Box>
          </Stack>
          <Box>
            <Button variant="contained">New Chore Template</Button>
          </Box>
        </Stack>
          <ChoresTable />
        </Stack>
      </Stack>
)
}