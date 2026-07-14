import { Box, Stack, Typography } from "@mui/material";
import SwapVertIcon from '@mui/icons-material/SwapVert';

export default function TableHeader({ sort, templates, headings }) {
  return(
    <Stack direction="row" sx={{flex: 1, width:"100%", border: 1, borderColor: "divider", }}>
      {headings.map(heading  => (
        <Stack key={heading}  direction="row" spacing={1} sx={{flex: 1, alignItems:"center", p: 1}}>
          <Typography 
            variant="h6" 
            sx={{ borderColor: "divider"}}>
              {heading}
          </Typography>
          <Box onClick={() => {sort(templates); console.log("clicked")}}><SwapVertIcon /></Box>
          </Stack>
      ))}
    </Stack>
  )
}