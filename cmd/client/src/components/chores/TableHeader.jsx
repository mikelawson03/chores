import { Box, Stack, Typography } from "@mui/material";

export default function TableHeader({ headings }) {
  console.log(headings)
  return(
    <Stack direction="row" sx={{width: "100%", border: 1, borderColor: "divider", justifyContent: "space-between"}}>
      {headings.map(heading  => (
        <Typography 
        key={heading}  
        variant="h6" 
          sx={{flex: 1, borderColor: "divider", p: 2}}>{heading}</Typography>
      ))}
    </Stack>
  )
}