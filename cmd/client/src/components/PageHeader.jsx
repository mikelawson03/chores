import { Typography } from "@mui/material";

export default function PageHeader({ title }) {
  return (
    
      <Typography variant="h4" sx={{ marginTop: 3 }}>
        {title}
      </Typography>
    
  );
}