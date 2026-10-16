import { Stack, Typography } from "@mui/material";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import RemoveCircleOutlinedIcon from "@mui/icons-material/RemoveCircleOutlined";

export default function StatusCell({ isActive }) {
  return(
    <Stack direction = "row" spacing={0.75} sx={{flex: 1, borderBottom: 1, p: 2,
            borderColor: "divider", alignItems: "center"}}>
        {isActive ? (
            <CheckCircleIcon color="success" fontSize="small" />
        ) : (
            <RemoveCircleOutlinedIcon color="disabled" fontSize="small" />
        )
    }
        <Typography 
        sx={{
            flex: 1, 
            overflow: "hidden",
            textOverflow: "ellipsis",
            whiteSpace: "nowrap"
        }}>
            {isActive ? "Active" : "Inactive" }
        </Typography>
    </Stack>
  )
}