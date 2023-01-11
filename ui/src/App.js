import React, { useState, useEffect } from 'react';
import PlayersList from './modules/player/PlayersList'
import PlayersTable from "./modules/player/PlayersTable";
import './App.css'

import MyNavbar from "./modules/header/Navbar";

export default App;


function App(){
    return (
        <div className="App" >
            <MyNavbar />
            <div className="contents">
                <PlayersTable />
                <PlayersList />
            </div>
        </div>
    )
}


