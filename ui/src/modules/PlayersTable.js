import React, { useEffect, useState } from 'react';
import { useTransition, animated } from 'react-spring';
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

    // Use the useTransition hook to animate the rows of the table as they are added or removed
    const transitions = useTransition(players, player => player.nickname, {
        from: { transform: 'translate3d(-100%,0,0)' },
        enter: { transform: 'translate3d(0%,0,0)' },
        leave: { transform: 'translate3d(100%,0,0)' },
    });

    return (
        <div>
            {isLoading ? (
                <div>Loading...</div>
            ) : (
                <table>
                    <thead>
                    <tr>
                        <th>Jugador</th>
                        <th>Edad</th>
                        <th>Rank</th>
                        <th>Elo</th>
                        <th>Promedio gol</th>
                        <th>PJ</th>
                        <th>PG</th>
                        <th>PP</th>
                        <th>DIF</th>
                    </tr>
                    </thead>
                    <tbody>
                    {/* Render the rows of the table using the transitions array */}
                    {transitions.map(({ item: player, key, props }) => (
                        <animated.tr key={key} style={props}>
                            <td>{player.nickname}</td>
                            <td>{player.age}</td>
                            <td>{player.rank}</td>
                            <td>{player.elo}</td>
                            <td>{player.goalsPerMatch}</td>
                            <td>{player.gamesPlayed}</td>
                            <td>{player.gamesWon}</td>
                            <td>{player.gamesLost}</td>
                            <td>{player.diff}</td>
                        </animated.tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default PlayersTableRank;
