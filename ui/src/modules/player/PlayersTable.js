import React, { useEffect, useState } from 'react';
import {getColor, getStars, getPhoto} from "../Util";
import './PlayersTable.css';
import Container from 'react-bootstrap/Container';
import jsonPlayers from '../../resources/players.json'
import 'animate.css';

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

    console.log(players)

    if (players.length === 0){
        setPlayers(jsonPlayers)
    }

    return (
        <Container className="Container animate__animated animate__bounceInDown">
            <div>
                <img src={require('../assets/carmen-league.png')} className="logo" alt="logo" />
                <img src={require('../assets/thropie.png')} className="thropie" alt="thropie" />
                <h1>TORNEO DE VERANO - 2023</h1>
                <h5>PREMIO: <b>100 USD</b></h5>
                <Container>
                <p className="note"><b>ELO:</b> Medida de estandarizacion que mide el rendimiento del jugador, <b>no influye en la tabla</b>, metrica usada en el algoritmo de armado de equipos. El ELO se basa en los partidos ganados/perdidos, goles, diferencia, etc. Donde el mayor multiplicador es el Rank. </p>
                <p className="note"><b>RANK:</b> Cantidad de libertadores del jugador, y hablando en serio, mide la <u>calidad</u> que tiene el jugador, me sirve para el calculo del ELO.</p>
                </Container>
            </div>
            <table>
                <thead>
                <tr>
                    <th>#</th>
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
                    <th>PTS</th>
                </tr>
                </thead>
                <tbody>
                {players.map((player, index) => (

                    <tr key={player.nickname}>
                        <td className={"index"}>{index+1} </td>
                        <td className={""}><img src={getPhoto(player.nickname)} className={"photo"}  alt={player.nickname}></img></td>
                        <td className={"nickname"}>{player.nickname}  </td>
                        <td className={"rank"}>{getStars(player.rank)}</td>
                        <td className={"elo"}>{player.elo}</td>
                        <td className={"position"} style={{color: getColor(player.position)}}>{player.position.substring(0,3).toUpperCase()}</td>
                        <td className={"info"}>{player.totalGoals}</td>
                        <td className={"info"}>{player.gamesPlayed}</td>
                        <td className={"info"}>{player.gamesWon}</td>
                        <td className={"info"}>{player.gamesLost}</td>
                        <td className={"info"}>{player.diff}</td>
                        <td className={"info"}>{player.points}</td>
                    </tr>
                ))}
                </tbody>
            </table>
      {/*      <div className="note">
                <p>
                    congelado tenes el pecho
                </p>
            </div>*/}
        </Container>
)}

export default PlayersTableRank;
