import React, { useState, useEffect } from 'react';
import PlayersList from './modules/player/PlayersList'
import PlayersTable from "./modules/player/PlayersTable";
import Header from "./modules/Header";

function App(){
    return (
        <div>
            <div className={"fa-2xs"}>
                <PlayersList />
            </div>
            <div>
                <PlayersTable />
            </div>
        </div>
    )
}


export default App;
