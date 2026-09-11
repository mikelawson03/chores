import { deepPurple, deepOrange, cyan } from "@mui/material/colors";


export const CADENCES = {
  "daily": { 
    label: "Daily",
    backgroundColor: cyan[50],
    outlineColor: cyan[200],
    color: cyan[800], 
    cadenceBadge: "D",
  },
  "weekly": { 
    label: "Weekly",
    backgroundColor: deepOrange[50],
    outlineColor: deepOrange[200],
    color: deepOrange[800],
    cadenceBadge: "W",
  },
  "monthly": { 
    label: "Monthly",
    backgroundColor: deepPurple[50],
    outlineColor: deepPurple[200],
    color: deepPurple[800], 
    cadenceBadge: "M",
  },
};