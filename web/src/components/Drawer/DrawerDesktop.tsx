import { Box, Drawer, makeStyles } from '@material-ui/core'
import React, { memo } from 'react'
import { DrawerType } from '@/@types/types'
import DrawerList from '@/components/Drawer/DrawerList'

const useStyles = makeStyles(({ palette: { primary }, spacing }) => {
  const drawerWidth = spacing(42);
  return {
    drawerRoot: {
      width: drawerWidth, // いみない
      flexShrink: 0,
    },
    drawerBg: {
      background: primary.main,
      color: primary.contrastText,
      width: drawerWidth,
    },
  };
});

const DrawerDesktop = () => {
  const classes = useStyles()

  return (
    <Drawer
      variant="permanent"
      anchor="left"
      classes={{
        root: classes.drawerRoot,
        paper: classes.drawerBg,
      }}
    >
      <Box p={3} pt={5}>
        <DrawerList type={DrawerType.DESKTOP1} />
      </Box>
    </Drawer>
  )
}

export default memo(DrawerDesktop);
