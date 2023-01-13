import React, { useState, useEffect, useTransition } from 'react';
import PlayerCard from './PlayerCard'; // Import the PlayerCard component
import './PlayersList.css'
import jsonPlayers from '../../resources/players.json'

const PlayerList = () => {
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
        fetchPlayers()
    }, []);


    if (players.length === 0){
        setPlayers(jsonPlayers)
    }

    return (
            <div className="player-list">
                    {/* Map over the players array and render a PlayerCard component for each player */}
                    {players.map(player => (
                        <PlayerCard
                            key={player.id}
                            nickname={player.nickname}
                            name={player.name}
                            rank={player.rank}
                            position={player.position}
                            description={player.description}
                            elo={player.elo}
                        />
                    ))}
                </div>
    );
};

export default PlayerList;
