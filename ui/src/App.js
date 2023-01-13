import React, { useState, useEffect } from 'react';
import PlayersList from './modules/player/PlayersList'
import Footer from './modules/footer/Footer'
import PlayersTable from "./modules/player/PlayersTable";
import Match from './modules/match/Match'
import {BrowserRouter as Router , Route, Routes} from "react-router-dom";
import './App.css'

import MyNavbar from "./modules/navbar/Navbar";
import PlayersTableRank from "./modules/player/PlayersTable";

function App(){

    return (
        <Router>
        <div className="App App-bg" >
            <MyNavbar />
            <Routes >
                <Route path="/" exact element={<PlayersTableRank />} />
                <Route path="players" element={<PlayersList />} />
                <Route path="rank" element={<PlayersTable />} />
                <Route path="match" element={<Match />} />
            </Routes>
            <Footer />
        </div>
        </Router>
    )
}


export default App;
