import { Typography } from "@mui/material";

export default function PageHeader({ title }) {
  return (
    <Typography variant="h4" sx={{ marginTop: 8 }}>
      {title}
    </Typography>
  );
}