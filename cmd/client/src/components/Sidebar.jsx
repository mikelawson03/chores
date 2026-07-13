import { Drawer, List, ListItem, ListItemText } from "@mui/material";
import { useState } from "react";
import { clickableText } from "../styles/typography";


export default function Sidebar() {
  const DRAWER_WIDTH = 240;
  const MINI_DRAWER_WIDTH = 72;

  const [open, setOpen] = useState(true)
  const navItems = [
    {
      text: "Dashboard"
    },
    {
      text: "Chores"
    },
    {
      text: "Planner"
    },
    {
      text: "Calendar"
    },
    {
      text: "Settings"
    },
  ]

  return(
    <Drawer 
      open={open} 
      variant="permanent"
      sx={{
        width:DRAWER_WIDTH,
        "& .MuiDrawer-paper": {
          width: DRAWER_WIDTH,
          boxSizing: "border-box",
        }
      }}
    >
    <List>
      {navItems.map(item => (
        <ListItem sx={{p: 3}} key={item.text}>
          <ListItemText 
          primary={item.text} 
          sx={clickableText}
          slotProps={{
            primary: {
              variant: "h6",
          }
        }}
          />
        </ListItem>
      ))
      }  
    </List>     
    </Drawer>)
}