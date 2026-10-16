import { Stack, Typography } from "@mui/material";
import { clickableText } from "../../styles/typography";
import SortIndicator from "../SortIndicator";

export default function TableHeader({ sort, headings, onSortClick }) {
  return(
    <Stack direction="row" sx={{width:"100%", border: 1, borderColor: "divider", }}>
      {headings.map(heading  => (
        <Stack key={heading.field} direction="row" 
          onClick={() => {heading.sortable ? onSortClick(heading) : undefined}} sx={[heading.sortable && clickableText, {flex: 1, alignItems:"center", p: 2}]}>
          <Typography 
            variant="h6" 
            sx={{ borderColor: "divider"}}>
              {heading.displayName}
          </Typography>
          {heading.sortable && <SortIndicator sort={sort} heading={heading} />}
          </Stack>
      ))}
    </Stack>
  )
}