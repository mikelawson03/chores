import { Drawer, List, ListItem, ListItemButton, ListItemText } from "@mui/material";
import { useState } from "react";
import { clickableText } from "../styles/typography";
import { NavLink } from "react-router-dom";


export default function Sidebar() {
  const DRAWER_WIDTH = 240;
  const MINI_DRAWER_WIDTH = 72;

  const [open, setOpen] = useState(true)
  const navItems = [
    {
      text: "Dashboard",
      route: "/",
    },
    {
      text: "Chores",
      route: "/chores",
    },
    {
      text: "Planner",
      route: "/planner",
    },
    {
      text: "Calendar",
      route: "/calendar",
    },
    {
      text: "Settings",
      route: "/admin",
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
          <ListItemButton component={NavLink} to={item.route}>
            <ListItemText 
              primary={item.text} 
              // sx={clickableText}
              slotProps={{
                primary: {
                  variant: "h6",
                }
              }}
            />
          </ListItemButton>
        </ListItem>
      ))
      }  
    </List>     
    </Drawer>)
}