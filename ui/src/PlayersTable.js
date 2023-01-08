import React from 'react';

function PlayersTable(props) {
    const { players } = props;

    return (
        <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Nickname</th>
                <th>Name</th>
                <th>Description</th>
                <th>Age</th>
                <th>Rank</th>
                <th>Position</th>
                <th>Elo</th>
                <th>Goals per match</th>
                <th>Games won</th>
            </tr>
            </thead>
            <tbody>
            {players.map(player => (
                <tr key={player.ID}>
                    <td>{player.ID}</td>
                    <td>{player.nickname}</td>
                    <td>{player.name}</td>
                    <td>{player.description}</td>
                    <td>{player.age}</td>
                    <td>{player.rank}</td>
                    <td>{player.position}</td>
                    <td>{player.elo}</td>
                    <td>{player.goalsPerMatch}</td>
                    <td>{player.gamesWon}</td>
                </tr>
            ))}
            </tbody>
        </table>
    );
}

export default PlayersTable;
