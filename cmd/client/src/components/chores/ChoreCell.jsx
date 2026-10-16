import { Typography } from "@mui/material"

export default function ChoreCell({ item }) {
  return(
    <Typography 
      sx={{
        flex: 1, 
        borderBottom: 1, 
        p: 2,
        borderColor: "divider", 
        overflow: "hidden",
        textOverflow: "ellipsis",
        whiteSpace: "nowrap"
      }}>
        {item}
    </Typography>
  )
}