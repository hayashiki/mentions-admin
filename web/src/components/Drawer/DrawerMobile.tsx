import { Box, Drawer, Fab, makeStyles } from '@material-ui/core'
import React, { Fragment, memo, useState } from 'react'
import MenuIcon from "@material-ui/icons/Menu";
import { DrawerType } from '@/@types/types'
import DrawerList from '@/components/Drawer/DrawerList'

const useStyles = makeStyles(({ palette: { primary }, spacing, zIndex }) => {
  const unit = spacing(1.5);
  return {
    btn: {
      zIndex: zIndex.modal,
      position: "fixed",
      right: 15,
      bottom: 15,
    },
    drawerBg: {
      background: primary.main,
      color: primary.contrastText,
      borderTopLeftRadius: unit,
      borderTopRightRadius: unit,
      height: spacing(45),
    },
  };
});

const DrawerMobile = () => {
  const classes = useStyles()

  const [drawerState, setDrawerState] = useState<boolean>(false)
  const setDrawer = (bool: boolean) => () => setDrawerState(bool)

  return (
    <Fragment>
      <Fab
        className={classes.btn}
        size="medium"
        color="primary"
        onClick={setDrawer(true)}
      >
        <MenuIcon />
      </Fab>
      <Drawer
        open={drawerState}
        variant="temporary"
        onClose={setDrawer(false)}
        anchor="bottom"
      >
        <Box>
          <DrawerList
            type={DrawerType.MOBILE2}
            onTap={setDrawer(false)}
          />
        </Box>
      </Drawer>
    </Fragment>
  )
}

export default memo(DrawerMobile);
