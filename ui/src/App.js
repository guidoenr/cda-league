import React, { useState, useEffect } from 'react';
import PlayersTable from './PlayersTable';
import PlayersCard from './PlayersCard';




function App(){
    return (
        <div>
            {/*<PlayersList />*/}
            <PlayersRanked />
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
        <div className="center">
            <h2>Players</h2>
            <PlayersTable players={players} />
        </div>
    );
}



function PlayersRanked() {
    const [players, setPlayers] = useState([]);

    useEffect(() => {
        // Fetch players data from API
        fetch('http://localhost:8080/players/rank')
            .then(response => response.json())
            .then(data => setPlayers(data.players));
    }, []);

    return (
        <div className="center">
            <h2>Players Cards</h2>
            <PlayersCard players={players} />
        </div>
    );
}


export default App;
