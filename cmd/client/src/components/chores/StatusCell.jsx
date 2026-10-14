import { Stack, Typography } from "@mui/material";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import RemoveCircleOutlinedIcon from "@mui/icons-material/RemoveCircleOutlined";

export default function StatusCell({ isActive }) {
  return(
    <Stack direction = "row" alignItems="center" spacing={0.75} sx={{flex: 1, }}>
        {isActive ? (
            <CheckCircleIcon color="success" fontSize="small" />
        ) : (
            <RemoveCircleOutlinedIcon color="disabled" fontSize="small" />
        )
    }
        <Typography 
        sx={{
            flex: 1, 
            borderBottom: 1, 
            borderColor: "divider", 
            overflow: "hidden",
            textOverflow: "ellipsis",
            whiteSpace: "nowrap"
        }}>
            {isActive ? "Active" : "Inactive" }
        </Typography>
    </Stack>
  )
}