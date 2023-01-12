import React, { useEffect, useState } from 'react';
import {getColor, getStars, getPhoto} from "../Util";
import './PlayersTable.css';
import Container from 'react-bootstrap/Container';

function PlayersTableRank() {
    // Use the useState hook to create state variables for the player data and the loading state
    const [players, setPlayers] = useState([]);

    // Use the useEffect hook to fetch the player data from the REST API when the component mounts
    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/rank');
                const data = await response.json();
                setPlayers(data.players);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);



    return (
        <Container >
            <div className="title-container">
                <img src={require('../assets/carmen-league.png')} className="logo" alt="logo" />
                <h1>RANK 2023</h1>
                <h2>PREMIO: <b>USD 200</b></h2>
            </div>
            <div>
                <div className="note">
                    <h4>
                        <b>ELO: </b> es una medida estandarizada para evaluar el rendimiento del jugador, que toma en cuenta factores como los goles totales, partidos ganados/perdidos, el ranking del jugador, entre otros.
                    </h4>
                </div>
                <div className="note">
                    <h4>
                    <b>Rank: </b>es la cantidad de libertadores del jugador, y hablando en serio, es simplemente un numero [1-5] que me ayuda para el algoritmo de armar equipos, fue elegido por criterio propio
                    </h4>
                </div>
            </div>

            <table>
                <thead>
                <tr>
                    <th></th>
                    <th>Jugador</th>
                    <th>Rank</th>
                    <th>ELO</th>
                    <th>POS</th>
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
                        <td ><img src={getPhoto(player.nickname)} className={"photo"}  alt={player.nickname}></img></td>
                        <td className={"nickname"}>{player.nickname}  </td>
                        <td className={"rank"}>{getStars(player.rank)}</td>
                        <td className={"elo"}>{player.elo}</td>
                        <td className={"position"} style={{color: getColor(player.position)}}>{player.position.substring(0,3).toUpperCase()}</td>
                        <td className={"info"}>{player.goalsPerMatch}</td>
                        <td className={"info"}>{player.gamesPlayed}</td>
                        <td className={"info"}>{player.gamesWon}</td>
                        <td className={"info"}>{player.gamesLost}</td>
                        <td className={"info"}>{player.diff}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </Container>
)}

export default PlayersTableRank;
