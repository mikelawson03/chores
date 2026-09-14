import { Typography } from "@mui/material"

export default function ChoreCell({ item }) {
  return(
    <Typography 
      sx={{
        flex: 1, 
        borderBottom: 1, 
        borderColor: "divider", 
        p: 2,
        overflow: "hidden",
        textOverflow: "ellipsis",
        whiteSpace: "nowrap"
      }}>
        {item}
    </Typography>
  )
}