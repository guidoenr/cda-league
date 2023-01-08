import React, { useState, useEffect } from 'react';
import PlayersTable from './PlayersTable';

function App(){
    return (
        <div>
            <PlayersList />
        </div>
    )
}

function PlayersList() {
    const [players, setPlayers] = useState([]);

    useEffect(() => {
        // Fetch players data from API
        fetch('http://localhost:8080/players/')
            .then(response => response.json())
            .then(data => setPlayers(data.players));
    }, []);

    return (
        <div>
            <h2>Players</h2>
            <PlayersTable players={players} />
        </div>
    );
}

export default App;
