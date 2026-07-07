import { useState } from "react";
import { Button, Drawer, Box, Typography } from "@mui/material";

export default function DrawerTest() {
  const [open, setOpen] = useState(false);

  return (
    <>
      <Button onClick={() => setOpen(true)}>
        Open Drawer
      </Button>

      <Drawer
        anchor="right"
        open={open}
        onClose={() => setOpen(false)}
      >
        <Box sx={{ width: 400, p: 2 }}>
          <Typography variant="h5">
            Test Drawer
          </Typography>
        </Box>
      </Drawer>
    </>
  );
}