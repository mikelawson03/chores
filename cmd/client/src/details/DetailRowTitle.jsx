import { TextField } from "@mui/material"

export default function DetailRowTitle({field, value, onValueChange}) {
    return(
        <TextField  
            variant="standard"
            value={value} 
            onChange={(event) => onValueChange(field, event.target.value)} 
            slotProps={{
                input: {
                    disableUnderline: true,
                },
            }}
            sx={{
                borderBottom: 1,
                borderColor: "divider",
                "&:hover": {
                    borderColor: "text.secondary"
                },
                "&:focus-within": {
                    borderColor: "text.primary"
                },
                "& .MuiInputBase-input": {
                    typography: "h2",
                    fontWeight: "inherit",
                    p: 0,
                },
            }}
        />
)
}