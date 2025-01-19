import { AppBar, Menu, MenuItem, Toolbar, Typography, IconButton, Box, Container } from "@mui/material";
import AccountCircle from '@mui/icons-material/AccountCircle';
import { Link, useNavigate } from "react-router-dom";
import { EDIT_USER_URL } from "../consts/constants";

interface TodoAppBarProps {
  id: string;
}

const TodoAppBar = (props :TodoAppBarProps) => {
  const navigate = useNavigate();
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          Todo
        </Typography>
        <IconButton
          size="large"
          aria-label="account of current user"
          aria-controls="menu-appbar"
          aria-haspopup="true"
          onClick={() => navigate(EDIT_USER_URL + '/' + props.id)}
          color="inherit"
        >
          <AccountCircle />
        </IconButton>
      </Toolbar>
    </AppBar>
  );
}

export default TodoAppBar;