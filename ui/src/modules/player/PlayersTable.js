import React, { useEffect, useState } from "react";
import { getColor, getStars, getPhoto } from "../Util";
import jsonPlayers from "../../resources/players.json";
import "./PlayersTable.css";

function PlayersTableRank() {
    // Use the useState hook to create state variables for the player data and the loading state
    const [players, setPlayers] = useState([]);

    // Use the useEffect hook to fetch the player data from the REST API when the component mounts
    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch("http://localhost:8080/players/rank");
                const data = await response.json();
                setPlayers(data.players);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);

    if (players.length === 0) {
        setPlayers(jsonPlayers.players);
    }

    return (
        <div className="container">
            <div className="headline">
                <img src={require("../../assets/cda-league.png")} className="logo" alt="logo"/>
                <div className="container">
                    <p className="note">
                        <b>ELO</b><br/>
                        Medida de estandarizacion que mide el rendimiento del
                        jugador, <b>no influye en la tabla</b>, metrica usada en el
                        algoritmo de armado de equipos.
                    </p>
                </div>
                <div className="container">
                    <p className="note">
                        {" "}
                        <b>RANK</b><br/>Cantidad de libertadores del jugador, y hablando en
                        serio, mide la <u>calidad</u> que tiene el jugador, me sirve para el
                        calculo del ELO.
                    </p>
                </div>
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
                        <td className={"index"}>{index + 1} </td>
                        <td className={""}>
                            <img
                                src={getPhoto(player.nickname)}
                                className={"photo"}
                                alt={player.nickname}
                            ></img>
                        </td>
                        <td className={"nickname"}>{player.nickname} </td>
                        <td className={"rank"}>{getStars(player.rank)}</td>
                        <td className={"elo"}>{player.elo}</td>
                        <td
                            className={"position"}
                            style={{ color: getColor(player.position) }}
                        >
                            {player.position.substring(0, 3).toUpperCase()}
                        </td>
                        <td className={"info"}>{player.totalGoals}</td>
                        <td className={"info"}>{player.gamesPlayed}</td>
                        <td className={"info"}>{player.gamesWon}</td>
                        <td className={"info"}>{player.gamesLost}</td>
                        <td className={"info"}>{player.diff}</td>
                        <td className={"points"}>{player.points}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
}

export default PlayersTableRank;
