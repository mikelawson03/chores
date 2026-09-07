import { deepPurple, deepOrange, cyan } from "@mui/material/colors";


export const CADENCES = {
  "daily": { 
    label: "Daily",
    backgroundColor: cyan[50],
    color: cyan[800], 
    cadenceBadge: "D",
  },
  "weekly": { 
    label: "Weekly",
    backgroundColor: deepOrange[50],
    color: deepOrange[800],
    cadenceBadge: "W",
  },
  "monthly": { 
    label: "Monthly",
    backgroundColor: deepPurple[50],
    color: deepPurple[800], 
    cadenceBadge: "M",
  },
};