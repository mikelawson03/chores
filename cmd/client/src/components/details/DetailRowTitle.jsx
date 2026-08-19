import { Stack, TextField, Typography } from "@mui/material"

export default function DetailRowTitle({field, placeholder, value, onValueChange, validateChoreField, error, required = false}) {
    return(
        <Stack>
            <TextField  
                variant="standard"
                value={value} 
                placeholder={placeholder + (required ? "*" : "")}
                onChange={(event) => onValueChange(field, event.target.value)}
                onBlur={() => validateChoreField(field, value)}
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
                    "&::placeholder": {
                        color: "text.secondary",
                        opacity: 1,
                    },
                }}
            />
            <Typography 
                variant="body2"
                sx={{
                    color: "error.main",
                }}
            >
                {error && error}
            </Typography>
        </Stack>
)
}