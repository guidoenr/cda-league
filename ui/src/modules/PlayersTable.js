import React, {useEffect, useState} from 'react';
import './PlayersTable.css';

function PlayersTable(props) {
    const { players } = props;

    return (
        <div className="table-container">
        <table className="players-table">
            <thead>
            <tr>
                <th>Apodo</th>
                <th>Nombre</th>
                <th>Edad</th>
                <th>Rank</th>
                <th>Posicion</th>
                <th>Elo</th>
                <th>Promedio gol</th>
                <th>PJ</th>
                <th>PG</th>
                <th>PP</th>
                <th>DIF</th>
            </tr>
            </thead>
            <tbody>
            {players.map(player => (
                <tr key={player.ID}>

                    <td>{player.nickname}</td>
                    <td>{player.name}</td>
                    <td>{player.age}</td>
                    <td>{player.rank}</td>
                    <td>{player.position}</td>
                    <td>{player.elo}</td>
                    <td>{player.goalsPerMatch}</td>
                    <td>{player.gamesPlayed}</td>
                    <td>{player.gamesWon}</td>
                    <td>{player.gamesLost}</td>
                    <td>{player.diff}</td>
                </tr>
            ))}
            </tbody>
        </table>
        </div>
    );
}

function PlayersTableRank() {
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

export default PlayersTableRank;
