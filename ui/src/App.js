import React, { useState, useEffect } from 'react';
import PlayersTableRank from './modules/PlayersTable';
import PlayersCardsList from './modules/PlayersCard';


function App(){
    return (
        <div>
            <PlayersTableRank />
            <PlayersCardsList />
        </div>
    )
}


export default App;
