import React, { FC } from 'react'
import { List, ListItem, ListItemIcon, ListItemText, makeStyles, Typography } from '@material-ui/core'
import { DrawerType } from '@/@types/types'
import ProjectIcon from "@material-ui/icons/AssignmentTwoTone";
import IssuesIcon from "@material-ui/icons/BugReportTwoTone";

type Props = {
  onTap?(): void;
  type: DrawerType;
};

const items = [
  {
    text: "Projects",
    icon: ProjectIcon,
    path: "/d/project",
  },
  {
    text: "Issues",
    icon: IssuesIcon,
    path: "/d/issue",
  },
]

const useStyles = makeStyles(
  ({ shape: {borderRadius}, spacing }) => ({
    listItem: {
      borderRadius,
      marginTop: spacing(0.5)

    }

  })
)

const DrawerList: FC<Props> = ({ onTap, type}) => {
  const styles = useStyles();

  const list = items.map(({ text, icon: Icon, path}) => {
    return (
      <ListItem
        key={path}
        button
        className={styles.listItem}
      >
        <ListItemIcon>
          <Icon fontSize="small" />
        </ListItemIcon>
        <ListItemText
          primary={<Typography color="inherit">{text}</Typography>}
        />
      </ListItem>
      )
  });

  return (
    <List disablePadding>
      {list}
    </List>
  )
}

export default DrawerList
