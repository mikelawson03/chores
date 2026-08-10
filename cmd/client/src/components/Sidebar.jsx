import { Drawer, List, ListItem, ListItemButton, ListItemText } from "@mui/material";
import { useState } from "react";
import { clickableText } from "../styles/typography";
import { NavLink } from "react-router-dom";
import { useAuth } from "../auth/useAuth";


export default function Sidebar() {
  const DRAWER_WIDTH = 240;
  const MINI_DRAWER_WIDTH = 72;

  const { user } = useAuth();
  const isAdmin = user?.role ==="admin"

  const [open, setOpen] = useState(true)
  const navItems = [
    {
      text: "Dashboard",
      route: "/",
    },
    {
      text: "Chores",
      route: "/chores",
      adminOnly: true,
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
      adminOnly: true,
    },
  ]

  const visibleNavItems = navItems.filter(
    (item) => !item.adminOnly || isAdmin
  );
  

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
      {visibleNavItems.map(item => (
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