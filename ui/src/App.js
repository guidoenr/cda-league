import React, { useState, useEffect } from 'react';
import PlayersList from './modules/player/PlayersList'
import PlayersTable from "./modules/player/PlayersTable";
import {BrowserRouter as Router , Route, Routes} from "react-router-dom";
import './App.css'

import MyNavbar from "./modules/header/Navbar";
import PlayersTableRank from "./modules/player/PlayersTable";

function App(){

    return (
        <Router>
        <div className="App" >
            <MyNavbar />
            <Routes >
                <Route path="/" exact element={<PlayersTableRank />} />
                <Route path="players" element={<PlayersList />} />
                <Route path="rank" element={<PlayersTable />} />
            </Routes>
        </div>
        </Router>
    )
}


export default App;
