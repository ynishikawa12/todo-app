import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { Login } from "../pages/Login";
import { SignUp } from "../pages/SignUp";
import { Profile } from "../pages/Profile";
import { Tag } from "../pages/Tag";
import { API_LOGIN, API_TAG, API_USER } from "../consts/consts";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path={API_LOGIN} element={<Login />} />
        <Route path={API_USER} element={<SignUp />} />
        <Route path={API_USER + "/:userId"} element={<Profile />} />
        <Route path={API_TAG} element={<Tag />} />
      </Routes>
    </Router>
  );
};

export default App;
