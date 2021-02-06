import React, { Fragment } from 'react';
import DrawerDesktop from '@/components/Drawer/DrawerDesktop'
import { Hidden } from '@material-ui/core'
import DrawerMobile from '@/components/Drawer/DrawerMobile'

const DrawerIndex = () => (
  <Fragment>
    <Hidden mdUp>
      <DrawerMobile />
    </Hidden>
    <Hidden smDown>
      <DrawerDesktop />
    </Hidden>
  </Fragment>
)

export default DrawerIndex;
