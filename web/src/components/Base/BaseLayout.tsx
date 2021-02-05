import React, { Suspense, FC } from "react";
import { Box, Grid, makeStyles } from '@material-ui/core'
import BaseLoader from '@/components/Base/BaseLoader'
import DrawerIndex from '@/components/Drawer/DrawerIndex'

const useStyles = makeStyles({
  container: {
    height: "100%"
  }
})

const BaseLayout: FC = ({ children }) => {
  const classes = useStyles()

  return (
    <Grid container className={classes.container}>
      <Grid item xs="auto">
        <DrawerIndex />
      </Grid>
      <Grid item xs>
        <Box
          p={{
            xs: 3,
            md: 5,
          }}
          flexGrow={1}
          component="main"
        >
          {children}
        </Box>
      </Grid>
    </Grid>
  )
}

export default BaseLayout;

// <Grid item xs>
//   <Box
//     p={{
//       xs: 3,
//       md: 5,
//     }}
//     flexGrow={1}
//     component="main"
//   >
//     <Suspense fallback={ <BaseLoader /> }>{children}</Suspense>
//   </Box>
// </Grid>
