import { Button, Stack, TextField } from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";


export default function TableToolbar(){
  return(
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
        <TextField size="small" />
        <FilterListIcon fontSize="large" />
      </Stack>
      <Button variant="contained">New Chore Template</Button>
    </Stack>
  )
}