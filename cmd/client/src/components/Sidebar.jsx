import { Button, Drawer, List, ListItem, ListItemButton, ListItemText, Stack } from "@mui/material";
import { useState } from "react";
import { clickableText } from "../styles/typography";
import { NavLink, useNavigate } from "react-router-dom";
import { useAuth } from "../auth/useAuth";
import { TbLayoutSidebarLeftCollapse, TbLayoutSidebarRightCollapse } from "react-icons/tb";
import DashboardIcon from '@mui/icons-material/Dashboard';
import CleaningServicesIcon from '@mui/icons-material/CleaningServices';
import CalendarViewWeekIcon from '@mui/icons-material/CalendarViewWeek';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import SettingsIcon from '@mui/icons-material/Settings';
import LogoutIcon from '@mui/icons-material/Logout';


export default function Sidebar() {
  const DRAWER_WIDTH = 240;
  const MINI_DRAWER_WIDTH = 90;

  const { user, logoutUser } = useAuth();
  const isAdmin = user?.role ==="admin"
  const navigate = useNavigate();

  const [sidebarExpanded, setSidebarExpanded] = useState(true)

  const onExpandClick = () => {
    setSidebarExpanded(current => !current);
    
  }

  const handleLogout = async () => {
    try {
      await logoutUser();
      navigate("/")
    } catch (err) {
      console.log(err)
    }
  }

  const navItems = [
    {
      text: "Dashboard",
      route: "/",
      icon: <DashboardIcon />
    },
    {
      text: "Chores",
      route: "/chores",
      adminOnly: true,
      icon: <CleaningServicesIcon />
    },
    {
      text: "Planner",
      route: "/planner",
      icon: <CalendarViewWeekIcon />
    },
    {
      text: "Calendar",
      route: "/calendar",
      icon: <CalendarMonthIcon />
    },
    {
      text: "Settings",
      route: "/admin",
      adminOnly: true,
      icon: <SettingsIcon />
    }
  ]

  const visibleNavItems = navItems.filter(
    (item) => !item.adminOnly || isAdmin
  );
  

  return(
    <Drawer 
      open={true} 
      variant="permanent"
      sx={{
        width: sidebarExpanded ? DRAWER_WIDTH : MINI_DRAWER_WIDTH,
        transition: "width 0.25s ease",
        "& .MuiDrawer-paper": {
          width: sidebarExpanded ? DRAWER_WIDTH : MINI_DRAWER_WIDTH,
          boxSizing: "border-box",
          transition: "width 0.25s ease",
        }
      }}
    >
    
    <Stack sx={{
      alignItems: sidebarExpanded ? "end" : "center", 
      pt: 2, 
      pb: 0
    }}>
      <Button onClick={onExpandClick}>
        {sidebarExpanded && <TbLayoutSidebarLeftCollapse size="1.5em" color="#424242" />}
        {!sidebarExpanded && <TbLayoutSidebarRightCollapse size="1.5em" color="#424242" />}
      </Button>
    </Stack>
    <List>
      {visibleNavItems.map(item => (
        <ListItem sx={{pb: 5}} key={item.text}>
          <ListItemButton component={NavLink} to={item.route}>
            {item.icon}
            <ListItemText 
              sx={[
                clickableText, 
                {
                  pl: 1,
                  opacity: sidebarExpanded ? 1 : 0,
                  width: sidebarExpanded ? "auto" : 0,
                  overflow: "hidden",
                  transition: "opacity 0.2s ease, width 0.25s ease"
                }
              ]}
              primary={item.text} 
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
      <ListItem >
        <ListItemButton onClick={handleLogout}>
          <LogoutIcon />
          {sidebarExpanded && <ListItemText
            sx={[clickableText, {pl: 1}]}
            primary="Logout"
            slotProps={{
              primary: {
                variant: "h6",
              }
            }}
          />}
        </ListItemButton>
      </ListItem>  
      
    </List>     
    </Drawer>)
}