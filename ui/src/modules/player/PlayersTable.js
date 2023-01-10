import React, { useEffect, useState } from 'react';
import {getColor, getStars, getPhoto} from "../Util";
import './PlayersTable.css';

const PlayersTableRank = () => {
    // Use the useState hook to create state variables for the player data and the loading state
    const [players, setPlayers] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    // Use the useEffect hook to fetch the player data from the REST API when the component mounts
    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/rank');
                const data = await response.json();
                setPlayers(data.players);
                setIsLoading(false);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);



    return (
        <div>
            {isLoading ? (
                <div>Loading...</div>
            ) : (
                <div className={"table-container"}>
                <table>
                    <thead>
                    <tr>
                        <th></th>
                        <th>JUGADOR</th>
                        <th>RANK</th>
                        <th>POS</th>
                        <th>ELO</th>
                        <th>GOL</th>
                        <th>PJ</th>
                        <th>PG</th>
                        <th>PP</th>
                        <th>DIF</th>
                    </tr>
                    </thead>
                    <tbody>
                    {players.map((player) => (

                        <tr key={player.nickname}>
                            <td><img src={getPhoto(player.nickname)} className={"player-photo"} alt={player.nickname}></img></td>
                            <td>{player.nickname} <a className={"player-table-name"}> {player.name}</a> </td>
                            <td className={"star"}>{getStars(player.rank)}</td>
                            <td className={"player-position"} style={{color: getColor(player.position)}}>{player.position.substring(0,3).toUpperCase()}</td>
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
            )}
        </div>
    );
};

export default PlayersTableRank;
